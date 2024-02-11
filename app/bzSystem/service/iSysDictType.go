package service

import (
	"baize/app/bzSystem/models"
	"github.com/gin-gonic/gin"
)

type IDictTypeService interface {
	SelectDictTypeList(c *gin.Context, dictType *models.SysDictTypeDQL) (list []*models.SysDictTypeVo, count *int64)
	ExportDictType(c *gin.Context, dictType *models.SysDictTypeDQL) (data []byte)
	SelectDictTypeById(c *gin.Context, dictId int64) (dictType *models.SysDictTypeVo)
	SelectDictTypeByIds(c *gin.Context, dictId []int64) (dictTypes []string)
	InsertDictType(c *gin.Context, dictType *models.SysDictTypeVo)
	UpdateDictType(c *gin.Context, dictType *models.SysDictTypeVo)
	DeleteDictTypeByIds(c *gin.Context, dictIds []int64)
	CheckDictTypeUnique(c *gin.Context, id int64, dictType string) bool
	DictTypeClearCache(c *gin.Context)
	SelectDictTypeAll(c *gin.Context) (list []*models.SysDictTypeVo)
}
