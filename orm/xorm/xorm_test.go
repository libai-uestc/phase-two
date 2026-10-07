package xorm_test

import (
	"libai/go/phase-two/orm/xorm"
	"log/slog"
	"os"
	"testing"
)

var (
	engine = xorm.CreateEngine("localhost", "test", "tester", "123456", 3306)
)

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

func TestCreate(t *testing.T) {
	xorm.Create(engine)
}

// go test -v ./orm/xorm -run=^TestCreate$ -count=1

func TestDelete(t *testing.T) {
	xorm.Delete(engine)
}

// go test -v ./orm/xorm -run=^TestDelete$ -count=1
func TestHandleError(t *testing.T) {

}
