package monitorRouter

import (
	"baize/app/bzMonitor/controller"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitServerRouter(router *gin.RouterGroup, sc *controller.InfoServer) {
	server := router.Group("/monitor/server")
	server.GET("", middlewares.HasPermission("monitor:server:list"), sc.GetInfoServer)

}
