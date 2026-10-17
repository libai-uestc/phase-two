package main

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func bizHandler(ctx fiber.Ctx) error {
	ctx.WriteString("welcome\n")
	return nil
}

func M1(ctx fiber.Ctx) error {
	ctx.WriteString("M1 Begin\n")
	ctx.Next()
	ctx.WriteString("M1 End\n")
	return nil
}

func M2(ctx fiber.Ctx) error {
	ctx.WriteString("Here is M2\n")
	return ctx.Next()
}

func M3(ctx fiber.Ctx) error {
	ctx.WriteString("Here is M3\n")
	return ctx.Next()
}

func M4(ctx fiber.Ctx) error {
	ctx.WriteString("Here is M4\n")
	return nil
}

func M5(ctx fiber.Ctx) error {
	ctx.WriteString("Here is M5\n")
	return ctx.Next()
}

func M6(ctx fiber.Ctx) error {
	slog.Info("visit", "path", ctx.Path)
	return ctx.Next()
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(M6)
	app.Use("/2", M2)
	app.Use([]string{"/3", "/4"}, M2, M4)

	app.Get("/1", bizHandler, M1, M2, M3, M4, M5)
	app.Get("/2", bizHandler, M3, M2, M4, M5)

	limit := limiter.New(
		limiter.Config{
			Next: func(c fiber.Ctx) bool {
				return c.IP() == "127.0.0.2"
			},
			Expiration: time.Minute,
			Max:        5,
			LimitReached: func(c fiber.Ctx) error {
				return c.SendString("请求过于频繁")
			},
			LimiterMiddleware: limiter.SlidingWindow{},
		},
	)

	app.Get("/lt", limit, bizHandler)

	if err := app.Listen("127.0.0.1:5678"); err != nil {
		slog.Error("fiber app start failed", "error", err)
	}
}
