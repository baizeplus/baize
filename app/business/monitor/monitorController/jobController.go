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

// JobList 查询定时任务列表
// @Summary 查询定时任务列表
// @Description 查询定时任务列表
// @Tags 定时任务
// @Param  object query monitorModels.JobDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ListData{Rows=[]monitorModels.Job}}  "成功"
// @Router /monitor/job/list  [get]
func (j *Job) JobList(c *gin.Context) {
	job := new(monitorModels.JobDQL)
	c.ShouldBind(job)
	list, total := j.ls.SelectJobList(c, job)
	baizeContext.SuccessListData(c, list, total)
}

// JobGetInfo 查询定时任务信息
// @Summary 查询定时任务信息
// @Description 查询定时任务信息
// @Tags 定时任务
// @Param  jobId path int true "jobId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=monitorModels.Job}  "成功"
// @Router /monitor/job/{jobId}  [get]
func (j *Job) JobGetInfo(c *gin.Context) {
	jobId := baizeContext.ParamInt64(c, "jobId")
	menu := j.ls.SelectJobById(c, jobId)
	baizeContext.SuccessData(c, menu)
}

// JobAdd 新增定时任务
// @Summary 新增定时任务
// @Description 新增定时任务
// @Tags 定时任务
// @Param  object body monitorModels.JobDML true "新增信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /monitor/job  [post]
func (j *Job) JobAdd(c *gin.Context) {
	job := new(monitorModels.JobDML)
	c.ShouldBind(job)
	job.SetCreateBy(baizeContext.GetUserId(c))
	j.ls.InsertJob(c, job)
	baizeContext.Success(c)
}

// JobEdit 编辑定时任务
// @Summary 编辑定时任务
// @Description 编辑定时任务
// @Tags 定时任务
// @Param  object body monitorModels.JobDML true "编辑信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /monitor/job  [put]
func (j *Job) JobEdit(c *gin.Context) {
	job := new(monitorModels.JobDML)
	c.ShouldBind(job)
	job.SetUpdateBy(baizeContext.GetUserId(c))
	j.ls.UpdateJob(c, job)
	baizeContext.Success(c)
}

// JobChangeStatus 修改定时任务状态
// @Summary 修改定时任务状态
// @Description 修改定时任务状态
// @Tags 定时任务
// @Param  object body monitorModels.JobDML true "修改信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /monitor/job/changeStatus  [put]
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

// JobRun 执行定时任务
// @Summary 执行定时任务
// @Description 执行定时任务
// @Tags 定时任务
// @Param  object body monitorModels.JobVo true "执行信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /monitor/job/run  [put]
func (j *Job) JobRun(c *gin.Context) {
	vo := new(monitorModels.JobVo)
	c.ShouldBindJSON(vo)
	j.ls.Run(c, vo)
}

// JobRemove 删除定时任务
// @Summary 删除定时任务
// @Description 删除定时任务
// @Tags 定时任务
// @Param  jobIds path string true "jobIds"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /monitor/job/{jobIds}  [delete]
func (j *Job) JobRemove(c *gin.Context) {
	j.ls.DeleteJobByIds(c, baizeContext.ParamInt64Array(c, "jobIds"))
	baizeContext.Success(c)

}
