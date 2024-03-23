package toolDaoImpl

import (
	"baize/app/business/tool/toolModels"
	"context"
	"github.com/baizeplus/sqly"
)

type GenTableColumnDao struct {
}

func NewGenTableColumnDao() *GenTableColumnDao {
	return &GenTableColumnDao{}
}

func (genTableColumnDao *GenTableColumnDao) SelectDbTableColumnsByName(ctx context.Context, db sqly.SqlyContext, tableName string) (list []*toolModels.InformationSchemaColumn) {
	list = make([]*toolModels.InformationSchemaColumn, 0, 0)
	db.SelectContext(ctx, &list, `
		select column_name, (case when (is_nullable = 'no'  &&  column_key != 'PRI') then '1' else '0' end) as is_required, (case when column_key = 'PRI' then '1' else '0' end) as is_pk, ordinal_position as sort, column_comment,  column_type
		from information_schema.columns where table_schema = (select database()) and table_name = (?)
		order by ordinal_position
`, tableName)

	return
}

func (genTableColumnDao *GenTableColumnDao) SelectGenTableColumnListByTableId(ctx context.Context, db sqly.SqlyContext, tableId int64) (list []*toolModels.GenTableColumnVo) {
	list = make([]*toolModels.GenTableColumnVo, 0)
	db.SelectContext(ctx, &list, `select column_id, table_id, column_name, column_comment, column_type, go_type, go_field,html_field, is_pk,  is_required, is_insert, is_edit, is_list, is_query, query_type, html_type, dict_type, sort, create_by, create_time, update_by, update_time 
	from gen_table_column where table_id = ? order by sort`, tableId)
	return
}

func (genTableColumnDao *GenTableColumnDao) BatchInsertGenTableColumn(ctx context.Context, db sqly.SqlyContext, genTables []*toolModels.GenTableColumnDML) {
	_, err := db.NamedExecContext(ctx, `insert into gen_table_column(column_id,table_id,column_name,column_comment,column_type,go_type,go_field,html_field,is_pk,is_required,is_insert,is_edit,is_list, is_query,query_type, html_type, dict_type, sort,create_by,create_time,update_by,update_time)
							values(:column_id,:table_id,:column_name,:column_comment,:column_type,:go_type,:go_field,:html_field,:is_pk,:is_required,:is_insert,:is_edit,:is_list, :is_query,  :query_type, :html_type, :dict_type, :sort,:create_by,now(),:update_by,now())`,
		genTables)
	if err != nil {
		panic(err)
	}
}

func (genTableColumnDao *GenTableColumnDao) UpdateGenTableColumn(ctx context.Context, db sqly.SqlyContext, column *toolModels.GenTableColumnDML) {
	_, err := db.NamedExecContext(ctx, "update gen_table_column set column_comment=:column_comment,go_type=:go_type,go_field=:go_field,html_field=:html_field,is_insert=:is_insert, is_edit=:is_edit,is_list=:is_list,is_query=:is_query,is_required=:is_required,query_type=:query_type,html_type=:html_type,dict_type=:dict_type,sort=:sort, update_by = :update_by,update_time = now()  where column_id = :column_id", column)
	if err != nil {
		panic(err)
	}
}

func (genTableColumnDao *GenTableColumnDao) DeleteGenTableColumnByIds(ctx context.Context, db sqly.SqlyContext, ids []int64) {
	query, i, err := sqly.In("  delete from gen_table_column where table_id in (?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}

//func selectDbTableList(GenTable *toolModels.GenTableDQL) (list []*toolModels.GenTableVo, total *int64, err error) {
//	var selectSql = `select table_name, table_comment, create_time, update_time `
//	var fromSql = ` from information_schema.tables`
//	whereSql := ` where table_schema = (select database())
//		AND table_name NOT LIKE 'qrtz_%' AND table_name NOT LIKE 'gen_%'
//		AND table_name NOT IN (select table_name from gen_table)`
//	if GenTable.TableName != "" {
//		whereSql += " AND lower(table_name) like lower(concat('%', :table_name, '%'))"
//	}
//	if GenTable.TableComment != "" {
//		whereSql += " AND lower(table_comment) like lower(concat('%', :table_comment, '%'))"
//	}
//	if GenTable.BeginTime != "" {
//		whereSql += " AND date_format(create_time,'%y%m%d') &gt;= date_format(:begin_time,'%y%m%d')"
//	}
//	if GenTable.EndTime != "" {
//		whereSql += " date_format(create_time,'%y%m%d') &lt;= date_format(:end_time,'%y%m%d')"
//	}
//
//	if whereSql != "" {
//		whereSql = " where " + whereSql[4:]
//	}
//	countSql := constants.MysqlCount + fromSql + whereSql
//
//	countRow, err := mysql.GetMasterDb().NamedQuery(countSql, GenTable)
//	if err != nil {
//		return
//	}
//	total = new(int64)
//	if countRow.Next() {
//		countRow.Scan(total)
//	}
//	defer countRow.Close()
//	list = make([]*toolModels.GenTableVo, 0, GenTable.Size)
//	if *total > GenTable.Offset {
//		if GenTable.Limit != "" {
//			whereSql += GenTable.Limit
//		}
//		listRows, err := mysql.GetMasterDb().NamedQuery(selectSql+fromSql+whereSql, GenTable)
//		if err != nil {
//			return nil, nil, err
//		}
//		for listRows.Next() {
//			postVo := new(toolModels.GenTableVo)
//			listRows.StructScan(postVo)
//			list = append(list, postVo)
//		}
//		defer listRows.Close()
//	}
//	return
//}
//
//func SelectDbTableListByNames(tableNames []string) (list []*toolModels.GenTableVo,  err error)  {
//	query, i, err := sqlx.In("select table_name, table_comment, create_time, update_time from information_schema.tables where table_name NOT LIKE 'gen_%' and table_schema = (select database()) and table_name in  (?)", tableNames)
//	list = make([]*toolModels.GenTableVo, 0,0)
//	mysql.GetMasterDb().Select(&list,query,i...)
//	return
//}
//
//func SelectGenTableById(id int64) (genTable *toolModels.GenTableVo)  {
//	genTable=new(toolModels.GenTableVo)
//	mysql.GetMasterDb().Get(&genTable,`SELECT t.table_id, t.table_name, t.table_comment, t.sub_table_name, t.sub_table_fk_name, t.class_name, t.tpl_category, t.package_name, t.module_name, t.business_name, t.function_name, t.function_author, t.gen_type, t.gen_path, t.options, t.remark,
//			   c.column_id, c.column_name, c.column_comment, c.column_type, c.java_type, c.java_field, c.is_pk, c.is_increment, c.is_required, c.is_insert, c.is_edit, c.is_list, c.is_query, c.query_type, c.html_type, c.dict_type, c.sort
//		FROM gen_table t
//			 LEFT JOIN gen_table_column c ON t.table_id = c.table_id
//		where t.table_id = ? order by c.sort`,id)
//	return
//}
//func SelectGenTableByName(name string) (genTable *toolModels.GenTableVo)  {
//	genTable=new(toolModels.GenTableVo)
//	mysql.GetMasterDb().Get(&genTable,`SELECT t.table_id, t.table_name, t.table_comment, t.sub_table_name, t.sub_table_fk_name, t.class_name, t.tpl_category, t.package_name, t.module_name, t.business_name, t.function_name, t.function_author, t.gen_type, t.gen_path, t.options, t.remark,
//			   c.column_id, c.column_name, c.column_comment, c.column_type, c.java_type, c.java_field, c.is_pk, c.is_increment, c.is_required, c.is_insert, c.is_edit, c.is_list, c.is_query, c.query_type, c.html_type, c.dict_type, c.sort
//		FROM gen_table t
//			 LEFT JOIN gen_table_column c ON t.table_id = c.table_id
//		where t.table_name = ? order by c.sort`,name)
//	return
//}
//func SelectGenTableAll() (list []*toolModels.GenTableVo)  {
//	list=make([]*toolModels.GenTableVo,0,0)
//	mysql.GetMasterDb().Select(&list,`SELECT t.table_id, t.table_name, t.table_comment, t.sub_table_name, t.sub_table_fk_name, t.class_name, t.tpl_category, t.package_name, t.module_name, t.business_name, t.function_name, t.function_author, t.gen_type, t.gen_path, t.options, t.remark,
//			   c.column_id, c.column_name, c.column_comment, c.column_type, c.java_type, c.java_field, c.is_pk, c.is_increment, c.is_required, c.is_insert, c.is_edit, c.is_list, c.is_query, c.query_type, c.html_type, c.dict_type, c.sort
//		FROM gen_table t
//			 LEFT JOIN gen_table_column c ON t.table_id = c.table_id
//		order by c.sort`)
//	return
//}
//
//func InsertGenTable(genTable *toolModels.GenTableDML) (err error)  {
//	insertSQL := `insert into gen_table(table_id,table_name,table_comment,class_name,tpl_category,package_name,module_name,business_name,function_name,function_author,gen_type,gen_path,create_by,create_time,update_by,update_time %s)
//					values(:table_id,:table_name,:table_comment,:class_name,:tpl_category,:package_name,:module_name,:business_name,:function_name,:function_author,:gen_type,:gen_path,:create_by,now(),:update_by,now() %s)`
//	key := ""
//	value := ""
//
//	if genTable.Remark != "" {
//		key += ",remark"
//		value += ",:remark"
//	}
//
//	insertStr := fmt.Sprintf(insertSQL, key, value)
//	_, err = mysql.GetMasterDb().NamedExec(insertStr, genTable)
//	if err != nil {
//		zap.L().Error("数据库数据信息错误", zap.Error(err))
//	}
//	return
//}
//
//func UpdateGenTable(genTable *toolModels.GenTableDML) (err error)  {
//	updateSQL := `update gen_table set update_time = now() , update_by = :update_by`
//	if genTable.TableName != "" {
//		updateSQL += ",table_name = :table_name"
//	}
//	if genTable.TableComment != "" {
//		updateSQL += ",table_comment = :table_comment"
//	}
//	if genTable.SubTableName != "" {
//		updateSQL += ",sub_table_name = :sub_table_name"
//	}
//	if genTable.SubTableFkName != "" {
//		updateSQL += ",sub_table_fk_name = :sub_table_fk_name"
//	}
//	if genTable.ClassName != "" {
//		updateSQL += ",class_name = :class_name"
//	}
//	if genTable.FunctionAuthor != "" {
//		updateSQL += ",function_author = :function_author"
//	}
//	if genTable.GenType != "" {
//		updateSQL += ",gen_type = :gen_type"
//	}
//	if genTable.GenPath != "" {
//		updateSQL += ",gen_path = :gen_path"
//	}
//	if genTable.TplCategory != "" {
//		updateSQL += ",tpl_category = :tpl_category"
//	}
//	if genTable.PackageName != "" {
//		updateSQL += ",package_name = :package_name"
//	}
//	if genTable.ModuleName != "" {
//		updateSQL += ",module_name = :module_name"
//	}
//	if genTable.BusinessName != "" {
//		updateSQL += ",business_name = :business_name"
//	}
//	if genTable.FunctionName != "" {
//		updateSQL += ",function_name = :function_name"
//	}
//	if genTable.Options != "" {
//		updateSQL += ",options = :options"
//	}
//	if genTable.Remark != "" {
//		updateSQL += ",remark = :remark"
//	}
//
//	updateSQL += " where table_id = :table_id"
//
//	_, err = mysql.GetMasterDb().NamedExec(updateSQL, genTable)
//	if err != nil {
//		zap.L().Error("数据库数据信息错误", zap.Error(err))
//	}
//	return
//}
//
//func DeleteGenTableByIds(ids []int64) (err error)   {
//	query, i, err := sqlx.In(" delete from gen_table where table_id in(?)", ids)
//	_, err = mysql.GetMasterDb().Exec(query, i...)
//	return
//}
