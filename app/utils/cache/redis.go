package cache

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"time"
)

var (
	ErrFailedToSetCache = errors.New("写入redis失败")
	lua                 = `if redis.call("exists", KEYS[1])
							then 
								return redis.call("hset", KEYS[1], ARGV[1],ARGV[2])
							else
								return -1
							end`
)

type RedisCache struct {
	client redis.Cmdable
	//client *redis.Client
}

func NewRedisCache(client redis.Cmdable) *RedisCache {
	//func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}

}

//func (r *RedisCache) Publish(ctx context.Context, channel string, message interface{}) error {
//	return r.client.Publish(ctx, channel, message).Err()
//}
//
//func (r *RedisCache) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
//	return r.client.Subscribe(ctx, channels...).Channel()
//}

func (r *RedisCache) Set(ctx context.Context, key string, val string, expiration time.Duration) error {
	result, err := r.client.Set(ctx, key, val, expiration).Result()
	if err != nil {
		return err
	}
	if result != "OK" {
		return ErrFailedToSetCache
	}
	return nil
}

func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RedisCache) Del(ctx context.Context, keys ...string) error {
	_, err := r.client.Del(ctx, keys...).Result()
	return err
}
func (r *RedisCache) HSet(ctx context.Context, key string, values ...any) error {
	_, err := r.client.HSet(ctx, key, values...).Result()
	return err
}
func (r *RedisCache) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	return r.client.Expire(ctx, key, expiration).Result()
}

func (r *RedisCache) Exists(ctx context.Context, keys ...string) (int64, error) {
	return r.client.Exists(ctx, keys...).Result()
}
func (r *RedisCache) HGet(ctx context.Context, key, field string) (string, error) {
	return r.client.HGet(ctx, key, field).Result()
}
func (r *RedisCache) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error) {
	return r.client.Scan(ctx, cursor, match, count).Result()
}

func (r *RedisCache) JudgmentAndHSet(ctx context.Context, ids, key string, gs any) error {
	_, err := r.client.Eval(ctx, lua, []string{ids}, key, gs).Int()
	return err
}
