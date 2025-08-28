package monitorServiceImpl

import (
	"baize/app/business/monitor/monitorDao"
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/monitor/monitorService"
	"baize/app/utils/excel"
	"github.com/gin-gonic/gin"
)

type OperLogService struct {
	old monitorDao.IOperLog
}

func NewOperLog(old monitorDao.IOperLog) monitorService.ISysOperLogService {
	return &OperLogService{old: old}
}

func (ols *OperLogService) SelectOperLogList(c *gin.Context, openLog *monitorModels.SysOperLogDQL) (list []*monitorModels.SysOperLog, total int64) {
	list, total = ols.old.SelectOperLogList(c, openLog)
	return

}
func (ols *OperLogService) ExportOperLog(c *gin.Context, openLog *monitorModels.SysOperLogDQL) (data []byte) {
	list := ols.old.SelectOperLogListAll(c, openLog)
	toExcel, err := excel.SliceToExcel(list)
	if err != nil {
		panic(err)
	}
	buffer, err := toExcel.WriteToBuffer()
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

func (ols *OperLogService) DeleteOperLogByIds(c *gin.Context, operIds []string) {
	ols.old.DeleteOperLogByIds(c, operIds)
}
func (ols *OperLogService) SelectOperLogById(c *gin.Context, operId string) (operLog *monitorModels.SysOperLog) {
	return ols.old.SelectOperLogById(c, operId)
}
func (ols *OperLogService) CleanOperLog(c *gin.Context) {
	ols.old.CleanOperLog(c)
}
