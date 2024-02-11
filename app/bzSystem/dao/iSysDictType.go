package dao

import (
	"baize/app/bzSystem/models"
	"context"
	"github.com/baizeplus/sqly"
)

type IDictTypeDao interface {
	SelectDictTypeList(ctx context.Context, db sqly.SqlyContext, dictType *models.SysDictTypeDQL) (list []*models.SysDictTypeVo, total *int64)
	SelectDictTypeAll(ctx context.Context, db sqly.SqlyContext) (list []*models.SysDictTypeVo)
	SelectDictTypeById(ctx context.Context, db sqly.SqlyContext, dictId int64) (dictType *models.SysDictTypeVo)
	SelectDictTypeByIds(ctx context.Context, db sqly.SqlyContext, dictId []int64) (dictTypes []string)
	InsertDictType(ctx context.Context, db sqly.SqlyContext, dictType *models.SysDictTypeVo)
	UpdateDictType(ctx context.Context, db sqly.SqlyContext, dictType *models.SysDictTypeVo)
	DeleteDictTypeByIds(ctx context.Context, db sqly.SqlyContext, dictIds []int64) (err error)
	CheckDictTypeUnique(ctx context.Context, db sqly.SqlyContext, dictType string) int64
}
