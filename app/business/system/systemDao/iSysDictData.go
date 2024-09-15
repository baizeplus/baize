package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IDictDataDao interface {
	SelectDictDataByType(ctx context.Context, db sqly.SqlyContext, dictType string) (SysDictDataList []*systemModels.SysDictDataVo)
	SelectDictDataList(ctx context.Context, db sqly.SqlyContext, dictData *systemModels.SysDictDataDQL) (list []*systemModels.SysDictDataVo, total int64)
	SelectDictDataById(ctx context.Context, db sqly.SqlyContext, dictCode int64) (dictData *systemModels.SysDictDataVo)
	InsertDictData(ctx context.Context, db sqly.SqlyContext, dictData *systemModels.SysDictDataVo)
	UpdateDictData(ctx context.Context, db sqly.SqlyContext, dictData *systemModels.SysDictDataVo)
	SelectDictTypesByDictCodes(ctx context.Context, db sqly.SqlyContext, dictCodes []int64) []string
	DeleteDictDataByIds(ctx context.Context, db sqly.SqlyContext, dictCodes []int64)
	CountDictDataByTypes(ctx context.Context, db sqly.SqlyContext, dictType []string) int
}
