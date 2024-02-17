package service

import (
	"baize/app/bzMonitor/models"
	"github.com/gin-gonic/gin"
)

type ILogininforService interface {
	SelectLogininforList(c *gin.Context, logininfor *models.LogininforDQL) (list []*models.Logininfor, total *int64)
	ExportLogininfor(c *gin.Context, logininfor *models.LogininforDQL) (data []byte)
	InserLogininfor(c *gin.Context, loginUser *models.Logininfor)
	DeleteLogininforByIds(c *gin.Context, infoIds []int64)
	CleanLogininfor(c *gin.Context)
}
