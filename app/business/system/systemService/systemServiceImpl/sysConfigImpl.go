package systemServiceImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemDao/systemDaoImpl"
	"baize/app/business/system/systemModels"
	"baize/app/datasource"
	"baize/app/utils/snowflake"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
)

type ConfigService struct {
	data *sqly.DB
	cd   systemDao.IConfigDao
}

func NewConfigService(data *sqly.DB, cd *systemDaoImpl.SysConfigDao) *ConfigService {
	return &ConfigService{data: data, cd: cd}
}

func (cs *ConfigService) SelectConfigList(c *gin.Context, config *systemModels.SysConfigDQL) (sysConfigList []*systemModels.SysConfigVo, total *int64) {
	return cs.cd.SelectConfigList(c, cs.data, config)
}

func (cs *ConfigService) SelectConfigById(c *gin.Context, configId int64) (Config *systemModels.SysConfigVo) {
	return cs.cd.SelectConfigById(c, cs.data, configId)
}

func (cs *ConfigService) InsertConfig(c *gin.Context, config *systemModels.SysConfigVo) {
	config.ConfigId = snowflake.GenID()
	cs.cd.InsertConfig(c, cs.data, config)
}

func (cs *ConfigService) UpdateConfig(c *gin.Context, config *systemModels.SysConfigVo) {
	cs.cd.UpdateConfig(c, cs.data, config)
	datasource.RedisDb.Del(c, cs.getCacheKey(config.ConfigKey))
}

func (cs *ConfigService) DeleteConfigById(c *gin.Context, configId int64) {
	key := cs.cd.SelectConfigById(c, cs.data, configId).ConfigKey
	cs.cd.DeleteConfigById(c, cs.data, configId)
	datasource.RedisDb.Del(c, cs.getCacheKey(key))
}

func (cs *ConfigService) CheckConfigKeyUnique(c *gin.Context, configId int64, configKey string) bool {
	id := cs.cd.SelectConfigIdByConfigKey(c, cs.data, configKey)
	if id == configId {
		return false
	}
	return true
}

func (cs *ConfigService) SelectConfigValueByKey(c *gin.Context, configKey string) string {
	value := datasource.RedisDb.Get(c, cs.getCacheKey(configKey)).Val()
	if value != "" {
		return value
	}
	value = cs.cd.SelectConfigValueByConfigKey(c, cs.data, configKey)
	if value != "" {
		datasource.RedisDb.Set(c, cs.getCacheKey(configKey), value, -1)
	}
	return value
}

func (cs *ConfigService) getCacheKey(configKey string) string {
	return "sys_config:" + configKey

}
