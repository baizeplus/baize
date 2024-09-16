package systemService

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/gin-gonic/gin"
)

type ISseService interface {
	BuildNotificationChannel(c *gin.Context)
	SendNotification(c context.Context, ss *systemModels.Sse)
}
