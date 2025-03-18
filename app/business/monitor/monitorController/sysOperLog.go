package monitorController

import (
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/monitor/monitorService"
	"baize/app/middlewares"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type OperLog struct {
	ls monitorService.ISysOperLogService
}

func NewOperLog(ls monitorService.ISysOperLogService) *OperLog {
	return &OperLog{ls: ls}
}

func (ol *OperLog) PrivateRoutes(router *gin.RouterGroup) {
	operlog := router.Group("/monitor/operlog")
	operlog.GET("/list", middlewares.HasPermission("system:monitor:operlog"), ol.OperLogList)
	operlog.POST("/export", middlewares.HasPermission("monitor:operlog:export"), ol.OperLogExport)
	operlog.DELETE("/:operIds", middlewares.SetLog("操作日志", middlewares.Delete), middlewares.HasPermission("monitor:operlog:remove"), ol.OperLogRemove)
	operlog.DELETE("/clean", middlewares.SetLog("操作日志", middlewares.Clear), middlewares.HasPermission("monitor:operlog:remove"), ol.OperLogClean)

}

func (ol *OperLog) OperLogList(c *gin.Context) {
	operLog := new(monitorModels.SysOperLogDQL)
	_ = c.ShouldBind(operLog)
	if operLog.OrderBy == "" {
		operLog.OrderBy = "oper_id"
		operLog.IsAsc = "desc"
	}
	list, count := ol.ls.SelectOperLogList(c, operLog)
	baizeContext.SuccessListData(c, list, count)

}

func (ol *OperLog) OperLogExport(c *gin.Context) {
	operLog := new(monitorModels.SysOperLogDQL)
	_ = c.ShouldBind(operLog)
	data := ol.ls.ExportOperLog(c, operLog)
	baizeContext.DataPackageExcel(c, data)
}

func (ol *OperLog) OperLogRemove(c *gin.Context) {
	ol.ls.DeleteOperLogByIds(c, baizeContext.ParamInt64Array(c, "operIds"))
	baizeContext.Success(c)
}

func (ol *OperLog) OperLogClean(c *gin.Context) {
	ol.ls.CleanOperLog(c)
	baizeContext.Success(c)

}
