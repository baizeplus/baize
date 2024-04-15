package monitorRouter

import (
	"baize/app/business/monitor/monitorController"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysUserOnlineRouter(router *gin.RouterGroup, uoc *monitorController.UserOnline) {
	online := router.Group("/monitor/online")
	online.GET("/list", middlewares.HasPermission("monitor:online:list"), uoc.UserOnlineList)
	online.DELETE("/:tokenId", middlewares.SetLog("在线用户", middlewares.ForcedRetreat), middlewares.HasPermission("monitor:online:forceLogout"), uoc.ForceLogout)

}
