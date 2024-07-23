package systemService

import (
	"baize/app/business/system/systemModels"
	"github.com/gin-gonic/gin"
)

type IDictDataService interface {
	SelectDictDataByType(c *gin.Context, dictType string) (data []byte)
	SelectDictDataList(c *gin.Context, dictData *systemModels.SysDictDataDQL) (list []*systemModels.SysDictDataVo, total int64)
	ExportDictData(c *gin.Context, dictData *systemModels.SysDictDataDQL) (data []byte)
	SelectDictDataById(c *gin.Context, dictCode int64) (dictData *systemModels.SysDictDataVo)
	InsertDictData(c *gin.Context, dictData *systemModels.SysDictDataVo)
	UpdateDictData(c *gin.Context, dictData *systemModels.SysDictDataVo)
	DeleteDictDataByIds(c *gin.Context, dictCodes []int64)
	CheckDictDataByTypes(c *gin.Context, dictType []string) bool
}
