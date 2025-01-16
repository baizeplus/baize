package systemServiceImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/datasource/cache"
	"baize/app/utils/excel"
	"baize/app/utils/snowflake"
	"github.com/gin-gonic/gin"
)

type ConfigService struct {
	cache cache.Cache
	cd    systemDao.IConfigDao
}

func NewConfigService(cd systemDao.IConfigDao,
	cache cache.Cache) systemService.IConfigService {
	return &ConfigService{cache: cache, cd: cd}
}

func (cs *ConfigService) SelectConfigList(c *gin.Context, config *systemModels.SysConfigDQL) (sysConfigList []*systemModels.SysConfigVo, total int64) {
	return cs.cd.SelectConfigList(c, config)
}
func (cs *ConfigService) ConfigExport(c *gin.Context, config *systemModels.SysConfigDQL) (data []byte) {
	list := cs.cd.SelectConfigListAll(c, config)
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

func (cs *ConfigService) SelectConfigById(c *gin.Context, configId int64) (Config *systemModels.SysConfigVo) {
	return cs.cd.SelectConfigById(c, configId)
}

func (cs *ConfigService) InsertConfig(c *gin.Context, config *systemModels.SysConfigVo) {
	config.ConfigId = snowflake.GenID()
	cs.cd.InsertConfig(c, config)
}

func (cs *ConfigService) UpdateConfig(c *gin.Context, config *systemModels.SysConfigVo) {
	cs.cd.UpdateConfig(c, config)
	cs.cache.Del(c, cs.getCacheKey(config.ConfigKey))
}

func (cs *ConfigService) DeleteConfigById(c *gin.Context, configId int64) {
	key := cs.cd.SelectConfigById(c, configId).ConfigKey
	cs.cd.DeleteConfigById(c, configId)
	cs.cache.Del(c, cs.getCacheKey(key))
}

func (cs *ConfigService) CheckConfigKeyUnique(c *gin.Context, configId int64, configKey string) bool {
	id := cs.cd.SelectConfigIdByConfigKey(c, configKey)
	if id == configId {
		return false
	}
	return true
}

func (cs *ConfigService) SelectConfigValueByKey(c *gin.Context, configKey string) string {
	v, err := cs.cache.Get(c, cs.getCacheKey(configKey))
	if err == nil {
		return v
	}

	value := cs.cd.SelectConfigValueByConfigKey(c, configKey)
	if value != "" {
		cs.cache.Set(c, cs.getCacheKey(configKey), value, 0)
	}
	return value
}

func (cs *ConfigService) getCacheKey(configKey string) string {
	return "sys_config:" + configKey

}
