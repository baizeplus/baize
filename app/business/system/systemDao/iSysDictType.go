package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IDictTypeDao interface {
	SelectDictTypeList(ctx context.Context, dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo, total int64)
	SelectDictTypeAll(ctx context.Context, dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo)
	SelectDictTypeById(ctx context.Context, dictId string) (dictType *systemModels.SysDictTypeVo)
	SelectDictTypeByIds(ctx context.Context, dictId []string) (dictTypes []string)
	InsertDictType(ctx context.Context, dictType *systemModels.SysDictTypeVo)
	UpdateDictType(ctx context.Context, dictType *systemModels.SysDictTypeVo)
	DeleteDictTypeByIds(ctx context.Context, dictIds []string)
	CheckDictTypeUnique(ctx context.Context, dictType string) string
}
