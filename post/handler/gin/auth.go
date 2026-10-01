package handler

import (
	"libai/go/phase-two/post/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	UID_IN_TOKEN = "uid"
	UID_IN_CTX   = "uid"
	COOKIE_NAME  = "jwt"
)

var (
	KeyConfig = util.InitViper("post/conf", "jwt", util.YAML)
)

func GetLoginUid(ctx *gin.Context) int {
	token := ""
	for _, cookie := range ctx.Request.Cookies() {
		if cookie.Name == COOKIE_NAME {
			token = cookie.Value
		}
	}
	return GetUidFromJwt(token)
}

func GetUidFromJwt(token string) int {
	_, payload, err := util.VerifyJwt(token, KeyConfig.GetString("secret"))
	if err != nil {
		return 0
	}
	for k, v := range payload.UserDefined {
		if k == UID_IN_TOKEN {
			return int(v.(float64))
		}
	}
	return 0
}

func Auth(ctx *gin.Context) {
	loginUid := GetLoginUid(ctx)
	if loginUid <= 0 {
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
		ctx.Abort()
	} else {
		ctx.Set(UID_IN_CTX, loginUid)
	}
}
