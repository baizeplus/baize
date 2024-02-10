package middlewares

import (
	"baize/app/baize"
	"baize/app/utils/baizeContext"
	"baize/app/utils/session"
	"github.com/gin-gonic/gin"
)

// SessionAuthMiddleware 基于Session的认证中间件
func SessionAuthMiddleware(noRefresh baize.Set[string]) func(c *gin.Context) {
	return func(c *gin.Context) {
		manager := session.NewManger()
		_, err := manager.GetSession(c)
		if err != nil {
			baizeContext.InvalidToken(c)
			c.Abort()
			return
		}
		if !noRefresh.Contains(c.Request.RequestURI) {
			_ = manager.RefreshSession(c)
		}
		c.Next()
	}
}
