package monitorController

import (
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/monitor/monitorService"
	"baize/app/business/monitor/monitorService/monitorServiceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type Job struct {
	ls monitorService.IJobService
}

func NewJob(ls *monitorServiceImpl.JobService) *Job {
	return &Job{ls: ls}
}
func (j *Job) JobList(c *gin.Context) {
	job := new(monitorModels.JobDQL)
	c.ShouldBind(job)
	list, total := j.ls.SelectJobList(c, job)
	baizeContext.SuccessListData(c, list, total)
}

func (j *Job) JobGetInfo(c *gin.Context) {
	jobId := baizeContext.ParamInt64(c, "jobId")
	menu := j.ls.SelectJobById(c, jobId)
	baizeContext.SuccessData(c, menu)
}
func (j *Job) JobAdd(c *gin.Context) {
	job := new(monitorModels.JobDML)
	c.ShouldBind(job)
	job.SetCreateBy(baizeContext.GetUserId(c))
	j.ls.InsertJob(c, job)
	baizeContext.Success(c)
}
func (j *Job) JobEdit(c *gin.Context) {
	job := new(monitorModels.JobDML)
	c.ShouldBind(job)
	job.SetUpdateBy(baizeContext.GetUserId(c))
	j.ls.UpdateJob(c, job)
	baizeContext.Success(c)
}
func (j *Job) JobChangeStatus(c *gin.Context) {
	job := new(monitorModels.JobDML)
	c.ShouldBind(job)
	status := j.ls.ChangeStatus(c, job)
	if !status {
		baizeContext.Waring(c, "目标方法未找到")
		return
	}
	baizeContext.Success(c)
}
func (j *Job) JobRun(c *gin.Context) {
	vo := new(monitorModels.JobVo)
	c.ShouldBindJSON(vo)
	j.ls.Run(c, vo)
}
func (j *Job) JobRemove(c *gin.Context) {

}
