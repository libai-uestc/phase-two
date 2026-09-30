package util_test

import (
	"encoding/base64"
	"fmt"
	"libai/go/phase-two/post/util"
	"strings"
	"testing"
	"time"
)

func TestBase64(t *testing.T) {
	text := "你好，李白"
	cipher := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Println(cipher)
	bs, _ := base64.StdEncoding.DecodeString(cipher)
	if string(bs) != text {
		t.Fail()
	}
}

func TestJWT(t *testing.T) {
	secret := "123456"
	header := util.DefautHeader
	payload := util.JwtPayload{
		ID:          "rj4t49tu49",
		Issue:       "微信",
		Audience:    "英雄联盟",
		Subject:     "购买道具",
		IssueAt:     time.Now().Unix(),
		Expiration:  time.Now().Add(2 * time.Hour).Unix(),
		UserDefined: map[string]any{"name": strings.Repeat("李白", 100)},
	}

	if token, err := util.GenJWT(header, payload, secret); err != nil {
		fmt.Printf("生成JSON web token失败:%v", err)
	} else {
		fmt.Println(token)
		if _, p, err := util.VerifyJwt(token, secret); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("JWT验证通过。欢迎 %s !\n", p.UserDefined["name"])
		}
	}
}

// go test -v ./post/util -run=^TestBase64$ -count=1
// go test -v ./post/util -run=^TestJWT$ -count=1
