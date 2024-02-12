package controller

import (
	"baize/app/bzSystem/models"
	"baize/app/bzSystem/service"
	"baize/app/bzSystem/service/serviceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type DictDataController struct {
	dds service.IDictDataService
}

func NewDictDataController(dds *serviceImpl.DictDataService) *DictDataController {
	return &DictDataController{
		dds: dds,
	}
}

// DictDataList 查询字典列表
// @Summary 查询字典列表
// @Description 查询字典列表
// @Tags 字典相关
// @Param  object query models.SysDictDataDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ResponseData{Rows=[]models.SysDictDataVo}}  "成功"
// @Router /system/dict/data/list  [get]
func (ddc *DictDataController) DictDataList(c *gin.Context) {
	dictData := new(models.SysDictDataDQL)
	_ = c.ShouldBind(dictData)
	list, count := ddc.dds.SelectDictDataList(c, dictData)
	baizeContext.SuccessListData(c, list, count)
}

func (ddc *DictDataController) DictDataExport(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//dictData := new(models.SysDictDataDQL)
	//_ = c.ShouldBind(dictData)
	//bzc.DataPackageExcel(ddc.dds.ExportDictData(dictData))
}

// DictDataGetInfo 根据dictCode获取字典信息
// @Summary 根据dictCode获取字典信息
// @Description 根据dictCode获取字典信息
// @Tags 字典相关
// @Param id path string true "dictCode"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=models.SysDictDataVo}  "成功"
// @Router /system/dict/data/{dictCode}  [get]
func (ddc *DictDataController) DictDataGetInfo(c *gin.Context) {
	dictCode := baizeContext.ParamInt64(c, "dictCode")
	if dictCode == 0 {
		baizeContext.ParameterError(c)
		return
	}
	dictData := ddc.dds.SelectDictDataById(c, dictCode)
	baizeContext.SuccessData(c, dictData)
}

// DictDataType 查询字典列表
// @Summary 查询字典列表
// @Description 查询字典列表
// @Tags 字典相关
// @Param id path string true "dictType"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ResponseData{Rows=[]models.SysDictDataVo}}  "成功"
// @Router /system/dict/data/type/{dictType}  [get]
func (ddc *DictDataController) DictDataType(c *gin.Context) {
	sysDictDataList := ddc.dds.SelectDictDataByType(c, c.Param("dictType"))
	baizeContext.SuccessData(c, sysDictDataList)
}

// DictDataAdd 添加字典数据
// @Summary 添加字典数据
// @Description 添加字典数据
// @Tags 字典相关
// @Param  object body models.SysDictDataVo true "字典"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/dict/data  [post]
func (ddc *DictDataController) DictDataAdd(c *gin.Context) {
	dictData := new(models.SysDictDataVo)
	_ = c.ShouldBindJSON(dictData)
	dictData.SetCreateBy(baizeContext.GetUserId(c))
	ddc.dds.InsertDictData(c, dictData)
	baizeContext.Success(c)
}

// DictDataEdit 修改字典数据
// @Summary 修改字典数据
// @Description 修改字典数据
// @Tags 字典相关
// @Param  object body models.SysDictDataVo true "字典"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/dict/data  [put]
func (ddc *DictDataController) DictDataEdit(c *gin.Context) {
	dictData := new(models.SysDictDataVo)
	_ = c.ShouldBindJSON(dictData)
	dictData.SetUpdateBy(baizeContext.GetUserId(c))
	ddc.dds.UpdateDictData(c, dictData)
	baizeContext.Success(c)
}

// DictDataRemove 删除字典数据
// @Summary 根据dictCode获取字典信息
// @Description 根据dictCode获取字典信息
// @Tags 字典相关
// @Param  dictCodes path []int64 true "dictCodes"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=models.SysDictDataVo}  "成功"
// @Router /system/dict/data/{dictCodes}  [delete]
func (ddc *DictDataController) DictDataRemove(c *gin.Context) {
	ddc.dds.DeleteDictDataByIds(c, baizeContext.ParamInt64Array(c, "dictCodes"))
	baizeContext.Success(c)
}
