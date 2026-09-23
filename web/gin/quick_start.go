package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func homeHandler(ctx *gin.Context) {
	fmt.Println("请求头：")
	for k, v := range ctx.Request.Header {
		fmt.Printf("%s=%s\n", k, v[0])
	}
	fmt.Println("请求体：")
	// io.ReadAll(ctx.Request.Body)
	io.Copy(os.Stdout, ctx.Request.Body)

	// 先设置响应头
	ctx.Writer.Header().Add("language", "go")
	ctx.Header("Strict-Transport-Security", "max-age=31536000;includeSubDomains; preload")

	// 再设置响应码/状态码
	ctx.Writer.WriteHeader(http.StatusOK)

	// 设置响应体
	ctx.Writer.WriteString("welcome")

	ctx.String(200, " to BeiJing ")
	ctx.JSON(200, map[string]any{"物理": 34, "化学": 70})
	ctx.JSON(200, gin.H{"物理": 34, "化学": 70})
}

func main1() {
	engine := gin.Default()

	engine.GET("/home", homeHandler)

	if err := engine.Run("127.0.0.1:5678"); err != nil {
		panic(err)
	}
}

// go run .\web\gin\
