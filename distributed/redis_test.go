package distributed_test

import (
	"context"
	"libai/go/phase-two/distributed"
	"log/slog"
	"testing"

	"github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		DB:       0,
		Password: "",
	})
	// 修改点 1：使用 context.Background() 替代类型引用 context.Context
	if err := client.Ping(context.Background()).Err(); err != nil {
		slog.Error("connect to redis failed", "error", err)
	} else {
		slog.Info("connect to redis")
	}
}

func TestStringValue(t *testing.T) {
	distributed.StringValue(context.Background(), client)
}

func TestStructValue(t *testing.T) {
	stu := &distributed.Student{Id: 1, Name: "李白"}
	distributed.WriteStudent2Redis(client, stu)
	stu2 := distributed.GetStudentFromRedis(client, 1)

	// 修改点 2：增加 nil 判断，避免直接获取 stu2.Id 导致空指针 panic
	if stu2 == nil {
		t.Fatalf("expected student to not be nil")
	}

	if stu2.Id != stu.Id {
		t.Fail()
	}
	if stu2.Name != stu.Name {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	distributed.DeleteKey(context.Background(), client)
}

func TestScan(t *testing.T) {
	distributed.Scan(context.Background(), client)
}

func TestListValue(t *testing.T) {
	distributed.ListValue(context.Background(), client)
}

// 修正了方法名拼写错误 SetgValue -> SetValue
func TestSetValue(t *testing.T) {
	distributed.SetValue(context.Background(), client)
}

func TestZSetValue(t *testing.T) {
	distributed.ZsetValue(context.Background(), client)
}

func TestHashTableValue(t *testing.T) {
	distributed.HashtableValue(context.Background(), client)
}

// go test -v ./distributed -run=^TestStringValue$ -count=1
// go test -v ./distributed -run=^TestStructValue$ -count=1
// go test -v ./distributed -run=^TestDelete$ -count=1
// go test -v ./distributed -run=^TestListValue$ -count=1
// go test -v ./distributed -run=^TestSetValue$ -count=1
// go test -v ./distributed -run=^TestZSetValue$ -count=1
// go test -v ./distributed -run=^TestHashTableValue$ -count=1
// go test -v ./distributed -run=^TestScan$ -count=1
