package io

import (
	"fmt"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitZap1(logFile string) *zap.Logger {
	rotateOut, err := rotatelogs.New(
		logFile+".%Y%m%d%H",
		rotatelogs.WithLinkName(logFile),
		rotatelogs.WithRotationTime(1*time.Hour),
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	// fmt.Println()

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(rotateOut),
		zapcore.InfoLevel,
	)
	logger := zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
		zap.Hooks(func(e zapcore.Entry) error {
			if e.Level >= zapcore.ErrorLevel {
				fmt.Println(e.Message)
			}
			return nil
		}),
	)
	logger = logger.With(
		zap.Namespace("libai"),
		zap.String("biz", "search"),
	)
	return logger
}

func InitZap2(logFile, level string) *zap.Logger {
	config := zap.Config{
		Encoding:         "console",
		OutputPaths:      []string{"stdout", logFile},
		ErrorOutputPaths: []string{"stderr"},
		InitialFields:    map[string]any{"biz": "search"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "msg",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,
		},
	}
	switch strings.ToLower(level) {
	case "debug":
		config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		config.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		config.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
	case "fatal":
		config.Level = zap.NewAtomicLevelAt(zap.FatalLevel)
	case "panic":
		config.Level = zap.NewAtomicLevelAt(zap.PanicLevel)
	default:
		panic(fmt.Errorf("invalid log level %s", level))
	}
	logger, _ := config.Build()
	logger = logger.With(
		zap.Namespace("libai"),
		zap.String("group", "game"),
	)
	return logger
}
