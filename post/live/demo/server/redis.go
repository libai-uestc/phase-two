package main

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	client    *redis.Client
	redisOnce sync.Once
)

func GetRedisClient() *redis.Client {
	redisOnce.Do(func() {
		client = redis.NewClient(
			&redis.Options{
				Addr: "localhost:6379",
				DB:   0,
			})
		if err := client.Ping(context.Background()).Err(); err != nil {
			panic(err)
		}
	})

	return client
}
