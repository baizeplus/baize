package monitorService

import (
	"baize/app/business/monitor/monitorModels"
	"github.com/gin-gonic/gin"
)

type ILogininforService interface {
	SelectLogininforList(c *gin.Context, logininfor *monitorModels.LogininforDQL) (list []*monitorModels.Logininfor, total int64)
	ExportLogininfor(c *gin.Context, logininfor *monitorModels.LogininforDQL) (data []byte)
	InserLogininfor(c *gin.Context, loginUser *monitorModels.Logininfor)
	DeleteLogininforByIds(c *gin.Context, infoIds []int64)
	CleanLogininfor(c *gin.Context)
}
