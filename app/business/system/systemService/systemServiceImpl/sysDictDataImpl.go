package systemServiceImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemDao/systemDaoImpl"
	"baize/app/business/system/systemModels"
	"baize/app/utils/cache"
	"baize/app/utils/excel"
	"baize/app/utils/response"
	"baize/app/utils/snowflake"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
	"time"
)

type DictDataService struct {
	data        *sqly.DB
	dictDataDao systemDao.IDictDataDao
	dictKey     string
	gzipNil     []byte
}

func NewDictDataService(data *sqly.DB, ddd *systemDaoImpl.SysDictDataDao) *DictDataService {
	return &DictDataService{
		data:        data,
		dictDataDao: ddd,
		dictKey:     "sys_dict:",
		gzipNil:     []byte{31, 139, 8, 0, 0, 0, 0, 0, 2, 255, 170, 86, 74, 206, 79, 73, 85, 178, 50, 50, 48, 208, 81, 202, 45, 78, 87, 178, 82, 42, 46, 77, 78, 78, 45, 46, 86, 170, 5, 4, 0, 0, 255, 255, 166, 20, 213, 245, 28, 0, 0, 0},
	}
}

func (dictDataService *DictDataService) SelectDictDataByType(c *gin.Context, dictType string) (data []byte) {

	data = dictDataService.getDictCache(c, dictType)
	if data != nil && len(data) != 0 {
		return
	}
	sysDictDataList := dictDataService.dictDataDao.SelectDictDataByType(c, dictDataService.data, dictType)
	if len(sysDictDataList) != 0 {
		responseData := response.ResponseData{Code: response.Success, Msg: response.Success.Msg(), Data: sysDictDataList}
		marshal, err := json.Marshal(responseData)
		if err != nil {
			panic(err)
		}
		var buf bytes.Buffer
		gz, err := gzip.NewWriterLevel(&buf, gzip.BestCompression)
		if err != nil {
			panic(err)
		}
		if _, err = gz.Write(marshal); err != nil {
			panic(err)
		}
		if err = gz.Close(); err != nil {
			panic(err)
		}
		compressedData := buf.Bytes()
		go cache.GetCache().Set(context.Background(), dictDataService.dictKey+dictType, string(compressedData), 0)
		return compressedData
	}
	return dictDataService.gzipNil
}
func (dictDataService *DictDataService) SelectDictDataList(c *gin.Context, dictData *systemModels.SysDictDataDQL) (list []*systemModels.SysDictDataVo, total int64) {
	return dictDataService.dictDataDao.SelectDictDataList(c, dictDataService.data, dictData)

}
func (dictDataService *DictDataService) ExportDictData(c *gin.Context, dictData *systemModels.SysDictDataDQL) (data []byte) {
	list, _ := dictDataService.dictDataDao.SelectDictDataList(c, dictDataService.data, dictData)
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
func (dictDataService *DictDataService) SelectDictDataById(c *gin.Context, dictCode int64) (dictData *systemModels.SysDictDataVo) {
	return dictDataService.dictDataDao.SelectDictDataById(c, dictDataService.data, dictCode)

}

func (dictDataService *DictDataService) InsertDictData(c *gin.Context, dictData *systemModels.SysDictDataVo) {
	dictData.DictCode = snowflake.GenID()
	dictDataService.dictDataDao.InsertDictData(c, dictDataService.data, dictData)
	dictDataService.deleteDictCache(dictData.DictType)

}

func (dictDataService *DictDataService) UpdateDictData(c *gin.Context, dictData *systemModels.SysDictDataVo) {
	dictDataService.dictDataDao.UpdateDictData(c, dictDataService.data, dictData)
	dictDataService.deleteDictCache(dictData.DictType)
}
func (dictDataService *DictDataService) DeleteDictDataByIds(c *gin.Context, dictCodes []int64) {

	codes := dictDataService.dictDataDao.SelectDictTypesByDictCodes(c, dictDataService.data, dictCodes)
	dictDataService.dictDataDao.DeleteDictDataByIds(c, dictDataService.data, dictCodes)
	for _, code := range codes {
		dictDataService.deleteDictCache(code)
	}

}
func (dictDataService *DictDataService) CheckDictDataByTypes(c *gin.Context, dictType []string) bool {
	return dictDataService.dictDataDao.CountDictDataByTypes(c, dictDataService.data, dictType) > 0

}
func (dictDataService *DictDataService) getDictCache(c context.Context, dictType string) (dictDataList []byte) {
	getString, err := cache.GetCache().Get(c, dictDataService.dictKey+dictType)
	if err != nil {
		return nil
	}
	return []byte(getString)
}

func (dictDataService *DictDataService) deleteDictCache(dictType string) {
	go func() {
		time.Sleep(time.Second * 3)
		cache.GetCache().Del(context.Background(), dictDataService.dictKey+dictType)
	}()
}
