package io_test

import (
	"libai/go/phase-two/io"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogrus(t *testing.T) {
	logger := io.InitLogrus("../log/logrus.log", "info")
	logger.Debug("this is debug log")
	logEntry := logger.WithFields(logrus.Fields{"name": "libai", "age": 18})
	logEntry.Info("this is info log")
	logEntry.Warnf("this is warn log, float=%.3f", 3.14)
	logger.Error("this is error log1", "this is error log2")

	// logger.Fatal("this is fatal log")                        //写完日志之后会调os.Exit(1)
	defer func() {
		recover()
	}()
	logger.Panic("this is panic log")
}
