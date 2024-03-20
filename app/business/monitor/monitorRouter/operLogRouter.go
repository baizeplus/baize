package monitorRouter

import (
	"baize/app/business/monitor/monitorController"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysOperLogRouter(router *gin.RouterGroup, sc *monitorController.OperLog) {
	operlog := router.Group("/monitor/operlog")
	operlog.GET("/list", middlewares.HasPermission("monitor:operlog:list"), sc.OperLogList)
	operlog.GET("/export", middlewares.HasPermission("monitor:operlog:list"), sc.OperLogExport)
	operlog.DELETE("/:operIds", middlewares.HasPermission("monitor:operlog:remove"), sc.OperLogRemove)
	operlog.DELETE("/clean", middlewares.HasPermission("monitor:operlog:remove"), sc.OperLogClean)
}
