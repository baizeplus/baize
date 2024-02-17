package dao

import (
	monitorModels "baize/app/bzMonitor/models"
	"context"
	"github.com/baizeplus/sqly"
)

type ILogininforDao interface {
	InserLogininfor(ctx context.Context, db sqly.SqlyContext, logininfor *monitorModels.Logininfor)
	SelectLogininforList(ctx context.Context, db sqly.SqlyContext, logininfor *monitorModels.LogininforDQL) (list []*monitorModels.Logininfor, total *int64)
	DeleteLogininforByIds(ctx context.Context, db sqly.SqlyContext, infoIds []int64)
	CleanLogininfor(ctx context.Context, db sqly.SqlyContext)
}
