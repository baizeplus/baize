package toolDao

import (
	"baize/app/business/tool/toolModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IGenTable interface {
	SelectGenTableList(ctx context.Context, db sqly.SqlyContext, table *toolModels.GenTableDQL) (list []*toolModels.GenTableVo, total int64)
	SelectDbTableList(ctx context.Context, db sqly.SqlyContext, table *toolModels.GenTableDQL) (list []*toolModels.DBTableVo, total int64)
	SelectDbTableListByNames(ctx context.Context, db sqly.SqlyContext, tableNames []string) (list []*toolModels.DBTableVo)
	SelectGenTableById(ctx context.Context, db sqly.SqlyContext, id int64) (table *toolModels.GenTableVo)
	SelectGenTableAll(ctx context.Context, db sqly.SqlyContext) (list []*toolModels.GenTableVo)
	BatchInsertGenTable(ctx context.Context, db sqly.SqlyContext, table []*toolModels.GenTableDML)
	InsertGenTable(ctx context.Context, db sqly.SqlyContext, table *toolModels.GenTableDML)
	UpdateGenTable(ctx context.Context, db sqly.SqlyContext, table *toolModels.GenTableDML)
	DeleteGenTableByIds(ctx context.Context, db sqly.SqlyContext, ids []int64)
}
