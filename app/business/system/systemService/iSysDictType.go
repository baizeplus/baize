package systemService

import (
	"baize/app/business/system/systemModels"

	"github.com/gin-gonic/gin"
)

type IDictTypeService interface {
	SelectDictTypeList(c *gin.Context, dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo, total int64)
	ExportDictType(c *gin.Context, dictType *systemModels.SysDictTypeDQL) (data []byte)
	SelectDictTypeById(c *gin.Context, dictId string) (dictType *systemModels.SysDictTypeVo)
	SelectDictTypeByIds(c *gin.Context, dictId []string) (dictTypes []string)
	InsertDictType(c *gin.Context, dictType *systemModels.SysDictTypeVo)
	UpdateDictType(c *gin.Context, dictType *systemModels.SysDictTypeVo)
	DeleteDictTypeByIds(c *gin.Context, dictIds []string)
	CheckDictTypeUnique(c *gin.Context, id string, dictType string) bool
	DictTypeClearCache(c *gin.Context)
	SelectDictTypeAll(c *gin.Context) (list []*systemModels.SysDictTypeVo)
}
