package monitorModels

import (
	"baize/app/baize"
	"time"
)

type SysOperLog struct {
	OperId        int64      `json:"operId,string" db:"oper_id"`
	Title         string     `json:"title" db:"title"`
	BusinessType  int8       `json:"businessType" db:"business_type"`
	Method        string     `json:"method" db:"method"`
	RequestMethod string     `json:"requestMethod" db:"request_method"`
	UserId        int64      `json:"userId" db:"user_id"`
	OperName      string     `json:"operName" db:"oper_name"`
	OperUrl       string     `json:"operUrl" db:"oper_url"`
	OperIp        string     `json:"operIp" db:"oper_ip"`
	OperParam     string     `json:"operParam" db:"oper_param"`
	JsonResult    string     `json:"jsonResult" db:"json_result"`
	Status        string     `json:"status" db:"status"`
	OperTime      *time.Time `json:"operTime" db:"oper_time"`
	CostTime      int64      `json:"costTime" db:"cost_time"`
}

type SysOperLogDQL struct {
	Title        string `form:"title" db:"title"`
	BusinessType *int8  `form:"businessType" db:"business_type"`
	Status       string `form:"status" db:"status"`
	OperName     string `form:"operName" db:"oper_name"`
	BeginTime    string `form:"beginTime" db:"end_time"`
	EndTime      string `form:"endTime" db:"end_time"`
	baize.BaseEntityDQL
}
