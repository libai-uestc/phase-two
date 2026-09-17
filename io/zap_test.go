package io_test

import (
	"libai/go/phase-two/io"
	"testing"

	"go.uber.org/zap"
)

func TestZap1(t *testing.T) {
	logger := io.InitZap1("../log/zap1.log")
	defer logger.Sync()
	logger.Debug("hello")
	logger.Info("hello", zap.Int("age", 18))
	logger.Error("hello", zap.Namespace("china"), zap.Int("age", 18))

	sugar := logger.Sugar()
	sugar.Infof("pi is %f", 3.1415926)
}

func TestZap2(t *testing.T) {
	logger := io.InitZap2("../log/zap2.log", "info")
	defer logger.Sync() //将缓存的内容同步到文件中
	logger.Debug("hello")
	logger.Info("hello", zap.Int("age", 18))
	logger.Error("hello", zap.Namespace("china"), zap.Int("age", 18))
}
