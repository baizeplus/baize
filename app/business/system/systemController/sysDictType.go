package systemController

import (
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/middlewares"
	"baize/app/utils/baizeContext"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DictType struct {
	dts systemService.IDictTypeService
	dds systemService.IDictDataService
}

func NewDictType(dts systemService.IDictTypeService, dds systemService.IDictDataService) *DictType {
	return &DictType{
		dts: dts,
		dds: dds,
	}
}
func (dtc *DictType) PrivateRoutes(router *gin.RouterGroup) {
	systemDictType := router.Group("/system/dict/type")
	systemDictType.GET("/list", middlewares.HasPermission("system:dict"), dtc.DictTypeList)
	systemDictType.POST("/export", middlewares.HasPermission("system:dict:export"), dtc.DictTypeExport)
	systemDictType.GET("/:dictId", middlewares.HasPermission("system:dict:query"), dtc.DictTypeGetInfo)
	systemDictType.POST("", middlewares.SetLog("字典管理", middlewares.Insert), middlewares.HasPermission("system:dict:add"), dtc.DictTypeAdd)
	systemDictType.PUT("", middlewares.SetLog("字典管理", middlewares.Update), middlewares.HasPermission("system:dict:edit"), dtc.DictTypeEdit)
	systemDictType.DELETE("/:dictIds", middlewares.SetLog("字典管理", middlewares.Delete), middlewares.HasPermission("system:dict:remove"), dtc.DictTypeRemove)
	systemDictType.DELETE("/refreshCache", middlewares.HasPermission("system:dict:remove"), dtc.DictTypeClearCache)
	systemDictType.GET("/optionSelect", dtc.DictTypeOptionSelect)
}

// DictTypeList 查询字典类型列表
// @Summary 查询字典类型列表
// @Description 查询字典类型列表
// @Tags 字典相关
// @Param  object query systemModels.SysDictDataDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ListData{Rows=[]systemModels.SysDictDataVo}}  "成功"
// @Router /system/dict/type/list  [get]
func (dtc *DictType) DictTypeList(c *gin.Context) {
	dictType := new(systemModels.SysDictTypeDQL)
	_ = c.ShouldBind(dictType)
	list, count := dtc.dts.SelectDictTypeList(c, dictType)
	baizeContext.SuccessListData(c, list, count)

}

// DictTypeExport 导出字典类型
// @Summary 导出字典类型
// @Description 导出字典类型
// @Tags 字典相关
// @Param  object query systemModels.SysDictDataDQL true "查询信息"
// @Security BearerAuth
// @Produce application/octet-stream
// @Success 200 {object} []byte
// @Router /system/config/export [post]
func (dtc *DictType) DictTypeExport(c *gin.Context) {
	dictType := new(systemModels.SysDictTypeDQL)
	_ = c.ShouldBind(dictType)
	baizeContext.DataPackageExcel(c, dtc.dts.ExportDictType(c, dictType))
}

// DictTypeGetInfo 根据dictCode获取字典类型类型
// @Summary 根据dictCode获取字典类型信息
// @Description 根据dictCode获取字典类型信息
// @Tags 字典相关
// @Param id path string true "dictCode"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=systemModels.SysDictDataVo}  "成功"
// @Router /system/dict/type/{dictCode}  [get]
func (dtc *DictType) DictTypeGetInfo(c *gin.Context) {
	dictId := c.Param("dictId")
	if dictId == "" {
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
// @Param  object body systemModels.SysDictDataVo true "字典"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/dict/type  [post]
func (dtc *DictType) DictTypeAdd(c *gin.Context) {
	dictType := new(systemModels.SysDictTypeVo)
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
// @Param  object body systemModels.SysDictDataVo true "字典"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/dict/type  [put]
func (dtc *DictType) DictTypeEdit(c *gin.Context) {
	dictType := new(systemModels.SysDictTypeVo)
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
func (dtc *DictType) DictTypeRemove(c *gin.Context) {
	dictIds := strings.Split(c.Param("dictIds"), ",")
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
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/dict/type/refreshCache  [put]
func (dtc *DictType) DictTypeClearCache(c *gin.Context) {
	dtc.dts.DictTypeClearCache(c)
	baizeContext.Success(c)
}

// DictTypeOptionSelect 查询字典列表
// @Summary 查询字典列表
// @Description 查询字典列表
// @Tags 字典相关
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=systemModels.SysDictDataVo}  "成功"
// @Router /system/dict/type/optionSelect  [put]
func (dtc *DictType) DictTypeOptionSelect(c *gin.Context) {
	baizeContext.SuccessData(c, dtc.dts.SelectDictTypeAll(c))
}
