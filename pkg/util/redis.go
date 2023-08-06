package util

import (
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/setting"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client
	Redis       = redisWrapper{}
)

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", setting.Config.Redis.Host, setting.Config.Redis.Port),
		Password: setting.Config.Redis.Password,
	})
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Default.Error("redis connect fail: " + err.Error())
		panic(err)
	}
}

type redisWrapper struct{}

func (redisWrapper) Set(key string, value interface{}, expiration time.Duration) error {
	return redisClient.Set(ctx, key, value, expiration).Err()
}

func (redisWrapper) Get(key string) (string, error) {
	return redisClient.Get(ctx, key).Result()
}

func (redisWrapper) Del(key string) error {
	return redisClient.Del(ctx, key).Err()
}

func (redisWrapper) Keys(pattern string) ([]string, error) {
	return redisClient.Keys(ctx, pattern).Result()
}

func (redisWrapper) PFAdd(key string, els ...interface{}) error {
	return redisClient.PFAdd(ctx, key, els...).Err()
}

func (redisWrapper) PFCount(keys ...string) (int64, error) {
	return redisClient.PFCount(ctx, keys...).Result()
}
