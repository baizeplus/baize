package session

import (
	"baize/app/utils/session/redis"
	"context"
	"github.com/gin-gonic/gin"
)

type Store interface {
	Generate(ctx context.Context, userId int64) (*redis.Session, error)
	Refresh(ctx context.Context, id string) error
	Remove(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*redis.Session, error)
}

type Session interface {
	Get(ctx context.Context, key string) string
	Set(ctx context.Context, key string, val any)
	Id() string
}

type Propagator interface {
	Extract(c *gin.Context) (string, error)
}
