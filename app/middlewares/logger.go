package middlewares

import (
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/monitor/monitorService/monitorServiceImpl"
	"baize/app/constant/sessionStatus"
	"baize/app/utils/baizeContext"
	"bytes"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"time"
)

type BusinessType string

const (
	Other         BusinessType = "Other"
	Insert        BusinessType = "Insert"
	Update        BusinessType = "Update"
	Delete        BusinessType = "Delete"
	ForcedRetreat BusinessType = "ForcedRetreat"
	Clear         BusinessType = "Clear"
)

var businessTypeMap = map[BusinessType]int8{
	Other:         0,
	Insert:        1,
	Update:        2,
	Delete:        3,
	ForcedRetreat: 4,
	Clear:         5,
}

func (c BusinessType) Msg() int8 {
	msg, ok := businessTypeMap[c]
	if !ok {
		msg = businessTypeMap[Other]
	}
	return msg
}

// SetLog 记录日志
func SetLog(title string, businessTy BusinessType) func(c *gin.Context) {
	return func(c *gin.Context) {
		start := time.Now()
		data, _ := c.GetRawData()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(data))
		ol := new(monitorModels.SysOperLog)
		ol.Title = title
		ol.BusinessType = businessTy.Msg()
		c.Next()
		value, ok := c.Get(sessionStatus.MsgKey)
		if ok {
			ol.Status = "0"
			marshal, err := json.Marshal(value)
			if err == nil {
				ol.JsonResult = string(marshal)
			}
		} else {
			ol.Status = "1"
		}
		ol.OperIp = c.ClientIP()
		ol.OperUrl = c.Request.URL.Path
		ol.RequestMethod = c.Request.Method
		ol.OperName = baizeContext.GetUserName(c)
		ol.OperParam = string(data) + c.Request.URL.RawQuery
		ol.CostTime = int64(time.Since(start))
		go func() {
			defer func() {
				if err := recover(); err != nil {
					zap.L().Error("操作日志错误", zap.Any("error", err))
				}
			}()
			monitorServiceImpl.OperLog.InsertOperLog(context.Background(), ol)
		}()
	}
}
