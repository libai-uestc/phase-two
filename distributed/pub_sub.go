package distributed

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type CS string

func publish(ctx context.Context, client *redis.Client, channel string, message any) {
	cmd := client.Publish(context.Background(), channel, message)
	if cmd.Err() == nil {
		n := cmd.Val()
		fmt.Printf("%s向频道%s发布了消息，此时频道也%d个订阅者\n", ctx.Value("publisher_name"), channel, n)
	} else {
		fmt.Printf("%s向频道%s发布消息失败%v\n", ctx.Value("publisher_name"), channel, cmd.Err())
	}
}

func subscribe(ctx context.Context, client *redis.Client, channels []string) {
	ps := client.Subscribe(context.Background(), channels...)
	defer ps.Close()

	for {
		if msg, err := ps.ReceiveMessage(context.Background()); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Printf("%s从频道%s里接收到消息:%s\n", ctx.Value("subscriber_name"), msg.Channel, msg.Payload)
		}
	}
}

func PubSub(ctx context.Context, client *redis.Client) {
	ctx1 := context.WithValue(ctx, CS("publisher_name"), "publisher1")
	ctx2 := context.WithValue(ctx, CS("publisher_name"), "publisher2")
	channel1 := "channel1"
	channel2 := "channel2"

	ctx3 := context.WithValue(ctx, CS("subscriber_name"), "subscriber3")
	ctx4 := context.WithValue(ctx, CS("subscriber_name"), "subscriber4")

	go subscribe(ctx3, client, []string{channel1})
	go subscribe(ctx4, client, []string{channel2})
	time.Sleep(1 * time.Second)

	go publish(ctx3, client, channel1, "黄河之水天上来")
	go publish(ctx2, client, channel1, "奔流到海不复回")
	time.Sleep(1 * time.Second)
	fmt.Println(strings.Repeat("-", 50))

	ctx5 := context.WithValue(ctx, CS("subscriber_name"), "subscriber5")
	go subscribe(ctx5, client, []string{channel1, channel2})
	time.Sleep(1 * time.Second)

	go publish(ctx1, client, channel2, "天生我材必有用")
	go publish(ctx2, client, channel2, "千金散尽还复来")
	time.Sleep(1 * time.Second)
}
