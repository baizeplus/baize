package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IConfigDao interface {
	SelectConfigList(ctx context.Context, db sqly.SqlyContext, config *systemModels.SysConfigDQL) (sysConfigList []*systemModels.SysConfigVo, total int64)
	SelectConfigListAll(ctx context.Context, db sqly.SqlyContext, config *systemModels.SysConfigDQL) (list []*systemModels.SysConfigVo)
	SelectConfigById(ctx context.Context, db sqly.SqlyContext, configId int64) (Config *systemModels.SysConfigVo)
	InsertConfig(ctx context.Context, db sqly.SqlyContext, config *systemModels.SysConfigVo)
	UpdateConfig(ctx context.Context, db sqly.SqlyContext, config *systemModels.SysConfigVo)
	DeleteConfigById(ctx context.Context, db sqly.SqlyContext, configId int64)

	SelectConfigIdByConfigKey(ctx context.Context, db sqly.SqlyContext, configKey string) int64
	SelectConfigValueByConfigKey(ctx context.Context, db sqly.SqlyContext, configKey string) string
}
