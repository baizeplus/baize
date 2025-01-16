package toolDao

import (
	"baize/app/business/tool/toolModels"
	"context"
)

type IGenTable interface {
	SelectGenTableList(ctx context.Context, table *toolModels.GenTableDQL) (list []*toolModels.GenTableVo, total int64)
	SelectDbTableList(ctx context.Context, table *toolModels.GenTableDQL) (list []*toolModels.DBTableVo, total int64)
	SelectDbTableListByNames(ctx context.Context, tableNames []string) (list []*toolModels.DBTableVo)
	SelectGenTableById(ctx context.Context, id int64) (table *toolModels.GenTableVo)
	SelectGenTableAll(ctx context.Context) (list []*toolModels.GenTableVo)
	BatchInsertGenTable(ctx context.Context, table []*toolModels.GenTableDML)
	InsertGenTable(ctx context.Context, table *toolModels.GenTableDML)
	UpdateGenTable(ctx context.Context, table *toolModels.GenTableDML)
	DeleteGenTableByIds(ctx context.Context, ids []int64)
}
