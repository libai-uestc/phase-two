package io_test

import (
	"log/slog"
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LOG      = "这是日志内容"
	TIME_FMT = "2006-01-02 15:04:05.000"
)

var (
	loc *time.Location
)

func init() {
	loc, _ = time.LoadLocation("asia/shanghai")
}

func BenchmarkZap(b *testing.B) {
	logFile := "../log/zap.log"
	fout, err := os.OpenFile(logFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(TIME_FMT)
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(fout),
		zapcore.InfoLevel,
	)
	logger := zap.New(
		core,
		zap.AddCaller(),
	)
	logger = logger.With(
		zap.Namespace("uber"),
		zap.String("bizID", "123456"),
	)

	b.ResetTimer()
	for b.Loop() {
		logger.Info(LOG, zap.String("name", "libai"), zap.Int("age", 18))
	}
	logger.Sync()

}

func BenchmarkSlog(b *testing.B) {
	logFile := "../log/slog.log"
	fout, err := os.OpenFile(logFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	handler := slog.NewJSONHandler(
		fout,
		&slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelInfo,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey {
					t := a.Value.Time()
					a.Value = slog.StringValue(t.Format(TIME_FMT))
				}
				return a
			},
		},
	)
	logger := slog.New(handler)

	logger = logger.WithGroup("google").With(slog.String("bizId", "123456"))

	b.ResetTimer()
	for b.Loop() {
		logger.Info(LOG, slog.String("name", "libai"), slog.Int("age", 18))
	}
}

func BenchmarkLogrus(b *testing.B) {
	logFile := "../log/logrus.log"
	fout, err := os.OpenFile(logFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	logger := logrus.New()
	logger.SetOutput(fout)
	logger.SetLevel(logrus.InfoLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: TIME_FMT,
	})
	logEntry := logger.WithField("bizId", "123456")
	b.ResetTimer()
	for b.Loop() {
		logEntry.WithFields(logrus.Fields{"name": "libai", "age": 18}).Info(LOG)
	}
}

// go test ./io -bench=^BenchmarkZap$ -run=^$
// go test ./io -bench=^BenchmarkSlog$ -run=^$
// go test ./io -bench=^BenchmarkLogrus$ -run=^$
