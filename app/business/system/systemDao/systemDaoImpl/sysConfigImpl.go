package systemDaoImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"errors"
	"github.com/baizeplus/sqly"
)

type sysConfigDao struct {
	ms        sqly.SqlyContext
	configSql string
}

func NewSysConfigDao(ms sqly.SqlyContext) systemDao.IConfigDao {
	return &sysConfigDao{
		ms:        ms,
		configSql: `select config_id,config_name,config_key,config_value,config_type,create_by,create_time,update_by,update_time,remark from sys_config`,
	}
}

func (s *sysConfigDao) SelectConfigList(ctx context.Context, config *systemModels.SysConfigDQL) (list []*systemModels.SysConfigVo, total int64) {
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
	err := s.ms.NamedSelectPageContext(ctx, &list, &total, s.configSql+whereSql, config)
	if err != nil {
		panic(err)
	}
	return
}
func (s *sysConfigDao) SelectConfigListAll(ctx context.Context, config *systemModels.SysConfigDQL) (list []*systemModels.SysConfigVo) {
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
	err := s.ms.NamedSelectContext(ctx, &list, s.configSql+whereSql, config)
	if err != nil {
		panic(err)
	}
	return
}

func (s *sysConfigDao) SelectConfigById(ctx context.Context, configId int64) (config *systemModels.SysConfigVo) {
	whereSql := ` where config_id = ?`
	config = new(systemModels.SysConfigVo)
	err := s.ms.GetContext(ctx, config, s.configSql+whereSql, configId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return config
}

func (s *sysConfigDao) InsertConfig(ctx context.Context, config *systemModels.SysConfigVo) {
	insertSQL := `insert into sys_config (config_id,config_name,config_key,config_value,config_type,remark,create_by,create_time,update_by,update_time)
					values (:config_id,:config_name,:config_key,:config_value,:config_type,:remark,:create_by,:create_time,:update_by,:update_time)`
	_, err := s.ms.NamedExecContext(ctx, insertSQL, config)
	if err != nil {
		panic(err)
	}
	return
}

func (s *sysConfigDao) UpdateConfig(ctx context.Context, config *systemModels.SysConfigVo) {
	updateSQL := `update sys_config set  update_time =:update_time , update_by = :update_by `

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

	_, err := s.ms.NamedExecContext(ctx, updateSQL, config)
	if err != nil {
		panic(err)
	}
	return
}

func (s *sysConfigDao) DeleteConfigById(ctx context.Context, configId int64) {
	_, err := s.ms.ExecContext(ctx, "delete from sys_config  where config_id = ?", configId)
	if err != nil {
		panic(err)
	}
	return
}

func (s *sysConfigDao) SelectConfigIdByConfigKey(ctx context.Context, configKey string) int64 {
	var configId int64 = 0
	err := s.ms.GetContext(ctx, &configId, "select config_id from sys_config where config_key = ?", configKey)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return configId
}

func (s *sysConfigDao) SelectConfigValueByConfigKey(ctx context.Context, configKey string) string {
	var configValue = ""
	err := s.ms.GetContext(ctx, &configValue, "select config_value from sys_config where config_key = ?", configKey)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return configValue
}
