package systemServiceImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/datasource/cache"
	"baize/app/utils/excel"
	"baize/app/utils/snowflake"

	"github.com/gin-gonic/gin"
)

type DictTypeService struct {
	cache       cache.Cache
	dictTypeDao systemDao.IDictTypeDao
	dictKey     string
}

func NewDictTypeService(dtd systemDao.IDictTypeDao,
	cache cache.Cache) systemService.IDictTypeService {
	return &DictTypeService{
		cache:       cache,
		dictTypeDao: dtd,
		dictKey:     "sys_dict:",
	}
}

func (dictTypeService *DictTypeService) SelectDictTypeList(c *gin.Context, dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo, total int64) {
	return dictTypeService.dictTypeDao.SelectDictTypeList(c, dictType)

}
func (dictTypeService *DictTypeService) ExportDictType(c *gin.Context, dictType *systemModels.SysDictTypeDQL) (data []byte) {
	list := dictTypeService.dictTypeDao.SelectDictTypeAll(c, dictType)
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

func (dictTypeService *DictTypeService) SelectDictTypeById(c *gin.Context, dictId string) (dictType *systemModels.SysDictTypeVo) {
	return dictTypeService.dictTypeDao.SelectDictTypeById(c, dictId)

}
func (dictTypeService *DictTypeService) SelectDictTypeByIds(c *gin.Context, dictId []string) (dictTypes []string) {
	return dictTypeService.dictTypeDao.SelectDictTypeByIds(c, dictId)
}

func (dictTypeService *DictTypeService) InsertDictType(c *gin.Context, dictType *systemModels.SysDictTypeVo) {
	dictType.DictId = snowflake.GenID()
	dictTypeService.dictTypeDao.InsertDictType(c, dictType)
}

func (dictTypeService *DictTypeService) UpdateDictType(c *gin.Context, dictType *systemModels.SysDictTypeVo) {
	dictTypeService.dictTypeDao.UpdateDictType(c, dictType)
}
func (dictTypeService *DictTypeService) DeleteDictTypeByIds(c *gin.Context, dictIds []string) {
	dictTypeService.dictTypeDao.DeleteDictTypeByIds(c, dictIds)
}

func (dictTypeService *DictTypeService) CheckDictTypeUnique(c *gin.Context, id string, dictType string) bool {
	dictId := dictTypeService.dictTypeDao.CheckDictTypeUnique(c, dictType)
	if dictId == id || dictId == "" {
		return false
	}
	return true
}
func (dictTypeService *DictTypeService) DictTypeClearCache(c *gin.Context) {
	dictTypeService.cache.Del(c, dictTypeService.dictKey+"*")
}
func (dictTypeService *DictTypeService) SelectDictTypeAll(c *gin.Context) (list []*systemModels.SysDictTypeVo) {
	return dictTypeService.dictTypeDao.SelectDictTypeAll(c, new(systemModels.SysDictTypeDQL))
}
