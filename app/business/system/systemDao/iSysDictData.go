package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IDictDataDao interface {
	SelectDictDataByType(ctx context.Context, dictType string) (SysDictDataList []*systemModels.SysDictDataVo)
	SelectDictDataList(ctx context.Context, dictData *systemModels.SysDictDataDQL) (list []*systemModels.SysDictDataVo, total int64)
	SelectDictDataById(ctx context.Context, dictCode int64) (dictData *systemModels.SysDictDataVo)
	InsertDictData(ctx context.Context, dictData *systemModels.SysDictDataVo)
	UpdateDictData(ctx context.Context, dictData *systemModels.SysDictDataVo)
	SelectDictTypesByDictCodes(ctx context.Context, dictCodes []int64) []string
	DeleteDictDataByIds(ctx context.Context, dictCodes []int64)
	CountDictDataByTypes(ctx context.Context, dictType []string) int
}
