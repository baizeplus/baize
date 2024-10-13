package monitorModels

import (
	"baize/app/baize"
	"time"
)

type JobVo struct {
	JobId          int64       `json:"jobId,string" db:"job_id"`
	JobName        string      `json:"jobName" db:"job_name"`
	JobParams      *baize.List `json:"jobParams" db:"job_params"`
	InvokeTarget   string      `json:"invokeTarget" db:"invoke_target"`
	CronExpression string      `json:"cronExpression" db:"cron_expression"`
	NextValidTime  *time.Time  `json:"nextValidTime"`
	Status         string      `json:"status" db:"status"`
	baize.BaseEntity
}
type JobDQL struct {
	JobName      string `form:"jobName" db:"job_name"`
	InvokeTarget string `form:"invokeTarget" db:"invoke_target"`
	Status       string `form:"status" db:"Status"`
	baize.BaseEntityDQL
}
type JobRun struct {
	JobId          int64       `json:"jobId,string" db:"job_id"`
	JobName        string      `json:"jobName" db:"job_name"`
	JobParams      *baize.List `json:"jobParams" db:"job_params"`
	InvokeTarget   string      `json:"invokeTarget" db:"invoke_target"`
	CronExpression string      `json:"cronExpression" db:"cron_expression"`
}

type JobDML struct {
	JobId          int64       `json:"jobId,string" db:"job_id"`
	JobName        *string     `json:"jobName" db:"job_name"`
	JobParams      *baize.List `json:"jobParams" db:"job_params"`
	InvokeTarget   string      `json:"invokeTarget" db:"invoke_target"`
	CronExpression string      `json:"cronExpression" db:"cron_expression"`
	Status         string      `json:"status" db:"status"`
	baize.BaseEntity
}

type JobLog struct {
	JobLogId      int64       `json:"jobLogId,string" db:"job_log_id"`
	JobId         int64       `json:"jobId,string" db:"job_id"`
	JobName       string      `json:"jobName" db:"job_name"`
	JobParams     *baize.List `json:"jobParams" db:"job_params"`
	InvokeTarget  string      `json:"invokeTarget" db:"invoke_target"`
	Status        string      `json:"status" db:"status"`
	ExceptionInfo string      `json:"exceptionInfo" db:"exception_info"`
	CreateTime    time.Time   `json:"createTime" db:"create_time"`
	CostTime      int64       `json:"costTime" db:"cost_time"`
}
type JobStatus struct {
	JobId  int64  `json:"jobId,string" db:"job_id" binding:"required"`
	Status string `json:"status" db:"status"`
}
type JobRedis struct {
	Id           int64
	RedisPublish bool
	Type         int8
}
