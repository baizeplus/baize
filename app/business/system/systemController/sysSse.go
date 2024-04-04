package systemController

import (
	"baize/app/business/system/systemService"
	"baize/app/business/system/systemService/systemServiceImpl"
	"github.com/gin-gonic/gin"
)

type Sse struct {
	ss systemService.ISseService
}

func NewSse(ss *systemServiceImpl.SseService) *Sse {
	return &Sse{ss: ss}
}

// BuildSse 建立SSE链接
// @Summary 建立SSE链接
// @Description 建立SSE链接
// @Tags 建立SSE链接
// @Security BearerAuth
// @Router /system/sse  [get]
func (s *Sse) BuildSse(c *gin.Context) {
	s.ss.BuildNotificationChannel(c)
}
