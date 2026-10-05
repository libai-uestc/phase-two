package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Name     string `form:"name"`
	Password string `form:"pass"`
}

func Login(ctx *gin.Context) {
	// 获取前端传过来的参数
	name := ctx.PostForm("name")
	password := ctx.PostForm("pass")
	fmt.Printf("=> 前端传来的name: [%s]\n", name)
	fmt.Printf("=> 前端传来的password: [%s]\n", password)
	// 对参数进行合法性校验
	if len(name) < 2 {
		ctx.String(http.StatusBadRequest, "用户名过短！")
		return
	}

	// var request LoginRequest
	// ctx.ShouldBind(&request) // 参数绑定和校验
	// 查询数据库
	dbPass, err := GetPasswordByName(name)
	fmt.Printf("=> 数据库查出的hash: [%s], 错误: %v\n", dbPass, err)
	if err == nil {
		// 123456 + Salt
		// Bcrypt算法
		if bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(password)) != nil {
			ctx.String(http.StatusBadRequest, "用户名或密码错误")
		} else {
			ctx.String(http.StatusOK, "登录成功")
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
	router.GET("/center", func(context *gin.Context) {
		context.String(http.StatusOK, "登录成功")
	})
	// 启动web server
	if err := router.Run("127.0.0.1:1234"); err != nil {
		panic(err)
	}
}
