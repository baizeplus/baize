package systemServiceImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemDao/systemDaoImpl"
	"baize/app/business/system/systemModels"
	"baize/app/datasource"
	"baize/app/utils/snowflake"
	"context"
	"encoding/json"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
)

type DictDataService struct {
	data        *sqly.DB
	dictDataDao systemDao.IDictDataDao
	dictKey     string
}

func NewDictDataService(data *sqly.DB, ddd *systemDaoImpl.SysDictDataDao) *DictDataService {
	return &DictDataService{
		data:        data,
		dictDataDao: ddd,
		dictKey:     "sys_dict:",
	}
}

func (dictDataService *DictDataService) SelectDictDataByType(c *gin.Context, dictType string) (sysDictDataList []*systemModels.SysDictDataVo) {

	sysDictDataList = dictDataService.getDictCache(c, dictType)
	if sysDictDataList != nil && len(sysDictDataList) != 0 {
		return
	}
	sysDictDataList = dictDataService.dictDataDao.SelectDictDataByType(c, dictDataService.data, dictType)
	if len(sysDictDataList) != 0 {
		go dictDataService.setDictCache(context.Background(), dictType, sysDictDataList)
	}
	return
}
func (dictDataService *DictDataService) SelectDictDataList(c *gin.Context, dictData *systemModels.SysDictDataDQL) (list []*systemModels.SysDictDataVo, count *int64) {
	return dictDataService.dictDataDao.SelectDictDataList(c, dictDataService.data, dictData)

}
func (dictDataService *DictDataService) ExportDictData(c *gin.Context, dictData *systemModels.SysDictDataDQL) (data []byte) {
	//list, _ := dictDataService.dictDataDao.SelectDictDataList(dictDataService.data.GetSlaveDb(), dictData)
	//rows := systemModels.SysDictDataListToRows(list)
	//return exceLize.SetRows(rows)
	return nil
}
func (dictDataService *DictDataService) SelectDictDataById(c *gin.Context, dictCode int64) (dictData *systemModels.SysDictDataVo) {
	return dictDataService.dictDataDao.SelectDictDataById(c, dictDataService.data, dictCode)

}

func (dictDataService *DictDataService) InsertDictData(c *gin.Context, dictData *systemModels.SysDictDataVo) {
	dictData.DictCode = snowflake.GenID()
	dictDataService.dictDataDao.InsertDictData(c, dictDataService.data, dictData)
	datasource.RedisDb.Del(c, dictDataService.dictKey+"*")
}

func (dictDataService *DictDataService) UpdateDictData(c *gin.Context, dictData *systemModels.SysDictDataVo) {
	dictDataService.dictDataDao.UpdateDictData(c, dictDataService.data, dictData)
	datasource.RedisDb.Del(c, dictDataService.dictKey+"*")
}
func (dictDataService *DictDataService) DeleteDictDataByIds(c *gin.Context, dictCodes []int64) {
	dictDataService.dictDataDao.DeleteDictDataByIds(c, dictDataService.data, dictCodes)
	datasource.RedisDb.Del(c, dictDataService.dictKey+"*")
}
func (dictDataService *DictDataService) CheckDictDataByTypes(c *gin.Context, dictType []string) bool {
	return dictDataService.dictDataDao.CountDictDataByTypes(c, dictDataService.data, dictType) > 0

}
func (dictDataService *DictDataService) getDictCache(c context.Context, dictType string) (dictDataList []*systemModels.SysDictDataVo) {
	getString := datasource.RedisDb.Get(c, dictDataService.dictKey+dictType).Val()
	if getString != "" {
		dictDataList = make([]*systemModels.SysDictDataVo, 0)
		_ = json.Unmarshal([]byte(getString), &dictDataList)
	}
	return
}

func (dictDataService *DictDataService) setDictCache(c context.Context, dictType string, dictDataList []*systemModels.SysDictDataVo) {
	datasource.RedisDb.Set(c, dictDataService.dictKey+dictType, dictDataList, 0)
}
