package main

import (
	database "libai/go/phase-two/post/database/gorm"
	handler "libai/go/phase-two/post/handler/gin"
	"libai/go/phase-two/post/util"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/robfig/cron/v3"
)

func Init() {
	util.InitSlog("./log/post.log")
	database.ConnectPostDB("./post/conf", "db", util.YAML, "./log")

	crontab := cron.New()
	crontab.AddFunc("*/30 * * * *", database.PingPostDB)
	crontab.Start()

}

func ListenTermSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	slog.Info("receive term signal " + sig.String() + ", going to exit")
	database.ClosePostDB()
	os.Exit(0)
}

func main() {
	Init()
	go ListenTermSignal()

	engine := gin.Default()

	engine.Static("/js", "post/views/js")
	engine.Static("/css", "post/views/css")
	engine.StaticFile("/favicon.ico", "post/views/img/迈克尔乔丹.png")
	engine.LoadHTMLGlob("post/views/html/*")

	engine.Use(handler.Metric)
	engine.GET("/metrics", func(ctx *gin.Context) {
		promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
	})
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
	engine.GET("/user", handler.GetUserInfo)
	engine.GET("/logout", handler.Logout)

	group := engine.Group("/news")
	group.GET("", handler.NewsList)
	group.GET("/issue", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "news_issue.html", nil)
	})
	group.POST("/issue/submit", handler.Auth, handler.PostNews)
	group.GET("/belong", handler.NewsBelong)
	group.GET("/:id", handler.GetNewsById)
	group.GET("/delete/:id", handler.Auth, handler.DeleteNews)
	group.POST("/update", handler.Auth, handler.UpdateNews)

	engine.GET("", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "news")
	})

	if err := engine.Run("localhost:5678"); err != nil {
		panic(err)
	}
}
