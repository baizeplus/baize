package monitorController

import (
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/monitor/monitorService"
	"baize/app/business/monitor/monitorService/monitorServiceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type OperLog struct {
	ls monitorService.ISysOperLogService
}

func NewOperLog(ls *monitorServiceImpl.OperLogService) *OperLog {
	return &OperLog{ls: ls}
}

func (ol *OperLog) OperLogList(c *gin.Context) {
	operLog := new(monitorModels.SysOperLogDQL)
	c.ShouldBind(operLog)
	list, count := ol.ls.SelectOperLogList(c, operLog)
	baizeContext.SuccessListData(c, list, count)

}

func (ol *OperLog) OperLogExport(c *gin.Context) {
}

func (ol *OperLog) OperLogRemove(c *gin.Context) {
	ol.ls.DeleteOperLogByIds(c, baizeContext.ParamInt64Array(c, "operIds"))
	baizeContext.Success(c)
}

func (ol *OperLog) OperLogClean(c *gin.Context) {
	ol.ls.CleanOperLog(c)
	baizeContext.Success(c)

}
