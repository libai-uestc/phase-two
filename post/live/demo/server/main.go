package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	// "sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Name     string `form:"name"`
	Password string `form:"pass"`
}

var (
// tokenMap = sync.Map{}
)

const (
	COOKIE_NAME = "auth"
)

func Login(ctx *gin.Context) {
	// 获取前端传过来的参数
	name := ctx.PostForm("name")
	password := ctx.PostForm("pass")
	// fmt.Printf("=> 前端传来的name: [%s]\n", name)
	// fmt.Printf("=> 前端传来的password: [%s]\n", password)
	// 对参数进行合法性校验
	if len(name) < 2 {
		ctx.String(http.StatusBadRequest, "用户名过短！")
		return
	}

	// var request LoginRequest
	// ctx.ShouldBind(&request) // 参数绑定和校验
	// 查询数据库
	user, err := GetPasswordByName(name)
	// fmt.Printf("=> 数据库查出的hash: [%s], 错误: %v\n", dbPass, err)
	if err == nil {
		// 123456 + Salt
		// Bcrypt算法
		if bcrypt.CompareHashAndPassword([]byte(user.PassWord), []byte(password)) != nil {
			ctx.String(http.StatusBadRequest, "用户名或密码错误")
		} else {
			log.Println("登录成功")
			uid := user.Id
			token := xid.New().String()
			// tokenMap.Store(token, uid)
			GetRedisClient().Set(context.Background(), token, uid, 7*24*time.Hour)
			// 响应头 放在响应头里面传给客户端
			ctx.SetCookie(COOKIE_NAME, token, 7*86400, "/", "localhost", false, true)
			// 所有的响应头必须在第一次TCP传输时传给TCP客户端。 客户端接收响应体可以分成多个批次来接收，而客户端接收响应体的时候，对响应体进行理解解析，依赖于响应头，因为响应头里可能有各种基本信息，如标识：响应体是什么格式，该怎么反序列化等信息
			ctx.String(http.StatusOK, "登录成功") // 响应体

		}
		// if dbPass!=password{
		// 	ctx.String(http.StatusBadRequest,"用户名或密码错误")
		// }else{
		// 	ctx.String(http.StatusOK,"登录成功")
		// }
	} else {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.String(http.StatusBadRequest, "用户名不存在")
		} else {
			ctx.String(http.StatusInternalServerError, "系统异常")
		}
	}
}

// kill -9 $pid
func ListenSignal() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	sig := <-ch
	fmt.Printf("接收到信号 %s, 准备终止go进程\n", sig.String())

	// 收尾工作
	CloseDB()

	os.Exit(0)
}

// Auth 登录验证的中间件
func Auth(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie(COOKIE_NAME)
	if err != nil {
		log.Printf("没传cookie %s", COOKIE_NAME)
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
		ctx.Abort() // 终止
		return
	}
	token := cookie.Value
	result := GetRedisClient().Get(context.Background(), token)
	// if v, exists := tokenMap.Load(token); !exists {
	if result.Err() != nil {
		// redis.Nil
		log.Printf("非法的cookie %s=%s", COOKIE_NAME, token)
		ctx.Redirect(http.StatusTemporaryRedirect, "/login")
		ctx.Abort() // 终止
		return
	} else {
		if v, err := result.Int(); err != nil {
			log.Printf("uid不是int: %#v", v)
		} else {
			log.Printf("uid=%d", v)
			ctx.Set("uid", v)
		}
		// if uid, ok := v.(int); ok {
		// 	log.Printf("uid=%d", uid)
		// 	ctx.Set("uid", uid)
		// } else {
		// 	log.Printf("uid不是int: %#v", v)
		// }
	}
}

func main() {
	go ListenSignal()

	router := gin.Default()
	// 加载模板文件
	router.LoadHTMLGlob("post/live/demo/views/*.html")
	// 路由
	router.GET("/login", func(context *gin.Context) {
		context.HTML(http.StatusOK, "login.html", gin.H{})
	})
	router.POST("/login/submit", Login)
	router.GET("/center", Auth, func(context *gin.Context) {
		v := context.Value("uid")
		if v != nil {
			if uid, ok := v.(int); ok {
				context.String(http.StatusOK, strconv.Itoa(uid)+" 登录成功")
			}
		}
		context.String(http.StatusOK, "登录成功")
	})
	// 启动web server
	if err := router.Run("127.0.0.1:1234"); err != nil {
		panic(err)
	}
}
