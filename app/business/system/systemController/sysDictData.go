package systemController

import (
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/middlewares"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
	"strings"
)

type DictData struct {
	dds systemService.IDictDataService
}

func NewDictData(dds systemService.IDictDataService) *DictData {
	return &DictData{
		dds: dds,
	}
}
func (ddc *DictData) PrivateRoutes(router *gin.RouterGroup) {
	systemDictData := router.Group("/system/dict/data")
	systemDictData.GET("/list", middlewares.HasPermission("system:dict"), ddc.DictDataList)
	systemDictData.GET("/export", middlewares.HasPermission("system:dict:export"), ddc.DictDataExport)
	systemDictData.GET("/:dictCode", middlewares.HasPermission("system:dict:query"), ddc.DictDataGetInfo)
	systemDictData.GET("/type/:dictType", ddc.DictDataType)
	systemDictData.POST("", middlewares.SetLog("字典数据管理", middlewares.Insert), middlewares.HasPermission("system:dict:add"), ddc.DictDataAdd)
	systemDictData.PUT("", middlewares.SetLog("字典数据管理", middlewares.Update), middlewares.HasPermission("system:dict:edit"), ddc.DictDataEdit)
	systemDictData.DELETE("/:dictCodes", middlewares.SetLog("字典数据管理", middlewares.Delete), middlewares.HasPermission("system:dict:remove"), ddc.DictDataRemove)

}

// DictDataList 查询字典列表
// @Summary 查询字典列表
// @Description 查询字典列表
// @Tags 字典相关
// @Param  object query systemModels.SysDictDataDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ListData{Rows=[]systemModels.SysDictDataVo}}  "成功"
// @Router /system/dict/data/list  [get]
func (ddc *DictData) DictDataList(c *gin.Context) {
	dictData := new(systemModels.SysDictDataDQL)
	_ = c.ShouldBind(dictData)
	list, count := ddc.dds.SelectDictDataList(c, dictData)
	baizeContext.SuccessListData(c, list, count)
}

// DictDataExport 导出配置
// @Summary 导出配置
// @Description 导出配置
// @Tags 配置相关
// @Param  object query systemModels.SysConfigDQL true "查询信息"
// @Security BearerAuth
// @Produce application/octet-stream
// @Success 200 {object} []byte
// @Router /system/config/export [post]
func (ddc *DictData) DictDataExport(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//dictData := new(systemModels.SysDictDataDQL)
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
// @Success 200 {object}  response.ResponseData{data=systemModels.SysDictDataVo}  "成功"
// @Router /system/dict/data/{dictCode}  [get]
func (ddc *DictData) DictDataGetInfo(c *gin.Context) {
	dictCode := c.Param("dictCode")
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
// @Success 200 {object}  response.ResponseData{data=response.ListData{Rows=[]systemModels.SysDictDataVo}}  "成功"
// @Router /system/dict/data/type/{dictType}  [get]
func (ddc *DictData) DictDataType(c *gin.Context) {
	sysDictDataList := ddc.dds.SelectDictDataByType(c, c.Param("dictType"))
	baizeContext.SuccessGzip(c, sysDictDataList)
}

// DictDataAdd 添加字典数据
// @Summary 添加字典数据
// @Description 添加字典数据
// @Tags 字典相关
// @Param  object body systemModels.SysDictDataVo true "字典"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/dict/data  [post]
func (ddc *DictData) DictDataAdd(c *gin.Context) {
	dictData := new(systemModels.SysDictDataVo)
	_ = c.ShouldBindJSON(dictData)
	dictData.SetCreateBy(baizeContext.GetUserId(c))
	ddc.dds.InsertDictData(c, dictData)
	baizeContext.Success(c)
}

// DictDataEdit 修改字典数据
// @Summary 修改字典数据
// @Description 修改字典数据
// @Tags 字典相关
// @Param  object body systemModels.SysDictDataVo true "字典"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/dict/data  [put]
func (ddc *DictData) DictDataEdit(c *gin.Context) {
	dictData := new(systemModels.SysDictDataVo)
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
// @Success 200 {object}  response.ResponseData{data=systemModels.SysDictDataVo}  "成功"
// @Router /system/dict/data/{dictCodes}  [delete]
func (ddc *DictData) DictDataRemove(c *gin.Context) {
	ddc.dds.DeleteDictDataByIds(c, strings.Split(c.Param("dictCodes"), ","))
	baizeContext.Success(c)
}
