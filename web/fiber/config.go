// package main

// import (
// 	"context"
// 	"os"
// 	"time"

// 	"github.com/bytedance/sonic"
// 	"github.com/gofiber/fiber/v3"
// 	fiberlog "github.com/gofiber/fiber/v3/log"
// )

// func main() {
// 	logger := fiberlog.DefaultLogger[context.Context]()
// 	logger.SetLevel(fiberlog.LevelInfo)
// 	logFile, _ := os.OpenFile("C:/Users/18101/Desktop/第二阶段/phase-two/log/fiber.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
// 	logger.SetOutput(logFile)
// 	logger.Info("start fiber")

// 	app := fiber.New(
// 		fiber.Config{
// 			AppName: "libai",
// 			ErrorHandler: func(c fiber.Ctx, err error) error {
// 				logger.Error(err)
// 				return nil
// 			},
// 			ServerHeader:      "Fiber",
// 			BodyLimit:         4 << 20,
// 			StreamRequestBody: false,
// 			Concurrency:       256 << 10,
// 			DisableKeepalive:  false,
// 			IdleTimeout:       time.Minute,
// 			ReadTimeout:       time.Minute,
// 			WriteTimeout:      time.Minute,
// 			ReadBufferSize:    4 << 10,
// 			WriteBufferSize:   4 << 10,
// 			JSONEncoder:       sonic.Marshal,
// 			JSONDecoder:       sonic.Unmarshal,
// 		},
// 	)
// 	app.Hooks().OnRoute(func(r fiber.Route) error {
// 		logger.Info("route path %s method %s", r.Path, r.Method)
// 		return nil
// 	})

// 	app.Get("/home", func(c fiber.Ctx) error {
// 		return c.SendString("这是一个安全的网站")
// 	})

//		if err := app.Listen("127.0.0.1:5678", fiber.ListenConfig{CertFile: "data/server.crt", CertKeyFile: "data/rsa_private_key.pem"}); err != nil {
//			logger.Errorf("fiber app start failed:%s", err)
//		}
//	}
package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	fiberlog "github.com/gofiber/fiber/v3/log"
)

func main1() {
	// 1. 确保日志目录存在并安全打开文件
	logDir := `C:\Users\18101\Desktop\第二阶段\phase-two\log`
	_ = os.MkdirAll(logDir, 0755)

	logFile, err := os.OpenFile(filepath.Join(logDir, "fiber.log"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		panic("无法创建/打开日志文件: " + err.Error())
	}
	defer logFile.Close()

	// 2. 直接使用 fiberlog 的包级函数（内部直接操作默认 Logger，不会报 nil）
	fiberlog.SetLevel(fiberlog.LevelInfo)
	fiberlog.SetOutput(logFile)
	fiberlog.Info("start fiber")

	// 3. 构建 Fiber 实例
	app := fiber.New(
		fiber.Config{
			AppName: "libai",
			ErrorHandler: func(c fiber.Ctx, err error) error {
				fiberlog.Error(err)
				return nil
			},
			ServerHeader:      "Fiber",
			BodyLimit:         4 << 20,
			StreamRequestBody: false,
			Concurrency:       256 << 10,
			DisableKeepalive:  false,
			IdleTimeout:       time.Minute,
			ReadTimeout:       time.Minute,
			WriteTimeout:      time.Minute,
			ReadBufferSize:    4 << 10,
			WriteBufferSize:   4 << 10,
			JSONEncoder:       sonic.Marshal,
			JSONDecoder:       sonic.Unmarshal,
		},
	)

	app.Hooks().OnRoute(func(r fiber.Route) error {
		fiberlog.Infof("route path %s method %s", r.Path, r.Method)
		return nil
	})

	app.Get("/home", func(c fiber.Ctx) error {
		return c.SendString("这是一个安全的网站")
	})

	if err := app.Listen("127.0.0.1:5678", fiber.ListenConfig{CertFile: "data/server.crt", CertKeyFile: "data/rsa_private_key.pem"}); err != nil {
		fiberlog.Errorf("fiber app start failed:%s", err)
	}
}
