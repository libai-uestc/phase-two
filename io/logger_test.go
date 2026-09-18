package io_test

import (
	"libai/go/phase-two/io"
	"log/slog"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitSlog() *slog.Logger {
	logFile := "../log/slog.log"
	fout, err := os.OpenFile(logFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logger := slog.New(
		&io.SlogContextHandler{
			slog.NewJSONHandler(fout, &slog.HandlerOptions{
				AddSource: true,
				Level:     slog.LevelInfo,
				ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
					if a.Key != slog.TimeKey {
						return a
					}
					t := a.Value.Time()
					a.Value = slog.StringValue(t.Format(TIME_FMT))

					return a
				},
			}),
		},
	)
	return logger
}

func InitLogrus() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: TIME_FMT,
	})

	logFile := "../log/logrus.log"
	fout, err := os.OpenFile(logFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	logger.SetOutput(fout)
	logger.SetReportCaller(true)
	return logger
}

func InitZap() *zap.Logger {
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
	return logger
}

func BenchmarkLoggerSlog(b *testing.B) {
	logger := InitSlog()
	b.ResetTimer()
	for b.Loop() {
		logger.Info(LOG, "name", "libai", "age", 18)
	}
}

func BenchmarkLoggerLogrus(b *testing.B) {
	logger := InitLogrus()
	b.ResetTimer()
	for b.Loop() {
		logger.WithFields(logrus.Fields{"name": "libai", "age": 18}).Info(LOG)
	}
}

func BenchmarkLoggerZap(b *testing.B) {
	logger := InitZap()
	b.ResetTimer()
	for b.Loop() {
		logger.Info(LOG, zap.String("name", "libai"), zap.Int("age", 18))
	}
	logger.Sync()
}

// go test ./io -bench=^BenchmarkLoggerSlog$ -run=^$
// go test ./io -bench=^BenchmarkLoggerLogrus$ -run=^$
// go test ./io -bench=^BenchmarkLoggerZap$ -run=^$
// go test ./io -bench=^BenchmarkLogger -run=^$
