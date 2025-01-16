package systemController

import (
	"baize/app/business/system/systemService"
	"github.com/gin-gonic/gin"
)

type Sse struct {
	ss systemService.ISseService
}

func NewSse(ss systemService.ISseService) *Sse {
	return &Sse{ss: ss}
}
func (s *Sse) PublicRoutes(router *gin.RouterGroup) {
	systemUser := router.Group("/system")
	systemUser.GET("/sse/:token", s.BuildSse)
}

// BuildSse 建立SSE链接
// @Summary 建立SSE链接
// @Description 建立SSE链接
// @Tags 建立SSE链接
// @Param  token path string true "token"
// @Security BearerAuth
// @Router /system/sse/{token}  [get]
func (s *Sse) BuildSse(c *gin.Context) {
	s.ss.BuildNotificationChannel(c)
}
