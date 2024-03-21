package cache

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/util/gconv"
	"strings"
	"sync"
	"time"
)

var (
	ErrKeyNotFound = errors.New("key不存在")
	ErrKeyType     = errors.New("key存储类型错误")
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

func (b *BuildInMapCache) Set(ctx context.Context, key string, val string, expiration time.Duration) error {
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
	return nil
}

func (b *BuildInMapCache) Get(ctx context.Context, key string) (string, error) {
	b.mutex.RLock()
	res, ok := b.data[key]
	b.mutex.RUnlock()
	if !ok {
		return "", ErrKeyNotFound
	}
	now := time.Now()
	if res.deadlineBefore(now) {
		b.mutex.Lock()
		res, ok = b.data[key]
		defer b.mutex.Unlock()
		if res.deadlineBefore(now) {
			delete(b.data, key)
			return "", ErrKeyNotFound
		}
	}
	return res.val.(string), nil
}

func (b *BuildInMapCache) Del(ctx context.Context, keys ...string) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	for _, key := range keys {
		delete(b.data, key)
	}
	return nil
}
func (b *BuildInMapCache) HSet(ctx context.Context, key string, values ...any) error {
	var dl time.Time
	if len(values) != 2 {
		return errors.New("参数错误")
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
		return ErrKeyType
	}
	m[gconv.String(values[0])] = gconv.String(values[1])
	b.data[key] = item{
		val:      m,
		deadline: dl,
	}
	return nil
}

func (b *BuildInMapCache) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	var dl time.Time
	if expiration > 0 {
		dl = time.Now().Add(expiration)
	}
	b.mutex.Lock()
	defer b.mutex.Unlock()
	res, ok := b.data[key]
	if !ok {
		return false, ErrKeyNotFound
	}
	now := time.Now()
	if res.deadlineBefore(now) {
		return false, ErrKeyNotFound
	}
	b.data[key] = item{
		val:      res.val,
		deadline: dl,
	}
	return true, nil
}

func (b *BuildInMapCache) Exists(ctx context.Context, keys ...string) (int64, error) {
	if len(keys) != 1 {
		return 0, errors.New("参数错误")
	}
	b.mutex.RLock()
	res, ok := b.data[keys[0]]
	b.mutex.RUnlock()
	if !ok {
		return 0, nil
	}
	now := time.Now()
	if res.deadlineBefore(now) {
		b.mutex.Lock()
		res, ok = b.data[keys[0]]
		defer b.mutex.Unlock()
		if res.deadlineBefore(now) {
			delete(b.data, keys[0])
			return 0, nil
		}
	}
	return 1, nil
}

func (b *BuildInMapCache) HGet(ctx context.Context, key, field string) (string, error) {
	b.mutex.RLock()
	res, ok := b.data[key]
	b.mutex.RUnlock()
	if !ok {
		return "", ErrKeyNotFound
	}
	now := time.Now()
	if res.deadlineBefore(now) {
		b.mutex.Lock()
		res, ok = b.data[key]
		defer b.mutex.Unlock()
		if res.deadlineBefore(now) {
			delete(b.data, key)
			return "", ErrKeyNotFound
		}
	}
	m, ok := res.val.(map[string]string)
	if !ok {
		return "", ErrKeyType
	}
	return m[field], nil
}

func (b *BuildInMapCache) Scan(ctx context.Context, cursor uint64, match string, count int64) ([]string, uint64, error) {
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
	return ss, 0, nil
}

func (b *BuildInMapCache) JudgmentAndHSet(ctx context.Context, rk, key string, gs any) error {
	b.mutex.RLock()
	res, ok := b.data[rk]
	b.mutex.RUnlock()
	if !ok {
		return ErrKeyNotFound
	}
	b.mutex.Lock()
	defer b.mutex.Unlock()
	now := time.Now()
	if res.deadlineBefore(now) {
		res, ok = b.data[rk]
		if res.deadlineBefore(now) {
			delete(b.data, rk)
			return ErrKeyNotFound
		}
	}
	m, ok := res.val.(map[string]string)
	if !ok {
		return ErrKeyType
	}
	m[key] = gconv.String(gs)
	b.data[rk] = item{
		val:      m,
		deadline: res.deadline,
	}
	return nil
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
