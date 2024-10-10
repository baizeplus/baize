package monitorService

import (
	"baize/app/business/monitor/monitorModels"
	"github.com/gin-gonic/gin"
)

type IJobService interface {
	SelectJobList(c *gin.Context, job *monitorModels.JobDQL) (list []*monitorModels.JobVo, total int64)
	SelectJobById(c *gin.Context, id int64) (job *monitorModels.JobVo)
	DeleteJobByIds(c *gin.Context, jobIds []int64)
	ChangeStatus(c *gin.Context, job *monitorModels.JobStatus)
	Run(c *gin.Context, job *monitorModels.JobStatus)
	InsertJob(c *gin.Context, job *monitorModels.JobDML)
	UpdateJob(c *gin.Context, job *monitorModels.JobDML)
	FunIsExist(invokeTarget string) bool
	GetFunList() []string
}
