package main

import (
	"log/slog"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	counter sync.Map
)

func init() {
	counter.Store("123456789", 0)
	counter.Store("987654321", 0)
}

func CountMD(ctx *gin.Context) {
	token, err := ctx.Cookie("token")
	if err != nil {
		ctx.Abort()
	}
	if v, exists := counter.Load(token); !exists {
		ctx.Abort()
	} else {
		c, _ := v.(int)
		counter.Store(token, c+1)
		slog.Info("visit counter", "token", token, "count", c+1)
	}
}

func SetCookie(ctx *gin.Context) {
	name := "language"
	value := "go"
	maxAge := 86400 * 7
	path := "/"
	domain := "www.baidu.com"
	secure := false
	httpOnly := true
	ctx.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
}

func main() {
	engine := gin.Default()
	engine.Use(CountMD)
	engine.GET("/ck", SetCookie)
	engine.Run("127.0.0.1:5678")
}
