package middlewares

import (
	"baize/app/baize"
	"baize/app/datasource/cache"
	"baize/app/middlewares/session"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type sessionAuthMiddlewareBuilder struct {
	paths baize.Set[string]
	cache cache.Cache
}

func NewSessionAuthMiddlewareBuilder(cache cache.Cache) *sessionAuthMiddlewareBuilder {
	return &sessionAuthMiddlewareBuilder{cache: cache, paths: baize.Set[string]{}}
}

func (s *sessionAuthMiddlewareBuilder) IgnorePaths(path string) *sessionAuthMiddlewareBuilder {
	s.paths.Add(path)
	return s
}

func (s *sessionAuthMiddlewareBuilder) Build() func(c *gin.Context) {
	return func(c *gin.Context) {
		manager := session.NewManger(s.cache)
		_, err := manager.GetSession(c)
		if err != nil {
			baizeContext.InvalidToken(c)
			c.Abort()
			return
		}
		if !s.paths.Contains(c.Request.RequestURI) {
			_ = manager.RefreshSession(c)
		}
		c.Next()
	}
}
