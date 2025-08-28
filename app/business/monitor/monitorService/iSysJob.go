package monitorService

import (
	"baize/app/business/monitor/monitorModels"
	"context"

	"github.com/gin-gonic/gin"
)

type IJobService interface {
	SelectJobList(c *gin.Context, job *monitorModels.JobDQL) (list []*monitorModels.JobVo, total int64)
	SelectJobById(c *gin.Context, id string) (job *monitorModels.JobVo)
	DeleteJobByIds(c *gin.Context, jobIds []string)
	ChangeStatus(c *gin.Context, job *monitorModels.JobStatus)
	Run(c *gin.Context, job *monitorModels.JobStatus)
	InsertJob(c *gin.Context, job *monitorModels.JobDML)
	UpdateJob(c *gin.Context, job *monitorModels.JobDML)
	StartRunCron(c context.Context, jo *monitorModels.JobRedis)
	DeleteRunCron(c context.Context, jo *monitorModels.JobRedis)
	FunIsExist(invokeTarget string) bool
	GetFunList() []string
	SelectJobLogList(c *gin.Context, job *monitorModels.JobLogDql) (list []*monitorModels.JobLog, total int64)
	SelectJobLogById(c *gin.Context, id string) (vo *monitorModels.JobLog)
	SelectJobIdAndNameAll(c *gin.Context) (list []*monitorModels.JobIdAndName)
}
