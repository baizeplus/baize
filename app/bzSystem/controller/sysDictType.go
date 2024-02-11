package controller

import (
	"baize/app/bzSystem/models"
	"baize/app/bzSystem/service"
	"baize/app/bzSystem/service/serviceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type DictTypeController struct {
	dts service.IDictTypeService
}

func NewDictTypeController(dts *serviceImpl.DictTypeService) *DictTypeController {
	return &DictTypeController{
		dts: dts,
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
	//bzc := baizeContext.NewBaiZeContext(c)
	//dictId := bzc.ParamInt64("dictId")
	//if dictId == 0 {
	//	zap.L().Error("参数错误")
	//	bzc.ParameterError()
	//	return
	//}
	//dictData := dtc.dts.SelectDictTypeById(dictId)
	//
	//bzc.SuccessData(dictData)
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
	//bzc := baizeContext.NewBaiZeContext(c)
	//dictType := new(models.SysDictTypeAdd)
	//if err := c.ShouldBindJSON(dictType); err != nil {
	//	bzc.ParameterError()
	//	return
	//}
	//if dtc.dts.CheckDictTypeUnique(dictType.DictId, dictType.DictType) {
	//	bzc.Waring("新增字典'" + dictType.DictName + "'失败，字典类型已存在")
	//	return
	//}
	//dictType.SetCreateBy(bzc.GetUserId())
	//dtc.dts.InsertDictType(dictType)
	//bzc.Success()
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
	//bzc := baizeContext.NewBaiZeContext(c)
	//dictType := new(models.SysDictTypeEdit)
	//if err := c.ShouldBindJSON(dictType); err != nil {
	//	bzc.ParameterError()
	//	return
	//}
	//if dtc.dts.CheckDictTypeUnique(dictType.DictId, dictType.DictType) {
	//	bzc.Waring("修改字典'" + dictType.DictName + "'失败，字典类型已存在")
	//	return
	//}
	//
	//dictType.SetUpdateBy(bzc.GetUserId())
	//dtc.dts.UpdateDictType(dictType)
	//bzc.Success()
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
	//bzc := baizeContext.NewBaiZeContext(c)
	//dictIds := bzc.ParamInt64Array("dictIds")
	////dictTypes := dtc.dts.SelectDictTypeByIds(dictIds)
	////if dtc.dts.CheckDictDataByTypes(dictTypes) {
	////	bzc.Waring("有已分配的字典,不能删除")
	////	return
	////}
	//dtc.dts.DeleteDictTypeByIds(dictIds)
	//bzc.Success()
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
	//bzc := baizeContext.NewBaiZeContext(c)
	//dtc.dts.DictTypeClearCache()
	//bzc.Success()
}

func (dtc *DictTypeController) DictTypeOptionselect(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//bzc.SuccessData(dtc.dts.SelectDictTypeAll())
}
