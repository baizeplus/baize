package systemService

import (
	"baize/app/business/system/systemModels"
	"github.com/gin-gonic/gin"
)

type IDictTypeService interface {
	SelectDictTypeList(c *gin.Context, dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo, total int64)
	ExportDictType(c *gin.Context, dictType *systemModels.SysDictTypeDQL) (data []byte)
	SelectDictTypeById(c *gin.Context, dictId int64) (dictType *systemModels.SysDictTypeVo)
	SelectDictTypeByIds(c *gin.Context, dictId []int64) (dictTypes []string)
	InsertDictType(c *gin.Context, dictType *systemModels.SysDictTypeVo)
	UpdateDictType(c *gin.Context, dictType *systemModels.SysDictTypeVo)
	DeleteDictTypeByIds(c *gin.Context, dictIds []int64)
	CheckDictTypeUnique(c *gin.Context, id int64, dictType string) bool
	DictTypeClearCache(c *gin.Context)
	SelectDictTypeAll(c *gin.Context) (list []*systemModels.SysDictTypeVo)
}
