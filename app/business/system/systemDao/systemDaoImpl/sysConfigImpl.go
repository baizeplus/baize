package systemDaoImpl

import (
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"errors"
	"github.com/baizeplus/sqly"
)

type SysConfigDao struct {
	configSql string
}

func NewSysConfigDao() *SysConfigDao {
	return &SysConfigDao{
		configSql: `select config_id,config_name,config_key,config_value,config_type,create_by,create_time,update_by,update_time,remark from sys_config`,
	}
}

func (s *SysConfigDao) SelectConfigList(ctx context.Context, db sqly.SqlyContext, config *systemModels.SysConfigDQL) (list []*systemModels.SysConfigVo, total int64) {
	whereSql := ``
	if config.ConfigName != "" {
		whereSql += " AND config_name like concat('%', :config_name, '%')"
	}
	if config.ConfigType != "" {
		whereSql += " AND  config_type = :config_type"
	}
	if config.ConfigKey != "" {
		whereSql += " AND config_key like concat('%', :config_key, '%')"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	err := db.NamedSelectPageContext(ctx, &list, &total, s.configSql+whereSql, config)
	if err != nil {
		panic(err)
	}
	return
}
func (s *SysConfigDao) SelectConfigListAll(ctx context.Context, db sqly.SqlyContext, config *systemModels.SysConfigDQL) (list []*systemModels.SysConfigVo) {
	whereSql := ``
	if config.ConfigName != "" {
		whereSql += " AND config_name like concat('%', :config_name, '%')"
	}
	if config.ConfigType != "" {
		whereSql += " AND  config_type = :config_type"
	}
	if config.ConfigKey != "" {
		whereSql += " AND post_name like concat('%', :config_key, '%')"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	list = make([]*systemModels.SysConfigVo, 0)
	err := db.NamedSelectContext(ctx, &list, s.configSql+whereSql, config)
	if err != nil {
		panic(err)
	}
	return
}

func (s *SysConfigDao) SelectConfigById(ctx context.Context, db sqly.SqlyContext, configId int64) (config *systemModels.SysConfigVo) {
	whereSql := ` where config_id = ?`
	config = new(systemModels.SysConfigVo)
	err := db.GetContext(ctx, config, s.configSql+whereSql, configId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return config
}

func (s *SysConfigDao) InsertConfig(ctx context.Context, db sqly.SqlyContext, config *systemModels.SysConfigVo) {
	insertSQL := `insert into sys_config (config_id,config_name,config_key,config_value,config_type,remark,create_by,create_time,update_by,update_time)
					values (:config_id,:config_name,:config_key,:config_value,:config_type,:remark,:create_by,now(),:update_by,now())`
	_, err := db.NamedExecContext(ctx, insertSQL, config)
	if err != nil {
		panic(err)
	}
	return
}

func (s *SysConfigDao) UpdateConfig(ctx context.Context, db sqly.SqlyContext, config *systemModels.SysConfigVo) {
	updateSQL := `update sys_config set  update_time = now() , update_by = :update_by `

	if config.ConfigName != "" {
		updateSQL += ",config_name = :config_name"
	}
	if config.ConfigKey != "" {
		updateSQL += ",config_key = :config_key"
	}
	if config.ConfigValue != "" {
		updateSQL += ",config_value = :config_value"
	}
	if config.ConfigType != "" {
		updateSQL += ",config_type = :config_type"
	}
	if config.Remark != "" {
		updateSQL += ",remark = :remark"
	}
	updateSQL += " where config_id = :config_id"

	_, err := db.NamedExecContext(ctx, updateSQL, config)
	if err != nil {
		panic(err)
	}
	return
}

func (s *SysConfigDao) DeleteConfigById(ctx context.Context, db sqly.SqlyContext, configId int64) {
	_, err := db.ExecContext(ctx, "delete from sys_config  where config_id = ?", configId)
	if err != nil {
		panic(err)
	}
	return
}

func (s *SysConfigDao) SelectConfigIdByConfigKey(ctx context.Context, db sqly.SqlyContext, configKey string) int64 {
	var configId int64 = 0
	err := db.GetContext(ctx, &configId, "select config_id from sys_config where config_key = ?", configKey)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return configId
}

func (s *SysConfigDao) SelectConfigValueByConfigKey(ctx context.Context, db sqly.SqlyContext, configKey string) string {
	var configValue = ""
	err := db.GetContext(ctx, &configValue, "select config_value from sys_config where config_key = ?", configKey)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return configValue
}
