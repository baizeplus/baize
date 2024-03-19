package systemController

import (
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/business/system/systemService/systemServiceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type Config struct {
	cs systemService.IConfigService
}

func NewConfig(cs *systemServiceImpl.ConfigService) *Config {
	return &Config{cs: cs}
}

// ConfigList 查询配置列表查询
// @Summary 查询配置列表查询
// @Description 查询配置列表查询
// @Tags 配置相关
// @Param  object query systemModels.SysConfigDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ListData{Rows=[]systemModels.SysConfigVo}}  "成功"
// @Router /system/config/list  [get]
func (cc *Config) ConfigList(c *gin.Context) {
	config := new(systemModels.SysConfigDQL)
	_ = c.ShouldBind(config)
	list, count := cc.cs.SelectConfigList(c, config)
	baizeContext.SuccessListData(c, list, count)
}

// ConfigExport 导出配置
// @Summary 导出配置
// @Description 导出配置
// @Tags 配置相关
// @Param  object query systemModels.SysConfigDQL true "查询信息"
// @Security BearerAuth
// @Produce application/octet-stream
// @Success 200 {object} []byte
// @Router /system/config/export [post]
func (cc *Config) ConfigExport(c *gin.Context) {
	config := new(systemModels.SysConfigDQL)
	_ = c.ShouldBind(config)
	baizeContext.DataPackageExcel(c, cc.cs.ConfigExport(c, config))
}

// ConfigGetInfo 根据配置ID获取配置信息
// @Summary 根据配置ID获取配置信息
// @Description 根据配置ID获取配置信息
// @Tags 配置相关
// @Param id path string true "ConfigId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=systemModels.SysConfigVo}  "成功"
// @Router /system/config/{configId}  [get]
func (cc *Config) ConfigGetInfo(c *gin.Context) {
	ConfigId := baizeContext.ParamInt64(c, "configId")
	if ConfigId == 0 {
		baizeContext.ParameterError(c)
		return
	}
	baizeContext.SuccessData(c, cc.cs.SelectConfigById(c, ConfigId))
}

// ConfigAdd 添加配置
// @Summary 添加配置
// @Description 添加配置
// @Tags 配置相关
// @Param  object body systemModels.SysConfigVo true "公司信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/Config  [post]
func (cc *Config) ConfigAdd(c *gin.Context) {
	sysConfig := new(systemModels.SysConfigVo)
	if err := c.ShouldBindJSON(sysConfig); err != nil {
		baizeContext.ParameterError(c)
		return
	}
	if cc.cs.CheckConfigKeyUnique(c, 0, sysConfig.ConfigKey) {
		baizeContext.Waring(c, "添加配置'"+sysConfig.ConfigKey+"'失败，Key已存在")
		return
	}
	sysConfig.SetCreateBy(baizeContext.GetUserId(c))
	cc.cs.InsertConfig(c, sysConfig)
	baizeContext.Success(c)
}

// ConfigEdit 修改配置
// @Summary 修改配置
// @Description 修改配置
// @Tags 配置相关
// @Param  object body systemModels.SysConfigVo true "公司信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/Config  [put]
func (cc *Config) ConfigEdit(c *gin.Context) {
	sysConfig := new(systemModels.SysConfigVo)
	if err := c.ShouldBindJSON(sysConfig); err != nil {
		baizeContext.ParameterError(c)
		return
	}
	if cc.cs.CheckConfigKeyUnique(c, sysConfig.ConfigId, sysConfig.ConfigKey) {
		baizeContext.Waring(c, "修改配置'"+sysConfig.ConfigKey+"'失败，Key已存在")
		return
	}
	sysConfig.SetUpdateBy(baizeContext.GetUserId(c))
	cc.cs.UpdateConfig(c, sysConfig)
	baizeContext.Success(c)

}

// ConfigRemove 删除配置
// @Summary 删除配置
// @Description 删除配置
// @Tags 配置相关
// @Param ids path string true "configId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/Config/{configId} [delete]
func (cc *Config) ConfigRemove(c *gin.Context) {
	cc.cs.DeleteConfigById(c, baizeContext.ParamInt64(c, "configId"))
	baizeContext.Success(c)
}
