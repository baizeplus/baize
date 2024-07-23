package monitorServiceImpl

import (
	"baize/app/business/monitor/monitorDao"
	"baize/app/business/monitor/monitorDao/monitorDaoImpl"
	"baize/app/business/monitor/monitorModels"
	"baize/app/utils/excel"
	"baize/app/utils/snowflake"
	"context"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
)

var OperLog *OperLogService

type OperLogService struct {
	data *sqly.DB
	old  monitorDao.IOperLog
}

func NewOperLog(data *sqly.DB, ld *monitorDaoImpl.OperLogDao) *OperLogService {
	OperLog = &OperLogService{data: data, old: ld}
	return OperLog
}

func (ols *OperLogService) InsertOperLog(c context.Context, operLog *monitorModels.SysOperLog) {
	operLog.OperId = snowflake.GenID()
	ols.old.InsertOperLog(c, ols.data, operLog)
}
func (ols *OperLogService) SelectOperLogList(c *gin.Context, openLog *monitorModels.SysOperLogDQL) (list []*monitorModels.SysOperLog, total int64) {
	list, total = ols.old.SelectOperLogList(c, ols.data, openLog)
	return

}
func (ols *OperLogService) ExportOperLog(c *gin.Context, openLog *monitorModels.SysOperLogDQL) (data []byte) {
	list := ols.old.SelectOperLogListAll(c, ols.data, openLog)
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

func (ols *OperLogService) DeleteOperLogByIds(c *gin.Context, operIds []int64) {
	ols.old.DeleteOperLogByIds(c, ols.data, operIds)
}
func (ols *OperLogService) SelectOperLogById(c *gin.Context, operId int64) (operLog *monitorModels.SysOperLog) {
	return ols.old.SelectOperLogById(c, ols.data, operId)
}
func (ols *OperLogService) CleanOperLog(c *gin.Context) {
	ols.old.CleanOperLog(c, ols.data)
}
