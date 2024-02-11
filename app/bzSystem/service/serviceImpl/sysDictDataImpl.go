package serviceImpl

import (
	"baize/app/bzSystem/dao"
	"baize/app/bzSystem/dao/daoImpl"
	"baize/app/bzSystem/models"
	"baize/app/datasource"
	"baize/app/utils/snowflake"
	"context"
	"encoding/json"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
)

type DictDataService struct {
	data        *sqly.DB
	dictDataDao dao.IDictDataDao
	dictKey     string
}

func NewDictDataService(data *sqly.DB, ddd *daoImpl.SysDictDataDao) *DictDataService {
	return &DictDataService{
		data:        data,
		dictDataDao: ddd,
		dictKey:     "sys_dict:",
	}
}

func (dictDataService *DictDataService) SelectDictDataByType(c *gin.Context, dictType string) (sysDictDataList []*models.SysDictDataVo) {

	sysDictDataList = dictDataService.getDictCache(c, dictType)
	if sysDictDataList != nil {
		return
	}
	sysDictDataList = dictDataService.dictDataDao.SelectDictDataByType(c, dictDataService.data, dictType)
	if len(sysDictDataList) != 0 {
		go dictDataService.setDictCache(context.Background(), dictType, sysDictDataList)
	}
	return
}
func (dictDataService *DictDataService) SelectDictDataList(c *gin.Context, dictData *models.SysDictDataDQL) (list []*models.SysDictDataVo, count *int64) {
	return dictDataService.dictDataDao.SelectDictDataList(c, dictDataService.data, dictData)

}
func (dictDataService *DictDataService) ExportDictData(c *gin.Context, dictData *models.SysDictDataDQL) (data []byte) {
	//list, _ := dictDataService.dictDataDao.SelectDictDataList(dictDataService.data.GetSlaveDb(), dictData)
	//rows := models.SysDictDataListToRows(list)
	//return exceLize.SetRows(rows)
	return nil
}
func (dictDataService *DictDataService) SelectDictDataById(c *gin.Context, dictCode int64) (dictData *models.SysDictDataVo) {
	return dictDataService.dictDataDao.SelectDictDataById(c, dictDataService.data, dictCode)

}

func (dictDataService *DictDataService) InsertDictData(c *gin.Context, dictData *models.SysDictDataVo) {
	dictData.DictCode = snowflake.GenID()
	dictDataService.dictDataDao.InsertDictData(c, dictDataService.data, dictData)
	datasource.RedisDb.Del(c, dictDataService.dictKey+"*")
}

func (dictDataService *DictDataService) UpdateDictData(c *gin.Context, dictData *models.SysDictDataVo) {
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
func (dictDataService *DictDataService) getDictCache(c context.Context, dictType string) (dictDataList []*models.SysDictDataVo) {
	getString := datasource.RedisDb.Get(c, dictDataService.dictKey+dictType).Val()
	if getString != "" {
		dictDataList = make([]*models.SysDictDataVo, 0)
		_ = json.Unmarshal([]byte(getString), &dictDataList)
	}
	return
}

func (dictDataService *DictDataService) setDictCache(c context.Context, dictType string, dictDataList []*models.SysDictDataVo) {
	datasource.RedisDb.Set(c, dictDataService.dictKey+dictType, dictDataList, 0)
}
