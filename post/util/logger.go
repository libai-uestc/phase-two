package util

import (
	"log/slog"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

func InitSlog(logFile string) {
	fout, err := rotatelogs.New(
		logFile+".%Y%m%d%H",
		rotatelogs.WithLinkName(logFile),
		rotatelogs.WithRotationTime(1*time.Hour),
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)
	if err != nil {
		panic(err)
	}

	handler := slog.NewTextHandler(
		fout,
		&slog.HandlerOptions{
			AddSource: true,
			Level:     slog.LevelInfo,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey {
					t := a.Value.Time()
					a.Value = slog.StringValue(t.Format("2006-01-02 15:04:05.000"))
				}
				return a
			},
		},
	)
	logger := slog.New(handler)

	slog.SetDefault(logger)
}
