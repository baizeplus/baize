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
	NextValidTime  string      `json:"nextValidTime"`
	Status         string      `json:"status" db:"status"`
	Remark         string      `json:"remark" db:"remark"`
	baize.BaseEntity
}
type JobDQL struct {
	JobName      string `form:"jobName" db:"job_name"`
	InvokeTarget string `form:"invokeTarget" db:"invoke_target"`
	Status       string `form:"status" db:"Status"`
	baize.BaseEntityDQL
}

type JobDML struct {
	JobId          int64       `json:"jobId,string" db:"job_id"`
	JobName        *string     `json:"jobName" db:"job_name"`
	JobParams      *baize.List `json:"jobParams" db:"job_params"`
	InvokeTarget   string      `json:"invokeTarget" db:"invoke_target"`
	CronExpression string      `json:"cronExpression" db:"cron_expression"`
	Status         string      `json:"status" db:"status"`
	Remark         *string     `json:"remark" db:"remark"`
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
