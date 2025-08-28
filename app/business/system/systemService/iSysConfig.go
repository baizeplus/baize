package systemService

import (
	"baize/app/business/system/systemModels"

	"github.com/gin-gonic/gin"
)

type IConfigService interface {
	SelectConfigList(c *gin.Context, config *systemModels.SysConfigDQL) (sysConfigList []*systemModels.SysConfigVo, total int64)
	ConfigExport(c *gin.Context, config *systemModels.SysConfigDQL) (data []byte)
	SelectConfigById(c *gin.Context, configId string) (Config *systemModels.SysConfigVo)
	InsertConfig(c *gin.Context, config *systemModels.SysConfigVo)
	UpdateConfig(c *gin.Context, config *systemModels.SysConfigVo)
	DeleteConfigById(c *gin.Context, configId string)
	CheckConfigKeyUnique(c *gin.Context, configId string, configKey string) bool
	SelectConfigValueByKey(c *gin.Context, configKey string) string
}
