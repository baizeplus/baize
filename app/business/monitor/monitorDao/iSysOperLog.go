package monitorDao

import (
	"baize/app/business/monitor/monitorModels"
	"context"
)

type IOperLog interface {
	InsertOperLog(ctx context.Context, operLog *monitorModels.SysOperLog)
	SelectOperLogList(ctx context.Context, openLog *monitorModels.SysOperLogDQL) (list []*monitorModels.SysOperLog, total int64)
	SelectOperLogListAll(ctx context.Context, openLog *monitorModels.SysOperLogDQL) (list []*monitorModels.SysOperLog)
	DeleteOperLogByIds(ctx context.Context, operIds []string)
	SelectOperLogById(ctx context.Context, operId string) (operLog *monitorModels.SysOperLog)
	CleanOperLog(ctx context.Context)
}
