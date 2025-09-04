package monitorModels

import (
	"baize/app/baize"
	"time"
)

type JobIdAndName struct {
	JobId   string `json:"jobId" db:"job_id"`
	JobName string `json:"jobName" db:"job_name"`
}

type JobVo struct {
	JobId          string      `json:"jobId" db:"job_id"`
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
	JobId          string      `json:"jobId" db:"job_id"`
	JobName        string      `json:"jobName" db:"job_name"`
	JobParams      *baize.List `json:"jobParams" db:"job_params"`
	InvokeTarget   string      `json:"invokeTarget" db:"invoke_target"`
	CronExpression string      `json:"cronExpression" db:"cron_expression"`
}

type JobDML struct {
	JobId          string      `json:"jobId" db:"job_id"`
	JobName        *string     `json:"jobName" db:"job_name"`
	JobParams      *baize.List `json:"jobParams" db:"job_params"`
	InvokeTarget   string      `json:"invokeTarget" db:"invoke_target"`
	CronExpression string      `json:"cronExpression" db:"cron_expression"`
	Status         string      `json:"status" db:"status"`
	baize.BaseEntity
}

type JobLog struct {
	JobLogId      string      `json:"jobLogId" db:"job_log_id"`
	JobId         string      `json:"jobId" db:"job_id"`
	JobName       string      `json:"jobName" db:"job_name"`
	JobParams     *baize.List `json:"jobParams" db:"job_params"`
	InvokeTarget  string      `json:"invokeTarget" db:"invoke_target"`
	Status        string      `json:"status" db:"status"`
	ExceptionInfo string      `json:"exceptionInfo" db:"exception_info"`
	CreateTime    time.Time   `json:"createTime" db:"create_time"`
	CostTime      int64       `json:"costTime" db:"cost_time"`
}
type JobLogDql struct {
	JobId  string `form:"jobId" db:"job_id"`
	Status string `form:"status" db:"status"`
	baize.BaseEntityDQL
}

type JobStatus struct {
	JobId  string `json:"jobId" db:"job_id" binding:"required"`
	Status string `json:"status" db:"status"`
}
type JobRedis struct {
	Id           string
	RedisPublish bool
	Type         int8
}
