package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

var (
	DefaultHeader = JwtHeader{
		Algo: "HS256",
		Type: "JWT",
	}
)

type JwtHeader struct {
	Algo string `json:"alg"`
	Type string `json:"typ"`
}

type JwtPayload struct {
	ID          string         `json:"jti"`
	Issue       string         `json:"iss"`
	Audience    string         `json:"aud"`
	Subject     string         `json:"sub"`
	IssueAt     int64          `json:"iat"`
	NotBefore   int64          `json:"nbf"`
	Expiration  int64          `json:"exp"`
	UserDefined map[string]any `json:"ud"`
}

func GenJWT(header JwtHeader, payload JwtPayload, secret string) (string, error) {
	var part1, part2, signature string
	if bs1, err := json.Marshal(header); err != nil {
		return "", err
	} else {
		part1 = base64.RawURLEncoding.EncodeToString(bs1)
	}
	if bs2, err := json.Marshal(payload); err != nil {
		return "", err
	} else {
		part2 = base64.RawURLEncoding.EncodeToString(bs2)
	}
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(part1 + "." + part2))
	signature = base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	return part1 + "." + part2 + "." + signature, nil
}

func VerifyJwt(token string, secret string) (*JwtHeader, *JwtPayload, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, nil, fmt.Errorf("token是%d部分", len(parts))
	}
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(parts[0] + "." + parts[1]))
	signature := base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	if signature != parts[2] {
		return nil, nil, fmt.Errorf("验证失败")
	}

	var part1, part2 []byte
	var err error
	if part1, err = base64.RawURLEncoding.DecodeString(parts[0]); err != nil {
		return nil, nil, fmt.Errorf("header Base64反解失败")
	}
	if part2, err = base64.RawURLEncoding.DecodeString(parts[1]); err != nil {
		return nil, nil, fmt.Errorf("payload Base64反解失败")
	}

	var header JwtHeader
	var payload JwtPayload
	if err = json.Unmarshal(part1, &header); err != nil {
		return nil, nil, fmt.Errorf("header json反解失败")
	}
	if err = json.Unmarshal(part2, &payload); err != nil {
		return nil, nil, fmt.Errorf("payload json反解失败")
	}
	if payload.Expiration > 0 && payload.Expiration < time.Now().Unix() {
		return nil, nil, fmt.Errorf("JWT token已过期")
	}

	return &header, &payload, nil
}
