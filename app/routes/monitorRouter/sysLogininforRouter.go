package monitorRouter

import (
	"baize/app/bzMonitor/controller"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysLogininforRouter(router *gin.RouterGroup, lc *controller.Logininfor) {
	logininfor := router.Group("/monitor/logininfor")
	logininfor.GET("/list", middlewares.HasPermission("monitor:logininfor:list"), lc.LogininforList)
	logininfor.GET("/export", middlewares.HasPermission("monitor:logininfor:list"), lc.LogininforExport)
	logininfor.DELETE("/:infoIds", middlewares.HasPermission("monitor:logininfor:remove"), lc.LogininforRemove)
	logininfor.DELETE("/clean", middlewares.HasPermission("monitor:logininfor:remove"), lc.LogininforClean)
}
