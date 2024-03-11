package middlewares

import (
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/monitor/monitorService/monitorServiceImpl"
	"baize/app/utils/baizeContext"
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"io"
)

type BusinessType string

const (
	OTHER  BusinessType = "OTHER"
	INSERT BusinessType = "INSERT"
	UPDATE BusinessType = "UPDATE"
	DELETE BusinessType = "DELETE"
)

var businessTypeMap = map[BusinessType]int8{
	OTHER:  1,
	INSERT: 2,
	UPDATE: 3,
	DELETE: 4,
}

func (c BusinessType) Msg() int8 {
	msg, ok := businessTypeMap[c]
	if !ok {
		msg = businessTypeMap[OTHER]
	}
	return msg
}

// SetLog 记录日志
func SetLog(title string, businessTy BusinessType) func(c *gin.Context) {
	return func(c *gin.Context) {
		data, _ := c.GetRawData()
		c.Request.Body = io.NopCloser(bytes.NewBuffer(data))
		ol := new(monitorModels.SysOperLog)
		ol.Title = title
		ol.BusinessType = businessTy.Msg()
		c.Next()
		//ol.Status = business.Success.Msg()
		ol.OperIp = c.ClientIP()
		ol.OperUrl = c.Request.URL.Path
		ol.RequestMethod = c.Request.Method
		ol.OperName = baizeContext.GetUserName(c)
		ol.OperParam = string(data) + c.Request.URL.RawQuery
		monitorServiceImpl.OperLog.InsertOperLog(context.Background(), ol)
	}
}
