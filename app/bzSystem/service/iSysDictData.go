package service

import (
	"baize/app/bzSystem/models"
	"github.com/gin-gonic/gin"
)

type IDictDataService interface {
	SelectDictDataByType(c *gin.Context, dictType string) (sysDictDataList []*models.SysDictDataVo)
	SelectDictDataList(c *gin.Context, dictData *models.SysDictDataDQL) (list []*models.SysDictDataVo, count *int64)
	ExportDictData(c *gin.Context, dictData *models.SysDictDataDQL) (data []byte)
	SelectDictDataById(c *gin.Context, dictCode int64) (dictData *models.SysDictDataVo)
	InsertDictData(c *gin.Context, dictData *models.SysDictDataVo)
	UpdateDictData(c *gin.Context, dictData *models.SysDictDataVo)
	DeleteDictDataByIds(c *gin.Context, dictCodes []int64)
	CheckDictDataByTypes(c *gin.Context, dictType []string) bool
}
