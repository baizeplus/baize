package monitorModels

import (
	"baize/app/baize"
)

type SysOperLog struct {
	OperId        int64       `json:"operId,string" db:"oper_id"`
	Title         string      `json:"title" db:"title"`
	BusinessType  int8        `json:"businessType" db:"business_type"`
	Method        string      `json:"method" db:"method"`
	RequestMethod string      `json:"requestMethod" db:"request_method"`
	OperatorType  int8        `json:"operatorType" db:"operator_type"`
	OperName      string      `json:"operName" db:"oper_name"`
	DeptName      string      `json:"deptName" db:"dept_name"`
	OperUrl       string      `json:"operUrl" db:"oper_url"`
	OperIp        string      `json:"operIp" db:"oper_ip"`
	OperLocation  string      `json:"operLocation" db:"oper_location"`
	OperParam     string      `json:"operParam" db:"oper_param"`
	JsonResult    string      `json:"jsonResult" db:"json_result"`
	Status        int8        `json:"status" db:"status"`
	ErrorMsg      string      `json:"errorMsg" db:"error_msg"`
	OperTime      *baize.Time `json:"operTime" db:"oper_time"`
}

type SysOperLogDQL struct {
	Title        string `form:"title" db:"title"`
	BusinessType *int8  `form:"businessType" db:"business_type"`
	Status       *int8  `form:"status" db:"status"`
	OperName     string `form:"operName" db:"oper_name"`
	BeginTime    string `form:"beginTime" db:"end_time"`
	EndTime      string `form:"endTime" db:"end_time"`
	baize.BaseEntityDQL
}
