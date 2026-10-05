package main

import (
	"context"
	"encoding/json"
	"io"
	"libai/go/phase-two/post/database"
	"libai/go/phase-two/post/login/sso"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

const (
	SESSION_KEY_PREFIX = "app2_session_"
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

func Home(ctx *gin.Context) {
	if cookie, err := ctx.Request.Cookie(sso.APP2_TOKEN_COOKIE_NAME); err == nil {
		sessionID := cookie.Value
		result := database.GetRedisClient().Get(context.Background(), SESSION_KEY_PREFIX+sessionID)
		if result.Err() == nil {
			v := result.Val()
			var mp map[string]string
			json.Unmarshal([]byte(v), &mp)
			userId := mp[sso.KEY_OF_UID]
			userName := mp[sso.KEY_OF_NAME]
			log.Println("根据app token，身份验证成功")
			ctx.String(200, "这里是app2, 欢迎 "+userName+"["+userId+"]")
			return
		}
	}

	var token string
	if cookie, err := ctx.Request.Cookie(sso.SSO_TOKEN_COOKIE_NAME); err == nil {
		token = cookie.Value
	} else {
		token = ctx.Query(sso.SSO_TOKEN_QUERY_NAME)
	}
	if len(token) > 0 {
		userName, userId, valid := checkToken(token)
		if valid {
			log.Println("根据sso token，身份验证成功")
			ctx.SetCookie(
				sso.SSO_TOKEN_COOKIE_NAME,
				token,
				sso.SSO_TOKEN_LIFE,
				"/",
				"localhost",
				false,
				true,
			)
			sessionID := xid.New().String()
			ctx.SetCookie(
				sso.APP2_TOKEN_COOKIE_NAME,
				sessionID,
				sso.APP_TOKEN_LIFE,
				"/",
				"localhost",
				false,
				true,
			)
			info, _ := json.Marshal(map[string]string{sso.KEY_OF_UID: userId, sso.KEY_OF_NAME: userName})
			database.GetRedisClient().Set(context.Background(), SESSION_KEY_PREFIX+sessionID, string(info), sso.APP_TOKEN_LIFE*time.Second)
			ctx.String(200, "这里是app2, 欢迎光临 "+userName+"["+userId+"]")
			return
		} else {
			ctx.SetCookie(
				sso.SSO_TOKEN_COOKIE_NAME,
				"",
				-1,
				"/",
				"localhost",
				false,
				true,
			)
			url := "http://" + sso.SSO_URL + "/login?service=" + sso.APP2_URL + "/home"
			ctx.Redirect(http.StatusFound, url)
			return
		}
	} else {
		url := "http://" + sso.SSO_URL + "/login?service=" + sso.APP2_URL + "/home"
		ctx.Redirect(http.StatusFound, url)
		return
	}
}

func main() {
	engine := gin.Default()
	engine.GET("/home", Home)
	if err := engine.Run(sso.APP2_URL); err != nil {
		panic(err)
	}
}

// go run ./post/login/sso/app2
