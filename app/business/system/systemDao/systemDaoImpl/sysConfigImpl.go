package systemDaoImpl

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type SysConfigDao struct {
	configSql string
}

func (s *SysConfigDao) SelectConfigList(ctx context.Context, db sqly.SqlyContext, Config *systemModels.SysConfigDQL) (sysConfigList []*systemModels.SysConfigVo, total *int64) {
	//TODO implement me
	panic("implement me")
}

func (s *SysConfigDao) SelectConfigById(ctx context.Context, db sqly.SqlyContext, ConfigId int64) (Config *systemModels.SysConfigVo) {
	//TODO implement me
	panic("implement me")
}

func (s *SysConfigDao) InsertConfig(ctx context.Context, db sqly.SqlyContext, Config *systemModels.SysConfigVo) {
	//TODO implement me
	panic("implement me")
}

func (s *SysConfigDao) UpdateConfig(ctx context.Context, db sqly.SqlyContext, Config *systemModels.SysConfigVo) {
	//TODO implement me
	panic("implement me")
}

func (s *SysConfigDao) DeleteConfigById(ctx context.Context, db sqly.SqlyContext, ConfigId int64) {
	//TODO implement me
	panic("implement me")
}

func NewSysConfigDao() *SysConfigDao {
	return &SysConfigDao{
		configSql: `select d.dept_id, d.parent_id, d.ancestors, d.dept_name, d.order_num, d.leader, d.phone, d.email, d.status, d.del_flag, d.create_by, d.create_time from sys_dept d`,
	}
}
