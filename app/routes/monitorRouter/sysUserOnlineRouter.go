package monitorRouter

import (
	"baize/app/bzMonitor/controller"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysUserOnlineRouter(router *gin.RouterGroup, uoc *controller.UserOnline) {
	online := router.Group("/monitor/online")
	online.GET("/list", middlewares.HasPermission("monitor:online:list"), uoc.UserOnlineList)
	online.DELETE("/:tokenId", middlewares.HasPermission("monitor:online:forceLogout"), uoc.ForceLogout)

}
