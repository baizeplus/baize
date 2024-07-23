package monitorDao

import (
	"baize/app/business/monitor/monitorModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IOperLog interface {
	InsertOperLog(ctx context.Context, db sqly.SqlyContext, operLog *monitorModels.SysOperLog)
	SelectOperLogList(ctx context.Context, db sqly.SqlyContext, openLog *monitorModels.SysOperLogDQL) (list []*monitorModels.SysOperLog, total int64)
	SelectOperLogListAll(ctx context.Context, db sqly.SqlyContext, openLog *monitorModels.SysOperLogDQL) (list []*monitorModels.SysOperLog)
	DeleteOperLogByIds(ctx context.Context, db sqly.SqlyContext, operIds []int64)
	SelectOperLogById(ctx context.Context, db sqly.SqlyContext, operId int64) (operLog *monitorModels.SysOperLog)
	CleanOperLog(ctx context.Context, db sqly.SqlyContext)
}
