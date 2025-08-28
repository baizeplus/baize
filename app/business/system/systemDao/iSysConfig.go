package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IConfigDao interface {
	SelectConfigList(ctx context.Context, config *systemModels.SysConfigDQL) (sysConfigList []*systemModels.SysConfigVo, total int64)
	SelectConfigListAll(ctx context.Context, config *systemModels.SysConfigDQL) (list []*systemModels.SysConfigVo)
	SelectConfigById(ctx context.Context, configId string) (Config *systemModels.SysConfigVo)
	InsertConfig(ctx context.Context, config *systemModels.SysConfigVo)
	UpdateConfig(ctx context.Context, config *systemModels.SysConfigVo)
	DeleteConfigById(ctx context.Context, configId string)
	SelectConfigIdByConfigKey(ctx context.Context, configKey string) string
	SelectConfigValueByConfigKey(ctx context.Context, configKey string) string
}
