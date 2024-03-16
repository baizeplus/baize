package sessionCache

import (
	"baize/app/constant/sessionStatus"
	"baize/app/utils/cache"
	"baize/app/utils/stringUtils"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/spf13/viper"
	"gopkg.in/errgo.v2/errors"
	"time"
)

var (
	SessionKey         = `session_key`
	ErrSessionNotFound = errors.New("session:id 对应的session不存在")
)

type Store struct {
	expiration time.Duration
}

func NewStore() *Store {
	return &Store{
		expiration: time.Duration(viper.GetInt("token.expire_time")) * time.Minute,
	}
}

func (s *Store) Generate(ctx context.Context, userId int64) (*Session, error) {
	sId := sessionId(userId)
	err := cache.GetCache().HSet(ctx, redisKey(sId), sessionStatus.UserId, userId)
	if err != nil {
		return nil, err
	}

	return NewSession(sId), nil
}

func (s *Store) Refresh(ctx context.Context, id string) error {
	ok, err := cache.GetCache().Expire(ctx, redisKey(id), s.expiration)
	if err != nil {
		return err
	}
	if !ok {
		return ErrSessionNotFound
	}
	return nil
}

func (s *Store) Remove(ctx context.Context, id string) error {
	return cache.GetCache().Del(ctx, redisKey(id))

}

func (s *Store) Get(ctx context.Context, id string) (*Session, error) {
	cnt, err := cache.GetCache().Exists(ctx, redisKey(id))
	if err != nil {
		return nil, err
	}
	if cnt != 1 {
		return nil, ErrSessionNotFound
	}
	return NewSession(id), nil
}

func NewSession(id string) *Session {
	return &Session{
		id:     id,
		values: make(map[string]string),
	}
}

type Session struct {
	id     string
	values map[string]string
}

func (s *Session) Get(ctx context.Context, key string) string {
	val := s.values[key]
	if val != "" {
		return val
	}
	result, err := cache.GetCache().HGet(ctx, redisKey(s.id), key)
	if err != nil {
		return ""
	}
	s.values[key] = result
	return result

}

func (s *Session) Set(ctx context.Context, key string, val any) {
	gs := gconv.String(val)
	s.values[key] = gs
	_ = cache.GetCache().JudgmentAndHSet(ctx, redisKey(s.id), key, gs)
}

func (s *Session) Id() string {
	return s.id
}

func sessionId(userId int64) string {
	return fmt.Sprintf("%d:%s", userId, stringUtils.GetUUID())
}

func redisKey(id string) string {
	return fmt.Sprintf("%s:%s", SessionKey, id)
}
