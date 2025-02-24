package toolDaoImpl

import (
	"baize/app/business/tool/toolDao"
	"baize/app/business/tool/toolModels"
	"context"
	"github.com/baizeplus/sqly"
)

type GenTableColumnDao struct {
	ms sqly.SqlyContext
}

func NewGenTableColumnDao(ms sqly.SqlyContext) toolDao.IGenTableColumn {
	return &GenTableColumnDao{ms: ms}
}

func (genTableColumnDao *GenTableColumnDao) SelectDbTableColumnsByName(ctx context.Context, tableName string) (list []*toolModels.InformationSchemaColumn) {
	list = make([]*toolModels.InformationSchemaColumn, 0, 0)
	genTableColumnDao.ms.SelectContext(ctx, &list, `
		select column_name, (case when (is_nullable = 'no'  &&  column_key != 'PRI') then '1' else '0' end) as is_required, (case when column_key = 'PRI' then '1' else '0' end) as is_pk, ordinal_position as sort, column_comment,  column_type
		from information_schema.columns where table_schema = (select database()) and table_name = (?)
		order by ordinal_position
`, tableName)

	return
}

func (genTableColumnDao *GenTableColumnDao) SelectGenTableColumnListByTableId(ctx context.Context, tableId int64) (list []*toolModels.GenTableColumnVo) {
	list = make([]*toolModels.GenTableColumnVo, 0)
	err := genTableColumnDao.ms.SelectContext(ctx, &list, `select column_id, table_id, column_name, column_comment, column_type, go_type, go_field,html_field, is_pk,  is_required, is_insert, is_edit, is_list, is_query, query_type, html_type, dict_type, sort, create_by, create_time, update_by, update_time 
	from gen_table_column where table_id = ? order by sort`, tableId)
	if err != nil {
		panic(err)
	}
	return
}

func (genTableColumnDao *GenTableColumnDao) BatchInsertGenTableColumn(ctx context.Context, genTables []*toolModels.GenTableColumnDML) {
	_, err := genTableColumnDao.ms.NamedExecContext(ctx, `insert into gen_table_column(column_id,table_id,column_name,column_comment,column_type,go_type,go_field,html_field,is_pk,is_required,is_insert,is_edit,is_list, is_query,query_type, html_type, dict_type, sort,create_by,create_time,update_by,update_time)
							values(:column_id,:table_id,:column_name,:column_comment,:column_type,:go_type,:go_field,:html_field,:is_pk,:is_required,:is_insert,:is_edit,:is_list, :is_query,  :query_type, :html_type, :dict_type, :sort,:create_by,:create_time,:update_by,:update_time)`,
		genTables)
	if err != nil {
		panic(err)
	}
}

func (genTableColumnDao *GenTableColumnDao) UpdateGenTableColumn(ctx context.Context, column *toolModels.GenTableColumnDML) {
	_, err := genTableColumnDao.ms.NamedExecContext(ctx, "update gen_table_column set column_comment=:column_comment,go_type=:go_type,go_field=:go_field,html_field=:html_field,is_insert=:is_insert, is_edit=:is_edit,is_list=:is_list,is_query=:is_query,is_required=:is_required,query_type=:query_type,html_type=:html_type,dict_type=:dict_type,sort=:sort, update_by = :update_by,update_time = :update_time  where column_id = :column_id", column)
	if err != nil {
		panic(err)
	}
}

func (genTableColumnDao *GenTableColumnDao) DeleteGenTableColumnByIds(ctx context.Context, ids []int64) {
	query, i, err := sqly.In("  delete from gen_table_column where table_id in (?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = genTableColumnDao.ms.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}
