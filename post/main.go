package main

import (
	database "libai/go/phase-two/post/database/gorm"
	handler "libai/go/phase-two/post/handler/gin"
	"libai/go/phase-two/post/util"

	"github.com/gin-gonic/gin"
)

func Init() {
	util.InitSlog("./log/post.log")
	database.ConnectPostDB("./post/conf", "db", util.YAML, "./log")
}

func main() {
	Init()

	engine := gin.Default()

	engine.Static("/js", "post/views/js")
	engine.Static("/css", "post/views/css")
	engine.StaticFile("/favicon.ico", "post/views/img/迈克尔乔丹.png")
	engine.LoadHTMLGlob("post/views/html/*")

	engine.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(200, "login.html", nil)
	})
	engine.POST("/login", func(ctx *gin.Context) { ctx.HTML(200, "login.html", nil) })
	engine.GET("/regist", func(ctx *gin.Context) {
		ctx.HTML(200, "user_regist.html", nil)
	})

	engine.GET("/modify_pass", func(ctx *gin.Context) {
		ctx.HTML(200, "update_pass.html", nil)
	})

	engine.POST("/login/submit", handler.Login)
	engine.POST("/regist/submit", handler.RegistUser)
	engine.POST("/modify_pass/submit", handler.UpdatePassword)
	// engine.GET("/user", handler.)
	engine.Run("localhost:5678")
}
