package monitorDao

import (
	"baize/app/business/monitor/monitorModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IJobDao interface {
	SelectJobList(ctx context.Context, db sqly.SqlyContext, job *monitorModels.JobDQL) (list []*monitorModels.JobVo, total int64)
	SelectJobAll(ctx context.Context, db sqly.SqlyContext) (list []*monitorModels.JobVo)
	SelectJobById(ctx context.Context, db sqly.SqlyContext, id int64) (job *monitorModels.JobVo)
	SelectJobByInvokeTarget(ctx context.Context, db sqly.SqlyContext, invokeTarget string) (job *monitorModels.JobVo)
	DeleteJobById(ctx context.Context, db sqly.SqlyContext, id int64)
	UpdateJob(ctx context.Context, db sqly.SqlyContext, job *monitorModels.JobDML)
	InsertJob(ctx context.Context, db sqly.SqlyContext, job *monitorModels.JobDML)
	DeleteJobByIds(ctx context.Context, db sqly.SqlyContext, id []int64)
	InsertJobLog(ctx context.Context, db sqly.SqlyContext, job *monitorModels.JobLog)
	SelectJobLogList(ctx context.Context, db sqly.SqlyContext, job *monitorModels.JobDQL) (list []*monitorModels.JobLog, total int64)
	SelectJobLogById(ctx context.Context, db sqly.SqlyContext, id int64) (vo *monitorModels.JobLog)
}
