package monitorRouter

import (
	"baize/app/business/monitor/monitorController"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysLogininforRouter(router *gin.RouterGroup, lc *monitorController.Logininfor) {
	logininfor := router.Group("/monitor/logininfor")
	logininfor.GET("/list", middlewares.HasPermission("monitor:logininfor:list"), lc.LogininforList)
	logininfor.GET("/export", middlewares.HasPermission("monitor:logininfor:list"), lc.LogininforExport)
	logininfor.DELETE("/:infoIds", middlewares.SetLog("登录日志", middlewares.Delete), middlewares.HasPermission("monitor:logininfor:remove"), lc.LogininforRemove)
	logininfor.DELETE("/clean", middlewares.SetLog("登录日志", middlewares.Clear), middlewares.HasPermission("monitor:logininfor:remove"), lc.LogininforClean)
}
