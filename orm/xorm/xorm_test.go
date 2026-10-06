package xorm_test

import (
	"libai/go/phase-two/orm/xorm"
	"log/slog"
	"os"
	"testing"
)

// var (
// 	engine = xorm.
// )

func init() {
	slog.SetDefault(
		slog.New(
			slog.NewTextHandler(os.Stdout,
				&slog.HandlerOptions{
					AddSource: true,
				},
			),
		),
	)
}

func TestXormQuickStart(t *testing.T) {
	xorm.XormQuickStart()
}

// go test -v ./orm/xorm -run=^TestXormQuickStart$ -count=1
