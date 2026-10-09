package xorm_test

import (
	"libai/go/phase-two/orm/xorm"
	"log/slog"
	"os"
	"sync"
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

func TestUpdate(t *testing.T) {
	xorm.Update(engine)
}

// go test -v ./orm/xorm -run=^TestUpdate$ -count=1

func TestUpdateByVersion(t *testing.T) {
	const P = 10
	wg := sync.WaitGroup{}
	wg.Add(P)
	for i := 0; i < P; i++ {
		go func() {
			defer wg.Done()
			xorm.UpdateByVersion(engine)
		}()
	}
	wg.Wait()
}

// go test -v ./orm/xorm -run=^TestUpdateByVersion$ -count=1

func TestRead(t *testing.T) {
	xorm.Read(engine)
}

// go test -v ./orm/xorm -run=^TestRead$ -count=1
func TestReadWithStatistics(t *testing.T) {
	xorm.ReadWithStatistics(engine)
}

// go test -v ./orm/xorm -run=^TestReadWithStatistics$ -count=1

func TestTransaction(t *testing.T) {
	xorm.Transaction(engine)
}

// go test -v ./orm/xorm -run=^TestTransaction$ -count=1

func TestHandleError(t *testing.T) {

}
