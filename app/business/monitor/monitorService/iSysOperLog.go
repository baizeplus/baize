package monitorService

import (
	"baize/app/business/monitor/monitorModels"
	"github.com/gin-gonic/gin"
)

type ISysOperLogService interface {
	SelectOperLogList(c *gin.Context, openLog *monitorModels.SysOperLogDQL) (list []*monitorModels.SysOperLog, total int64)
	ExportOperLog(c *gin.Context, logininfor *monitorModels.SysOperLogDQL) (data []byte)
	DeleteOperLogByIds(c *gin.Context, operIds []int64)
	SelectOperLogById(c *gin.Context, operId int64) (operLogList *monitorModels.SysOperLog)
	CleanOperLog(c *gin.Context)
}
