package monitorDao

import (
	"baize/app/business/monitor/monitorModels"
	"context"
	"github.com/baizeplus/sqly"
)

type ILogininforDao interface {
	InserLogininfor(ctx context.Context, db sqly.SqlyContext, logininfor *monitorModels.Logininfor)
	SelectLogininforList(ctx context.Context, db sqly.SqlyContext, logininfor *monitorModels.LogininforDQL) (list []*monitorModels.Logininfor, total int64)
	SelectLogininforListAll(ctx context.Context, db sqly.SqlyContext, logininfor *monitorModels.LogininforDQL) (list []*monitorModels.Logininfor)
	DeleteLogininforByIds(ctx context.Context, db sqly.SqlyContext, infoIds []int64)
	CleanLogininfor(ctx context.Context, db sqly.SqlyContext)
}
