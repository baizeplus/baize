package monitorRouter

import (
	"baize/app/business/monitor/monitorController"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitJobRouter(router *gin.RouterGroup, jc *monitorController.Job) {
	job := router.Group("/monitor/job")
	job.GET("/list", middlewares.HasPermission("monitor:job:list"), jc.JobList)
	job.GET("/:jobId", middlewares.HasPermission("monitor:job:query"), jc.JobGetInfo)
	job.POST("", middlewares.HasPermission("monitor:job:add"), jc.JobAdd)
	job.PUT("", middlewares.HasPermission("monitor:job:edit"), jc.JobEdit)
	job.PUT("/changeStatus", middlewares.HasPermission("monitor:job:changeStatus"), jc.JobChangeStatus)
	job.PUT("/run", middlewares.HasPermission("monitor:job:changeStatus"), jc.JobRun)
	job.DELETE("/:jobIds", middlewares.HasPermission("monitor:job:remove"), jc.JobRemove)
	job.GET("/funList", middlewares.HasPermission("monitor:job:list"), jc.FunList)
}
