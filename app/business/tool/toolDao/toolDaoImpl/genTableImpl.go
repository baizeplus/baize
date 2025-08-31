package toolDaoImpl

import (
	"baize/app/business/tool/toolDao"
	"baize/app/business/tool/toolModels"
	"context"

	"github.com/baizeplus/sqly"
)

type GenTableDao struct {
	ms sqly.SqlyContext
}

func GetGenTableDao(ms sqly.SqlyContext) toolDao.IGenTable {
	return &GenTableDao{ms: ms}
}

func (genTableDao *GenTableDao) SelectGenTableList(ctx context.Context, table *toolModels.GenTableDQL) (list []*toolModels.GenTableVo, total int64) {
	var selectSql = `select table_id, table_name, table_comment, sub_table_name, sub_table_fk_name, struct_name, tpl_category, package_name, module_name, business_name, function_name, function_author, options, create_by, create_time, update_by, update_time, remark  from gen_table`
	whereSql := ``
	if table.TableName != "" {
		whereSql += " AND lower(table_name) like lower(concat('%', :table_name, '%'))"
	}
	if table.TableComment != "" {
		whereSql += " AND lower(table_comment) like lower(concat('%', :table_comment, '%'))"
	}
	if table.BeginTime != "" {
		whereSql += " AND date_format(create_time,'%y%m%d') &gt;= date_format(:begin_time,'%y%m%d')"
	}
	if table.EndTime != "" {
		whereSql += " date_format(create_time,'%y%m%d') &lt;= date_format(:end_time,'%y%m%d')"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	err := genTableDao.ms.NamedSelectPageContext(ctx, &list, &total, selectSql+whereSql, table)
	if err != nil {
		panic(err)
	}
	return

}
func (genTableDao *GenTableDao) SelectDbTableList(ctx context.Context, table *toolModels.GenTableDQL) (list []*toolModels.DBTableVo, total int64) {
	var selectSql = `select table_name , table_comment, create_time, update_time  from information_schema.tables where table_schema = (select database())
		AND table_name NOT LIKE 'gen_%'
		AND table_name NOT IN (select table_name from gen_table)`
	if table.TableName != "" {
		selectSql += " AND lower(table_name) like lower(concat('%', :table_name, '%'))"
	}
	if table.TableComment != "" {
		selectSql += " AND lower(table_comment) like lower(concat('%', :table_comment, '%'))"
	}
	if table.BeginTime != "" {
		selectSql += " AND date_format(create_time,'%y%m%d') &gt;= date_format(:begin_time,'%y%m%d')"
	}
	if table.EndTime != "" {
		selectSql += " date_format(create_time,'%y%m%d') &lt;= date_format(:end_time,'%y%m%d')"
	}
	err := genTableDao.ms.NamedSelectPageContext(ctx, &list, &total, selectSql, table)
	if err != nil {
		panic(err)
	}
	return
}

func (genTableDao *GenTableDao) SelectDbTableListByNames(ctx context.Context, tableNames []string) (list []*toolModels.DBTableVo) {
	query, i, err := sqly.In("select table_name, table_comment, create_time, update_time from information_schema.tables where table_name NOT LIKE 'gen_%' and table_schema = (select database()) and table_name in  (?)", tableNames)
	if err != nil {
		panic(err)
	}
	list = make([]*toolModels.DBTableVo, 0)
	err = genTableDao.ms.SelectContext(ctx, &list, query, i...)
	if err != nil {
		panic(err)
	}
	return
}

func (genTableDao *GenTableDao) SelectGenTableById(ctx context.Context, id string) (genTable *toolModels.GenTableVo) {
	genTable = new(toolModels.GenTableVo)
	err := genTableDao.ms.GetContext(ctx, genTable, `SELECT
       table_id, table_name, table_comment, sub_table_name,sub_table_fk_name, struct_name, parent_menu_id,
      tpl_category, package_name,module_name, business_name,function_name, function_author, options, remark
		FROM gen_table 
		where table_id = ?`, id)
	if err != nil {
		panic(err)
	}
	return
}
func (genTableDao *GenTableDao) SelectGenTableByName(ctx context.Context, name string) (genTable *toolModels.GenTableVo) {
	genTable = new(toolModels.GenTableVo)
	err := genTableDao.ms.GetContext(ctx, genTable, `SELECT t.table_id, t.table_name, t.table_comment, t.sub_table_name, t.sub_table_fk_name, t.struct_name,t.tpl_category, t.package_name, t.module_name, t.business_name, t.function_name, t.function_author,  t.options, t.remark
		FROM gen_table t
		where t.table_name = ? `, name)
	if err != nil {
		panic(err)
	}
	return
}
func (genTableDao *GenTableDao) SelectGenTableAll(ctx context.Context) (list []*toolModels.GenTableVo) {
	list = make([]*toolModels.GenTableVo, 0)
	err := genTableDao.ms.SelectContext(ctx, &list, `SELECT t.table_id, t.table_name, t.table_comment, t.sub_table_name, t.sub_table_fk_name, t.struct_name,t.tpl_category, t.package_name, t.module_name, t.business_name, t.function_name, t.function_author, t.options, t.remark
		FROM gen_table t`)
	if err != nil {
		panic(err)
	}
	return
}

func (genTableDao *GenTableDao) BatchInsertGenTable(ctx context.Context, genTables []*toolModels.GenTableDML) {

	_, err := genTableDao.ms.NamedExecContext(ctx, `insert into gen_table(table_id,table_name,table_comment,struct_name,tpl_category,package_name,module_name,business_name,function_name,function_author,create_by,create_time,update_by,update_time,remark)
							values(:table_id,:table_name,:table_comment,:struct_name,:tpl_category,:package_name,:module_name,:business_name,:function_name,:function_author,:create_by,:create_time,:update_by,:update_time,:remark)`,
		genTables)
	if err != nil {
		panic(err)
	}

}

func (genTableDao *GenTableDao) InsertGenTable(ctx context.Context, genTable *toolModels.GenTableDML) {
	insertSQL := `insert into gen_table(table_id,table_name,table_comment,struct_name,tpl_category,package_name,module_name,business_name,function_name,function_author,create_by,create_time,update_by,update_time,remark)
					values(:table_id,:table_name,:table_comment,:struct_name,:tpl_category,:package_name,:module_name,:business_name,:function_name,:function_author,:create_by,:create_time,:update_by,:update_time,:remark)`

	_, err := genTableDao.ms.NamedExecContext(ctx, insertSQL, genTable)
	if err != nil {
		panic(err)
	}
	return
}

func (genTableDao *GenTableDao) UpdateGenTable(ctx context.Context, genTable *toolModels.GenTableDML) {
	updateSQL := `update gen_table set update_time = :update_time , update_by = :update_by`
	if genTable.TableName != "" {
		updateSQL += ",table_name = :table_name"
	}
	if genTable.TableComment != "" {
		updateSQL += ",table_comment = :table_comment"
	}
	if genTable.SubTableName != "" {
		updateSQL += ",sub_table_name = :sub_table_name"
	}
	if genTable.SubTableFkName != "" {
		updateSQL += ",sub_table_fk_name = :sub_table_fk_name"
	}
	if genTable.StructName != "" {
		updateSQL += ",struct_name = :struct_name"
	}
	if genTable.FunctionAuthor != "" {
		updateSQL += ",function_author = :function_author"
	}
	if genTable.TplCategory != "" {
		updateSQL += ",tpl_category = :tpl_category"
	}
	if genTable.PackageName != "" {
		updateSQL += ",package_name = :package_name"
	}
	if genTable.ModuleName != "" {
		updateSQL += ",module_name = :module_name"
	}
	if genTable.BusinessName != "" {
		updateSQL += ",business_name = :business_name"
	}
	if genTable.FunctionName != "" {
		updateSQL += ",function_name = :function_name"
	}
	if genTable.Options != "" {
		updateSQL += ",options = :options"
	}
	if genTable.Remark != "" {
		updateSQL += ",remark = :remark"
	}
	if genTable.ParentMenuId != "" {
		updateSQL += ",parent_menu_id = :parent_menu_id"
	}

	updateSQL += " where table_id = :table_id"

	_, err := genTableDao.ms.NamedExecContext(ctx, updateSQL, genTable)
	if err != nil {
		panic(err)
	}
	return
}

func (genTableDao *GenTableDao) DeleteGenTableByIds(ctx context.Context, ids []string) {
	query, i, err := sqly.In(" delete from gen_table where table_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = genTableDao.ms.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}

}
