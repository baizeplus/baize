package cache

import (
	"context"
	"errors"
	"sync"
	"time"
)

var (
	ErrKeyNotFound = errors.New("key不存在")
)

type BuildInMapCache struct {
	data  map[string]item
	mutex sync.RWMutex
	close chan struct{}
}

func NewBuildInMapCache(internal time.Duration) *BuildInMapCache {
	res := &BuildInMapCache{
		data: make(map[string]item, 128),
	}
	go func() {
		ticker := time.NewTicker(internal)
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

type item struct {
	val      any
	deadline time.Time
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

func (b *BuildInMapCache) Close() error {
	select {
	case b.close <- struct{}{}:
	default:
		return errors.New("重复关闭")
	}
	return nil
}

func (i *item) deadlineBefore(t time.Time) bool {
	return i.deadline.IsZero() && i.deadline.Before(t)
}
