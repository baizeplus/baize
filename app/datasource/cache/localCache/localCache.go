package localCache

import (
	"baize/app/datasource/cache/cacheError"
	"baize/app/utils/converts"
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"strings"
	"sync"
	"time"
)

type BuildInMapCache struct {
	data  map[string]item
	mutex sync.RWMutex
	close chan struct{}
}

func NewBuildInMapCache() *BuildInMapCache {
	res := &BuildInMapCache{
		data: make(map[string]item, 128),
	}
	go func() {
		ticker := time.NewTicker(time.Minute)
		i := 0
		for {
			select {
			case t := <-ticker.C:
				res.mutex.Lock()
				for key, val := range res.data {
					if i > 1000 {
						break
					}
					if val.deadlineBefore(t) {
						delete(res.data, key)
					}
					i++
				}
				res.mutex.Unlock()
			case <-res.close:
				return
			}
		}
	}()
	return res
}

func (b *BuildInMapCache) Set(ctx context.Context, key string, val string, expiration time.Duration) {
	var dl time.Time
	if expiration > 0 {
		dl = time.Now().Add(expiration)
	}
	b.mutex.Lock()
	defer b.mutex.Unlock()
	b.data[key] = item{
		val:      val,
		deadline: dl,
	}
}

func (b *BuildInMapCache) Get(ctx context.Context, key string) (string, error) {
	b.mutex.RLock()
	res, ok := b.data[key]
	b.mutex.RUnlock()
	if !ok {
		return "", cacheError.Nil
	}
	now := time.Now()
	if res.deadlineBefore(now) {
		b.mutex.Lock()
		res, ok = b.data[key]
		defer b.mutex.Unlock()
		if res.deadlineBefore(now) {
			delete(b.data, key)
			return "", cacheError.Nil
		}
	}
	return res.val.(string), nil
}

func (b *BuildInMapCache) Del(ctx context.Context, keys ...string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	for _, key := range keys {
		delete(b.data, key)
	}
}
func (b *BuildInMapCache) HSet(ctx context.Context, key string, values ...any) {
	var dl time.Time
	if len(values) != 2 {
		panic("参数错误")
	}
	b.mutex.Lock()
	defer b.mutex.Unlock()
	res, ok := b.data[key]
	if !ok {
		res = item{
			val:      make(map[string]string),
			deadline: dl,
		}
	}
	m, ok := res.val.(map[string]string)
	if !ok {
		panic("key已被使用并类型不对")
	}
	m[converts.String(values[0])] = converts.String(values[1])
	b.data[key] = item{
		val:      m,
		deadline: dl,
	}
}

func (b *BuildInMapCache) Expire(ctx context.Context, key string, expiration time.Duration) bool {
	var dl time.Time
	if expiration > 0 {
		dl = time.Now().Add(expiration)
	}
	b.mutex.Lock()
	defer b.mutex.Unlock()
	res, ok := b.data[key]
	if !ok {
		return false
	}
	now := time.Now()
	if res.deadlineBefore(now) {
		return false
	}
	b.data[key] = item{
		val:      res.val,
		deadline: dl,
	}
	return true
}

func (b *BuildInMapCache) Exists(ctx context.Context, keys ...string) int64 {
	if len(keys) != 1 {
		return 0
	}
	b.mutex.RLock()
	res, ok := b.data[keys[0]]
	b.mutex.RUnlock()
	if !ok {
		return 0
	}
	now := time.Now()
	if res.deadlineBefore(now) {
		b.mutex.Lock()
		res, ok = b.data[keys[0]]
		defer b.mutex.Unlock()
		if res.deadlineBefore(now) {
			delete(b.data, keys[0])
			return 0
		}
	}
	return 1
}

func (b *BuildInMapCache) HGet(ctx context.Context, key, field string) string {
	b.mutex.RLock()
	res, ok := b.data[key]
	b.mutex.RUnlock()
	if !ok {
		return ""
	}
	now := time.Now()
	if res.deadlineBefore(now) {
		b.mutex.Lock()
		res, ok = b.data[key]
		defer b.mutex.Unlock()
		if res.deadlineBefore(now) {
			delete(b.data, key)
			return ""
		}
	}
	m, ok := res.val.(map[string]string)
	if !ok {
		return ""
	}
	return m[field]
}

func (b *BuildInMapCache) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64) {
	ss := make([]string, 0)
	suffix := strings.TrimSuffix(match, "*")
	now := time.Now()
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	for s, v := range b.data {
		if strings.HasPrefix(s, suffix) && !v.deadlineBefore(now) {
			ss = append(ss, s)
		}
	}
	return ss, 0
}

func (b *BuildInMapCache) JudgmentAndHSet(ctx context.Context, rk, key string, gs any) {
	b.mutex.RLock()
	res, ok := b.data[rk]
	b.mutex.RUnlock()
	if !ok {
		return
	}
	b.mutex.Lock()
	defer b.mutex.Unlock()
	now := time.Now()
	if res.deadlineBefore(now) {
		res, ok = b.data[rk]
		if res.deadlineBefore(now) {
			delete(b.data, rk)
			return
		}
	}
	m, ok := res.val.(map[string]string)
	if !ok {
		return
	}
	m[key] = converts.String(gs)
	b.data[rk] = item{
		val:      m,
		deadline: res.deadline,
	}
	return
}
func (b *BuildInMapCache) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) bool {
	panic("local 不支持请切换redis")
}
func (b *BuildInMapCache) Publish(ctx context.Context, channel string, message interface{}) {
	panic("local 不支持请切换redis")
}
func (b *BuildInMapCache) Subscribe(ctx context.Context, channels ...string) *redis.PubSub {
	panic("local 不支持请切换redis")
}
func (b *BuildInMapCache) Close() error {
	select {
	case b.close <- struct{}{}:
	default:
		return errors.New("重复关闭")
	}
	return nil
}

type item struct {
	val      any
	deadline time.Time
}

func (i *item) deadlineBefore(t time.Time) bool {
	return (!i.deadline.IsZero()) && i.deadline.Before(t)
}
