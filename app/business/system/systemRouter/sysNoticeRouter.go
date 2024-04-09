package systemRoutes

import (
	"baize/app/business/system/systemController"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysRouterRouter(router *gin.RouterGroup, dc *systemController.Notice) {
	systemNoticeData := router.Group("/system/notice")
	systemNoticeData.GET("/list", middlewares.HasPermission("system:notice:list"), dc.NoticeList)
	systemNoticeData.GET("/:id", middlewares.HasPermission("system:notice:query"), dc.NoticeGetInfo)
	systemNoticeData.POST("", middlewares.HasPermission("system:notice:add"), dc.NoticeAdd)
	systemConsumptionData := router.Group("/system/consumption")
	systemConsumptionData.GET("/newMessage", dc.NewMessage)
	systemConsumptionData.GET("/:id", dc.UserNoticeGetInfo)
	systemConsumptionData.GET("/userNoticeList", dc.UserNoticeList)
	systemConsumptionData.PUT("/noticeRead/:id", dc.NoticeRead)
	systemConsumptionData.PUT("/noticeReadAll", dc.NoticeReadAll)
	systemConsumptionData.DELETE("/noticeDelete/:ids", dc.NoticeDelete)

}
