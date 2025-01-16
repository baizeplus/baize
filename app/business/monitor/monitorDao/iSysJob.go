package monitorDao

import (
	"baize/app/business/monitor/monitorModels"
	"context"
)

type IJobDao interface {
	SelectJobList(ctx context.Context, job *monitorModels.JobDQL) (list []*monitorModels.JobVo, total int64)
	SelectJobAll(ctx context.Context) (list []*monitorModels.JobVo)
	SelectJobById(ctx context.Context, id int64) (job *monitorModels.JobVo)
	SelectJobByInvokeTarget(ctx context.Context, invokeTarget string) (job *monitorModels.JobVo)
	DeleteJobById(ctx context.Context, id int64)
	UpdateJob(ctx context.Context, job *monitorModels.JobDML)
	InsertJob(ctx context.Context, job *monitorModels.JobDML)
	DeleteJobByIds(ctx context.Context, id []int64)
	InsertJobLog(ctx context.Context, job *monitorModels.JobLog)
	SelectJobLogList(ctx context.Context, job *monitorModels.JobLogDql) (list []*monitorModels.JobLog, total int64)
	SelectJobLogById(ctx context.Context, id int64) (vo *monitorModels.JobLog)
	SelectJobIdAndNameAll(ctx context.Context) (list []*monitorModels.JobIdAndName)
}
