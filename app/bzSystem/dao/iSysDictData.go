package systemDao

import (
	"baize/app/bzSystem/models"
	"context"
	"github.com/baizeplus/sqly"
)

type IDictDataDao interface {
	SelectDictDataByType(ctx context.Context, db sqly.SqlyContext, dictType string) (SysDictDataList []*models.SysDictDataVo)
	SelectDictDataList(ctx context.Context, db sqly.SqlyContext, dictData *models.SysDictDataDQL) (list []*models.SysDictDataVo, total *int64)
	SelectDictDataById(ctx context.Context, db sqly.SqlyContext, dictCode int64) (dictData *models.SysDictDataVo)
	InsertDictData(ctx context.Context, db sqly.SqlyContext, dictData *models.SysDictDataVo)
	UpdateDictData(ctx context.Context, db sqly.SqlyContext, dictData *models.SysDictDataVo)
	DeleteDictDataByIds(ctx context.Context, db sqly.SqlyContext, dictCodes []int64)
	CountDictDataByTypes(ctx context.Context, db sqly.SqlyContext, dictType []string) int
}
