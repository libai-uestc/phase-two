package main

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	// database "libai/go/phase-two/post/database/gorm"
	"libai/go/phase-two/post/database"
	mysql "libai/go/phase-two/post/database/gorm"
	"libai/go/phase-two/post/handler/model"
	"libai/go/phase-two/post/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

const (
	SESSION_KEY_PREFIX = "session_"
	SESSION_LIFE       = 86400
	COOKIE_NAME        = "session_id"
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

	sessionID := xid.New().String()
	ctx.SetCookie(
		COOKIE_NAME,
		sessionID,
		SESSION_LIFE,
		"/",
		"localhost",
		false,
		true,
	)
	log.Printf("种下cookie %s", COOKIE_NAME)
	info, _ := json.Marshal(userInfo{Name: user2.Name, Id: user2.Id})
	database.GetRedisClient().Set(context.Background(), SESSION_KEY_PREFIX+sessionID, info, SESSION_LIFE*time.Second)
	// return
}

func AuthMiddleWare(ctx *gin.Context) {
	if cookie, err := ctx.Request.Cookie(COOKIE_NAME); err == nil {
		sessionID := cookie.Value
		result := database.GetRedisClient().Get(context.Background(), SESSION_KEY_PREFIX+sessionID)
		if result.Err() == nil {
			info := result.Val()
			var user userInfo
			if err := json.Unmarshal([]byte(info), &user); err == nil {
				ctx.Set("user_name", user.Name)
				ctx.Set("user_id", strconv.Itoa(user.Id))
				return
			} else {
				log.Printf("user info反序列化失败: %s", err)
			}
		} else {
			log.Printf("找不到session id %s: %s", sessionID, err)
		}
	} else {
		log.Printf("读不到cookie %s", COOKIE_NAME)
	}
	ctx.Redirect(http.StatusFound, "/login")
}

func main() {
	mysql.ConnectPostDB("./post/conf", "db", util.YAML, "./log")
	router := gin.Default()

	router.Static("/js", "post/views/js")
	router.Static("/css", "post/views/js")
	router.StaticFile("/favicon.ico", "post/views/img/迈克尔乔丹.png")
	router.LoadHTMLGlob("post/views/html/*")

	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "session_login.html", nil)
	})
	router.POST("/login/submit", Login)
	router.GET("/page1", AuthMiddleWare, func(ctx *gin.Context) {
		uname, _ := ctx.Value("user_name").(string)
		uid, _ := ctx.Value("user_id").(string)
		ctx.String(200, "这是page1,欢迎光临"+uname+"["+uid+"]")
	})
	router.GET("/page2", AuthMiddleWare, func(ctx *gin.Context) {
		uname, _ := ctx.Value("user_name").(string)
		uid, _ := ctx.Value("user_id").(string)
		ctx.String(200, "这是page2，欢迎光临"+uname+"["+uid+"]")
	})
	if err := router.Run("127.0.0.1:5678"); err != nil {
		panic(err)
	}
}
