package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func M1(ctx *gin.Context) {
	ctx.String(200, "M1 Begin\n")
	ctx.Next()
	ctx.String(200, "M1 End\n")
}

func M2(ctx *gin.Context) {
	ctx.String(200, "Here is M2\n")
}

func M3(ctx *gin.Context) {
	ctx.String(200, "Here is M3\n")
}

func M4(ctx *gin.Context) {
	ctx.String(200, "Here is M4\n")
	ctx.Abort()
	ctx.String(200, "M4 Second Line\n")
}

func M5(ctx *gin.Context) {
	ctx.String(200, "Here is M5\n")
}

func M6(ctx *gin.Context) {
	slog.Info("visit", "path", ctx.Request.URL)
}

func main2() {
	engine := gin.Default()
	engine.Use(M6)
	engine.GET("/", M1, M2, M3, M2, M4, M5)
	engine.GET("/2", M2, M3, M2, M4, M5)
	engine.Run("127.0.0.1:5678")
}
