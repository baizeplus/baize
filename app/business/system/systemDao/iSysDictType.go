package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IDictTypeDao interface {
	SelectDictTypeList(ctx context.Context, db sqly.SqlyContext, dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo, total int64)
	SelectDictTypeAll(ctx context.Context, db sqly.SqlyContext, dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo)
	SelectDictTypeById(ctx context.Context, db sqly.SqlyContext, dictId int64) (dictType *systemModels.SysDictTypeVo)
	SelectDictTypeByIds(ctx context.Context, db sqly.SqlyContext, dictId []int64) (dictTypes []string)
	InsertDictType(ctx context.Context, db sqly.SqlyContext, dictType *systemModels.SysDictTypeVo)
	UpdateDictType(ctx context.Context, db sqly.SqlyContext, dictType *systemModels.SysDictTypeVo)
	DeleteDictTypeByIds(ctx context.Context, db sqly.SqlyContext, dictIds []int64)
	CheckDictTypeUnique(ctx context.Context, db sqly.SqlyContext, dictType string) int64
}
