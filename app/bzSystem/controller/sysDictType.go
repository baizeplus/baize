package controller

import (
	"baize/app/bzSystem/models"
	"baize/app/bzSystem/service"
	"baize/app/bzSystem/service/serviceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DictTypeController struct {
	dts service.IDictTypeService
	dds service.IDictDataService
}

func NewDictTypeController(dts *serviceImpl.DictTypeService, dds *serviceImpl.DictDataService) *DictTypeController {
	return &DictTypeController{
		dts: dts,
		dds: dds,
	}
}

// DictTypeList 查询字典类型列表
// @Summary 查询字典类型列表
// @Description 查询字典类型列表
// @Tags 字典相关
// @Param  object query models.SysDictDataDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ResponseData{Rows=[]models.SysDictDataVo}}  "成功"
// @Router /system/dict/type/list  [get]
func (dtc *DictTypeController) DictTypeList(c *gin.Context) {
	dictType := new(models.SysDictTypeDQL)
	_ = c.ShouldBind(dictType)
	list, count := dtc.dts.SelectDictTypeList(c, dictType)
	baizeContext.SuccessListData(c, list, count)

}

func (dtc *DictTypeController) DictTypeExport(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//dictType := new(models.SysDictTypeDQL)
	//c.ShouldBind(dictType)
	//bzc.DataPackageExcel(dtc.dts.ExportDictType(dictType))
}

// DictTypeGetInfo 根据dictCode获取字典类型类型
// @Summary 根据dictCode获取字典类型信息
// @Description 根据dictCode获取字典类型信息
// @Tags 字典相关
// @Param id path string true "dictCode"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=models.SysDictDataVo}  "成功"
// @Router /system/dict/type/{dictCode}  [get]
func (dtc *DictTypeController) DictTypeGetInfo(c *gin.Context) {
	dictId := baizeContext.ParamInt64(c, "dictId")
	if dictId == 0 {
		zap.L().Error("参数错误")
		baizeContext.ParameterError(c)
		return
	}
	dictData := dtc.dts.SelectDictTypeById(c, dictId)
	baizeContext.SuccessData(c, dictData)
}

// DictTypeAdd 添加字典类型数据
// @Summary 添加字典类型数据
// @Description 添加字典类型数据
// @Tags 字典相关
// @Param  object body models.SysDictDataVo true "字典"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/dict/type  [post]
func (dtc *DictTypeController) DictTypeAdd(c *gin.Context) {
	dictType := new(models.SysDictTypeVo)
	_ = c.ShouldBindJSON(dictType)
	if dtc.dts.CheckDictTypeUnique(c, dictType.DictId, dictType.DictType) {
		baizeContext.Waring(c, "新增字典'"+dictType.DictName+"'失败，字典类型已存在")
		return
	}
	dictType.SetCreateBy(baizeContext.GetUserId(c))
	dtc.dts.InsertDictType(c, dictType)
	baizeContext.Success(c)
}

// DictTypeEdit 修改字典类型数据
// @Summary 修改字典类型数据
// @Description 修改字典类型数据
// @Tags 字典相关
// @Param  object body models.SysDictDataVo true "字典"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/dict/type  [put]
func (dtc *DictTypeController) DictTypeEdit(c *gin.Context) {
	dictType := new(models.SysDictTypeVo)
	_ = c.ShouldBindJSON(dictType)
	if dtc.dts.CheckDictTypeUnique(c, dictType.DictId, dictType.DictType) {
		baizeContext.Waring(c, "修改字典'"+dictType.DictName+"'失败，字典类型已存在")
		return
	}

	dictType.SetUpdateBy(baizeContext.GetUserId(c))
	dtc.dts.UpdateDictType(c, dictType)
	baizeContext.Success(c)
}

// DictTypeRemove 删除字典类型数据
// @Summary 根据dictCode获取字典类型信息
// @Description 根据dictCode获取字典类型信息
// @Tags 字典相关
// @Param  object body []string true "字典"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{}  "成功"
// @Router /system/dict/type  [delete]
func (dtc *DictTypeController) DictTypeRemove(c *gin.Context) {
	dictIds := baizeContext.ParamInt64Array(c, "dictIds")
	dictTypes := dtc.dts.SelectDictTypeByIds(c, dictIds)
	if dtc.dds.CheckDictDataByTypes(c, dictTypes) {
		baizeContext.Waring(c, "有已分配的字典,不能删除")
		return
	}
	dtc.dts.DeleteDictTypeByIds(c, dictIds)
	baizeContext.Success(c)
}

// DictTypeClearCache 更新字典缓存
// @Summary 更新字典缓存
// @Description 更新字典缓存
// @Tags 字典相关
// @Param  object body []string true "字典"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=models.SysDictDataVo}  "成功"
// @Router /system/dict/clearCache  [put]
func (dtc *DictTypeController) DictTypeClearCache(c *gin.Context) {
	dtc.dts.DictTypeClearCache(c)
	baizeContext.Success(c)
}

func (dtc *DictTypeController) DictTypeOptionSelect(c *gin.Context) {
	baizeContext.SuccessData(c, dtc.dts.SelectDictTypeAll(c))
}
