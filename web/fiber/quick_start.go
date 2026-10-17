package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func homeHandler(ctx fiber.Ctx) error {
	fmt.Println("协议:", ctx.Protocol())
	fmt.Println("请求方法:", ctx.Method())
	fmt.Println("请求URL:", ctx.OriginalURL())
	fmt.Println("请求路径:", ctx.Path())
	fmt.Println("请求头:")
	for k, v := range ctx.GetReqHeaders() {
		fmt.Printf("%s=%s\n", k, v[0])
	}
	fmt.Println("请求体:", string(ctx.Request().Body()))
	fmt.Println("请求体:", string(ctx.Request().Body()))

	ctx.Set("Set-Cookie", "language=go; Path=/; Domain=www.baidu.com; Max-Age=604800; HttpOnly")
	ctx.SendStatus(http.StatusOK)
	ctx.Status(http.StatusOK)
	ctx.SendString("Welcome")
	ctx.WriteString(" to 江浙沪")

	fmt.Println()
	return nil

}

func main2() {
	app := fiber.New()
	logFile, err := os.OpenFile("C:/Users/18101/Desktop/第二阶段/phase-two/log/fiber.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	app.Use(logger.New(
		logger.Config{
			TimeFormat:    "2006-01-02 - 15:04:05",
			TimeZone:      "Asia/Shanghai",
			DisableColors: true,
		},
	))
	app.Use(recover.New())

	app.Get("/home", homeHandler)
	app.Post("/home", homeHandler)
	app.Head("/ping", func(ctx fiber.Ctx) error {
		ctx.Status(200)
		return nil
	})
	if err := app.Listen("127.0.0.1:5678"); err != nil {
		slog.Error("fiber app start failed", "error", err)
	}
}

// go run ./web/fiber
