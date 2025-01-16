package toolDao

import (
	"baize/app/business/tool/toolModels"
	"context"
)

type IGenTableColumn interface {
	SelectDbTableColumnsByName(ctx context.Context, tableName string) (list []*toolModels.InformationSchemaColumn)
	SelectGenTableColumnListByTableId(ctx context.Context, tableId int64) (list []*toolModels.GenTableColumnVo)
	BatchInsertGenTableColumn(ctx context.Context, genTables []*toolModels.GenTableColumnDML)
	UpdateGenTableColumn(ctx context.Context, column *toolModels.GenTableColumnDML)
	DeleteGenTableColumnByIds(ctx context.Context, ids []int64)
}
