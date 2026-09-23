package main

import "github.com/gin-gonic/gin"

func main3() {
	engine := gin.Default()

	{
		g1 := engine.Group("/v1")
		g1.Use(M6)
		g1.GET("/a", func(ctx *gin.Context) {
			ctx.String(200, "name=libai")
		})
		g1.GET("/b", func(ctx *gin.Context) {
			ctx.String(200, "age=18")
		})
	}
	{
		g2 := engine.Group("/v2")
		g2.GET("/a", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"name": "libai"})
		})
		g2.GET("/b", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{"age": 18})
		})
	}

	engine.Run("127.0.0.1:5678")
}
