package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IDictDataDao interface {
	SelectDictDataByType(ctx context.Context, dictType string) (SysDictDataList []*systemModels.SysDictDataVo)
	SelectDictDataList(ctx context.Context, dictData *systemModels.SysDictDataDQL) (list []*systemModels.SysDictDataVo, total int64)
	SelectDictDataById(ctx context.Context, dictCode string) (dictData *systemModels.SysDictDataVo)
	InsertDictData(ctx context.Context, dictData *systemModels.SysDictDataVo)
	UpdateDictData(ctx context.Context, dictData *systemModels.SysDictDataVo)
	SelectDictTypesByDictCodes(ctx context.Context, dictCodes []string) []string
	DeleteDictDataByIds(ctx context.Context, dictCodes []string)
	CountDictDataByTypes(ctx context.Context, dictType []string) int
}
