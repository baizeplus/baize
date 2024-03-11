package monitorRouter

import (
	"baize/app/business/monitor/monitorController"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysOperLogRouter(router *gin.RouterGroup, sc *monitorController.OperLog) {
	operlog := router.Group("/monitor/operlog")
	operlog.GET("/list", middlewares.HasPermission("system:operlog:list"), sc.OperLogList)
	operlog.GET("/export", middlewares.HasPermission("system:operlog:list"), sc.OperLogExport)
	operlog.DELETE("/:operIds", middlewares.HasPermission("system:operlog:remove"), sc.OperLogRemove)
	operlog.DELETE("/clean", middlewares.HasPermission("system:operlog:remove"), sc.OperLogClean)
}
