package systemServiceImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemDao/systemDaoImpl"
	"baize/app/business/system/systemModels"
	"baize/app/utils/snowflake"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
)

type DictTypeService struct {
	data        *sqly.DB
	dictTypeDao systemDao.IDictTypeDao
}

func NewDictTypeService(data *sqly.DB, dtd *systemDaoImpl.SysDictTypeDao) *DictTypeService {
	return &DictTypeService{
		data:        data,
		dictTypeDao: dtd,
	}
}

func (dictTypeService *DictTypeService) SelectDictTypeList(c *gin.Context, dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo, count *int64) {
	return dictTypeService.dictTypeDao.SelectDictTypeList(c, dictTypeService.data, dictType)

}
func (dictTypeService *DictTypeService) ExportDictType(c *gin.Context, dictType *systemModels.SysDictTypeDQL) (data []byte) {
	//list, _ := dictTypeService.dictTypeDao.SelectDictTypeList(dictTypeService.data, dictType)
	//rows := systemModels.SysDictTypeListToRows(list)
	//return exceLize.SetRows(rows)
	return nil
}

func (dictTypeService *DictTypeService) SelectDictTypeById(c *gin.Context, dictId int64) (dictType *systemModels.SysDictTypeVo) {
	return dictTypeService.dictTypeDao.SelectDictTypeById(c, dictTypeService.data, dictId)

}
func (dictTypeService *DictTypeService) SelectDictTypeByIds(c *gin.Context, dictId []int64) (dictTypes []string) {
	return dictTypeService.dictTypeDao.SelectDictTypeByIds(c, dictTypeService.data, dictId)
}

func (dictTypeService *DictTypeService) InsertDictType(c *gin.Context, dictType *systemModels.SysDictTypeVo) {
	dictType.DictId = snowflake.GenID()
	dictTypeService.dictTypeDao.InsertDictType(c, dictTypeService.data, dictType)
}

func (dictTypeService *DictTypeService) UpdateDictType(c *gin.Context, dictType *systemModels.SysDictTypeVo) {
	dictTypeService.dictTypeDao.UpdateDictType(c, dictTypeService.data, dictType)
}
func (dictTypeService *DictTypeService) DeleteDictTypeByIds(c *gin.Context, dictIds []int64) {
	dictTypeService.dictTypeDao.DeleteDictTypeByIds(c, dictTypeService.data, dictIds)
}

func (dictTypeService *DictTypeService) CheckDictTypeUnique(c *gin.Context, id int64, dictType string) bool {
	dictId := dictTypeService.dictTypeDao.CheckDictTypeUnique(c, dictTypeService.data, dictType)
	if dictId == id || dictId == 0 {
		return false
	}
	return true
}
func (dictTypeService *DictTypeService) DictTypeClearCache(c *gin.Context) {
	//redis.Delete(constants.SysDictKey + "*")
}
func (dictTypeService *DictTypeService) SelectDictTypeAll(c *gin.Context) (list []*systemModels.SysDictTypeVo) {
	return dictTypeService.dictTypeDao.SelectDictTypeAll(c, dictTypeService.data)
}
