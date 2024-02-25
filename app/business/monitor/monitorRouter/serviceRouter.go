package monitorRouter

import (
	"baize/app/business/monitor/monitorController"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitServerRouter(router *gin.RouterGroup, sc *monitorController.InfoServer) {
	server := router.Group("/monitor/server")
	server.GET("", middlewares.HasPermission("monitor:server:list"), sc.GetInfoServer)

}
