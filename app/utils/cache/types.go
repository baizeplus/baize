package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"time"
)

var cache Cache
var RedisClient *redis.Client

func GetCache() Cache {
	return cache
}
func init() {
	switch viper.GetString("cache.type") {
	case "redis":
		r := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", viper.GetString("cache.redis.host"), viper.GetInt("cache.redis.port")),
			Password: viper.GetString("cache.redis.password"),
			DB:       viper.GetInt("cache.redis.db"),
		})
		RedisClient = r
		cache = NewRedisCache(r)
	default:
		cache = NewBuildInMapCache()
	}

}

type Cache interface {
	Set(ctx context.Context, key string, val string, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) error
	HSet(ctx context.Context, key string, values ...any) error
	Expire(ctx context.Context, key string, expiration time.Duration) (bool, error)
	Exists(ctx context.Context, keys ...string) (int64, error)
	HGet(ctx context.Context, key, field string) (string, error)
	Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error)

	JudgmentAndHSet(ctx context.Context, rk, key string, gs any) error
}
