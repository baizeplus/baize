package toolDao

import (
	"baize/app/business/tool/toolModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IGenTableColumn interface {
	SelectDbTableColumnsByName(ctx context.Context, db sqly.SqlyContext, tableName string) (list []*toolModels.InformationSchemaColumn)
	SelectGenTableColumnListByTableId(ctx context.Context, db sqly.SqlyContext, tableId int64) (list []*toolModels.GenTableColumnVo)
	BatchInsertGenTableColumn(ctx context.Context, db sqly.SqlyContext, genTables []*toolModels.GenTableColumnDML)
	UpdateGenTableColumn(ctx context.Context, db sqly.SqlyContext, column *toolModels.GenTableColumnDML)
	DeleteGenTableColumnByIds(ctx context.Context, db sqly.SqlyContext, ids []int64)
}
