package systemService

import (
	"baize/app/business/system/systemModels"
	"github.com/gin-gonic/gin"
)

type ISseService interface {
	BuildNotificationChannel(c *gin.Context)
	SendNotification(userId []int64, ss *systemModels.Sse)
}
