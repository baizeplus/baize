package cache

import (
	"baize/app/datasource/cache/localCache"
	redis2 "baize/app/datasource/cache/redis"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

func NewCache() (cache Cache) {
	switch viper.GetString("cache.type") {
	case "redis":
		r := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", viper.GetString("cache.redis.host"), viper.GetInt("cache.redis.port")),
			Password: viper.GetString("cache.redis.password"),
			DB:       viper.GetInt("cache.redis.db"),
		})

		cache = redis2.NewRedisCache(r)
	default:
		cache = localCache.NewBuildInMapCache()
	}
	return
}

type Cache interface {
	Set(ctx context.Context, key string, val string, expiration time.Duration)
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string)
	HSet(ctx context.Context, key string, values ...any)
	Expire(ctx context.Context, key string, expiration time.Duration) bool
	Exists(ctx context.Context, keys ...string) int64
	HGet(ctx context.Context, key, field string) string
	Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64)
	JudgmentAndHSet(ctx context.Context, rk, key string, gs any)
	SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) bool
	Publish(ctx context.Context, channel string, message interface{})
	Subscribe(ctx context.Context, channels ...string) *redis.PubSub
}
