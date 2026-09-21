package distributed

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func TryLock(rc *redis.Client, key string, expire time.Duration) bool {
	cmd := rc.SetNX(context.Background(), key, "value", expire)
	if cmd.Err() != nil {
		return false
	} else {
		return cmd.Val()
	}
}

func ReleaseLock(rc *redis.Client, key string) {
	rc.Del(context.Background(), key)
}
