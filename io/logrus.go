package io

import (
	"fmt"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

func InitLogrus(logFile, level string) *logrus.Logger {
	logger := logrus.New()
	switch strings.ToLower(level) {
	case "debug":
		logger.SetLevel(logrus.DebugLevel)
	case "info":
		logger.SetLevel(logrus.InfoLevel)
	case "warn":
		logger.SetLevel(logrus.WarnLevel)
	case "error":
		logger.SetLevel(logrus.ErrorLevel)
	case "fatal":
		logger.SetLevel(logrus.FatalLevel)
	case "panic":
		logger.SetLevel(logrus.PanicLevel)
	default:
		panic(fmt.Errorf("invalid log level %s", level))
	}
	// 普通文本格式
	logger.SetFormatter(&logrus.TextFormatter{
		// ForceColors: true,
		DisableColors:   true,
		TimestampFormat: "2006-01-02 15:04:05.000", // ms
	})
	// json格式
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})

	fout, err := rotatelogs.New(
		logFile+".%Y%m%d%H",
		rotatelogs.WithLinkName(logFile),
		rotatelogs.WithRotationTime(1*time.Hour),
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(fout)
	// logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)

	return logger

}

type AppHook struct {
	AppName string
}

func (h *AppHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel}
}

func (h *AppHook) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = h.AppName
	fmt.Println(entry.Message)
	return nil
}
