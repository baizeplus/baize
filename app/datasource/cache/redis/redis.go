package redis

import (
	"baize/app/datasource/cache/cacheError"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	lua = `if redis.call("exists", KEYS[1])
							then 
								return redis.call("hset", KEYS[1], ARGV[1],ARGV[2])
							else
								return -1
							end`
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

//func (r *RedisCache) Publish(ctx context.Context, channel string, message interface{}) error {
//	return r.client.Publish(ctx, channel, message).Err()
//}
//
//func (r *RedisCache) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
//	return r.client.Subscribe(ctx, channels...).Channel()
//}

func (r *RedisCache) Set(ctx context.Context, key string, val string, expiration time.Duration) {
	_, err := r.client.Set(ctx, key, val, expiration).Result()
	if err != nil {
		panic(err)
	}
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := r.client.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", cacheError.Nil
		}
		return "", err
	}
	return result, nil
}

func (r *RedisCache) Del(ctx context.Context, keys ...string) {
	_, err := r.client.Del(ctx, keys...).Result()
	if err != nil {
		panic(err)
	}
}
func (r *RedisCache) HSet(ctx context.Context, key string, values ...any) {
	_, err := r.client.HSet(ctx, key, values...).Result()
	if err != nil {
		panic(err)
	}
}
func (r *RedisCache) Expire(ctx context.Context, key string, expiration time.Duration) bool {
	result, err := r.client.Expire(ctx, key, expiration).Result()
	if err != nil {
		panic(err)
	}
	return result
}

func (r *RedisCache) Exists(ctx context.Context, keys ...string) int64 {
	result, err := r.client.Exists(ctx, keys...).Result()
	if err != nil {
		panic(err)
	}
	return result
}
func (r *RedisCache) HGet(ctx context.Context, key, field string) string {
	result, err := r.client.HGet(ctx, key, field).Result()
	if err != nil {
		panic(err)
	}
	return result
}
func (r *RedisCache) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64) {
	keys, c, err := r.client.Scan(ctx, cursor, match, count).Result()
	if err != nil {
		panic(err)
	}
	return keys, c
}

func (r *RedisCache) JudgmentAndHSet(ctx context.Context, ids, key string, gs any) {
	_, err := r.client.Eval(ctx, lua, []string{ids}, key, gs).Int()
	if err != nil {
		panic(err)
	}
}

func (r *RedisCache) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) bool {
	result, err := r.client.SetNX(ctx, key, value, expiration).Result()
	if err != nil {
		panic(err)
	}
	return result
}

func (r *RedisCache) Publish(ctx context.Context, channel string, message interface{}) {
	err := r.client.Publish(ctx, channel, message).Err()
	if err != nil {
		panic(err)
	}
}

func (r *RedisCache) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	return r.client.Subscribe(ctx, channels...)
}
