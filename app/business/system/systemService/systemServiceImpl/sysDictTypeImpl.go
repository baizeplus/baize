package systemServiceImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemDao/systemDaoImpl"
	"baize/app/business/system/systemModels"
	"baize/app/utils/cache"
	"baize/app/utils/excel"
	"baize/app/utils/snowflake"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
)

type DictTypeService struct {
	data        *sqly.DB
	dictTypeDao systemDao.IDictTypeDao
	dictKey     string
}

func NewDictTypeService(data *sqly.DB, dtd *systemDaoImpl.SysDictTypeDao) *DictTypeService {
	return &DictTypeService{
		data:        data,
		dictTypeDao: dtd,
		dictKey:     "sys_dict:",
	}
}

func (dictTypeService *DictTypeService) SelectDictTypeList(c *gin.Context, dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo, total int64) {
	return dictTypeService.dictTypeDao.SelectDictTypeList(c, dictTypeService.data, dictType)

}
func (dictTypeService *DictTypeService) ExportDictType(c *gin.Context, dictType *systemModels.SysDictTypeDQL) (data []byte) {
	list := dictTypeService.dictTypeDao.SelectDictTypeAll(c, dictTypeService.data, dictType)
	toExcel, err := excel.SliceToExcel(list)
	if err != nil {
		panic(err)
	}
	buffer, err := toExcel.WriteToBuffer()
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
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
	cache.GetCache().Del(c, dictTypeService.dictKey+"*")
}
func (dictTypeService *DictTypeService) SelectDictTypeAll(c *gin.Context) (list []*systemModels.SysDictTypeVo) {
	return dictTypeService.dictTypeDao.SelectDictTypeAll(c, dictTypeService.data, new(systemModels.SysDictTypeDQL))
}
