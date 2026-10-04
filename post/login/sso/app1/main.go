package main

import (
	"encoding/json"
	"io"
	"libai/go/phase-two/post/login/sso"
	"log"
	"net/http"
)

const (
	SESSION_KEY_PREFIX = "app1_session_"
)

func checkToken(token string) (userName, userId string, valid bool) {
	resp, err := http.Get("http://" + sso.SSO_URL + "/identify?" + sso.SSO_TOKEN_QUERY_NAME + "=" + token)
	if err != nil {
		log.Println(err)
		valid = false
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		valid = false
		return
	}
	var mp map[string]string
	bs, _ := io.ReadAll(resp.Body)
	json.Unmarshal(bs, &mp)
	userId = mp[sso.KEY_OF_UID]
	userName = mp[sso.KEY_OF_NAME]
	valid = true
	return
}
