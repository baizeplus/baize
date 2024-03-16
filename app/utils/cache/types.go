package cache

import (
	"context"
	"time"
)

var cache Cache

func GetCache() Cache {
	return cache
}
func SetCache(c Cache) {
	cache = c
}

type Cache interface {
	Set(ctx context.Context, key string, val string, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) error
	HSet(ctx context.Context, key string, values ...any) error
	Expire(ctx context.Context, key string, expiration time.Duration) (bool, error)
	Exists(ctx context.Context, keys ...string) (int64, error)
	HGet(ctx context.Context, key, field string) (string, error)
	JudgmentAndHSet(ctx context.Context, ids, key string, gs any) error
}
