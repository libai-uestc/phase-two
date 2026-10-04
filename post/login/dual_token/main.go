package main

import (
	"context"
	"libai/go/phase-two/post/database"
	mysql "libai/go/phase-two/post/database/gorm"
	"libai/go/phase-two/post/handler/model"
	"libai/go/phase-two/post/util"
	"log"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

const (
	REFRESH_KEY_PREFIX  = "session_"
	REFRESH_TOKEN_LIFE  = 60
	REFRESH_COOKIE_NAME = "refresh"
	ACCESS_COOKIE_NAME  = "access"

	SECRET = "f4398"
)

type userInfo struct {
	Name string
	Id   int
}

func Login(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.String(http.StatusBadRequest, util.BindErrMsg(err))
		return
	}
	user2 := mysql.GetUserByName(user.Name)
	if user2 == nil {
		ctx.String(http.StatusBadRequest, "用户名不存在")
		return
	}
	if user2.PassWord != user.PassWord {
		ctx.String(http.StatusBadRequest, "密码错误")
		return
	}

	refreshToken := xid.New().String()
	ctx.SetCookie(
		REFRESH_COOKIE_NAME,
		refreshToken,
		REFRESH_TOKEN_LIFE,
		"/",
		"localhost",
		false,
		true,
	)
	header := util.DefaultHeader
	payload := util.JwtPayload{
		Issue:       "dual_token",
		IssueAt:     time.Now().Unix(),
		Expiration:  0,
		UserDefined: map[string]any{"user_id": strconv.Itoa(user2.Id), "user_name": user2.Name},
	}
	if accessToken, err := util.GenJWT(header, payload, SECRET); err != nil {
		slog.Error("生成access token失败", "error", err)
	} else {
		ctx.SetCookie(
			ACCESS_COOKIE_NAME,
			accessToken,
			0,
			"/",
			"localhost",
			false,
			true,
		)
		database.GetRedisClient().Set(context.Background(), REFRESH_KEY_PREFIX+refreshToken, accessToken, REFRESH_TOKEN_LIFE*time.Second)
	}
	// return
}

func AuthMiddleWare(ctx *gin.Context) {
	if cookie, err := ctx.Request.Cookie(ACCESS_COOKIE_NAME); err == nil {
		accessToken := cookie.Value
		_, payload, err := util.VerifyJwt(accessToken, SECRET)
		if err == nil {
			log.Println("直接根据access token拿到了用户的身份信息")
			ctx.Set("user_name", payload.UserDefined["user_name"])
			ctx.Set("user_id", payload.UserDefined["user_id"])
			return
		}
	} else {
		log.Println("cookie里没有access token")
	}
	if cookie, err := ctx.Request.Cookie(REFRESH_COOKIE_NAME); err == nil {
		refreshToken := cookie.Value
		result := database.GetRedisClient().Get(context.Background(), REFRESH_KEY_PREFIX+refreshToken)
		if result.Err() == nil {
			log.Println("根据cookie里的refresh token重新获得了access token")
			accessToken := result.Val()
			_, payload, err := util.VerifyJwt(accessToken, SECRET)
			if err == nil {
				ctx.SetCookie(
					ACCESS_COOKIE_NAME,
					accessToken,
					0,
					"/",
					"localhost",
					false,
					true,
				)
				log.Println("把access token种到了浏览器的cookie里")
				ctx.Set("user_name", payload.UserDefined["user_name"])
				ctx.Set("user_id", payload.UserDefined["user_id"])
				return
			}
		}
	}
	ctx.Redirect(http.StatusFound, "/login")
}

func main() {
	mysql.ConnectPostDB("./post/conf", "db", util.YAML, "./log")
	router := gin.Default()

	router.Static("/js", "post/views/js") //在url是访问目录/js相当于访问文件系统中的views/js目录
	router.Static("/css", "post/views/css")
	router.StaticFile("/favicon.ico", "post/views/img/迈克尔乔丹.png") //在url中访问文件/favicon.ico，相当于访问文件系统中的views/img/dqq.png文件
	router.LoadHTMLGlob("post/views/html/*")                      //使用这些.html文件时就不需要加路径了

	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "session_login.html", nil)
	})
	router.POST("/login/submit", Login)
	router.GET("/page1", AuthMiddleWare, func(ctx *gin.Context) { // 使用身份认证中间件
		// 从ctx里取出用户信息
		uname, _ := ctx.Value("user_name").(string)
		uid, _ := ctx.Value("user_id").(string)
		ctx.String(200, "这是page1, 欢迎 "+uname+"["+uid+"]")
	})
	//需要登录才能访问的页面2
	router.GET("/page2", AuthMiddleWare, func(ctx *gin.Context) { // 使用身份认证中间件
		// 从ctx里取出用户信息
		uname, _ := ctx.Value("user_name").(string)
		uid, _ := ctx.Value("user_id").(string)
		ctx.String(200, "这是page2, 欢迎 "+uname+"["+uid+"]")
	})

	if err := router.Run("127.0.0.1:5678"); err != nil {
		panic(err)
	}
}
