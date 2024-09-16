package systemController

import (
	"baize/app/business/system/systemService"
	"baize/app/business/system/systemService/systemServiceImpl"
	"baize/app/constant/sessionStatus"
	"baize/app/utils/baizeContext"
	"baize/app/utils/session"
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
// @Param  token path string true "token"
// @Security BearerAuth
// @Router /system/sse/{token}  [get]
func (s *Sse) BuildSse(c *gin.Context) {
	manager := session.NewManger()
	sess, err := manager.Get(c, c.Param("token"))
	if err != nil {
		baizeContext.InvalidToken(c)
	}
	c.Set(sessionStatus.SessionKey, sess)
	s.ss.BuildNotificationChannel(c)
}
