package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IConfigDao interface {
	SelectConfigList(ctx context.Context, db sqly.SqlyContext, Config *systemModels.SysConfigDQL) (sysConfigList []*systemModels.SysConfigVo, total *int64)
	SelectConfigById(ctx context.Context, db sqly.SqlyContext, ConfigId int64) (Config *systemModels.SysConfigVo)
	InsertConfig(ctx context.Context, db sqly.SqlyContext, Config *systemModels.SysConfigVo)
	UpdateConfig(ctx context.Context, db sqly.SqlyContext, Config *systemModels.SysConfigVo)
	DeleteConfigById(ctx context.Context, db sqly.SqlyContext, ConfigId int64)
}
