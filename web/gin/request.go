package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func url(engine *gin.Engine) {
	engine.GET("/student", func(ctx *gin.Context) {
		a := ctx.Query("name")
		b := ctx.DefaultQuery("addr", "China")
		ctx.String(http.StatusOK, a+" live in "+b)
	})
}

func restful(engine *gin.Engine) {
	engine.GET("/student/:name/*addr", func(ctx *gin.Context) {
		name := ctx.Param("name")
		addr := ctx.Param("addr")
		ctx.String(http.StatusOK, name+" live in "+addr)
	})
}
