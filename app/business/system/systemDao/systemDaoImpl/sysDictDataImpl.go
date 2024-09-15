package systemDaoImpl

import (
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"errors"
	"github.com/baizeplus/sqly"
)

type SysDictDataDao struct {
	dictDataSql string
}

func NewSysDictDataDao() *SysDictDataDao {
	return &SysDictDataDao{
		dictDataSql: `select dict_code, dict_sort, dict_label, dict_value, dict_type, css_class, list_class, is_default , status , create_by, create_time, remark  from sys_dict_data`,
	}
}

func (sysDictDataDao *SysDictDataDao) SelectDictDataByType(ctx context.Context, db sqly.SqlyContext, dictType string) (SysDictDataList []*systemModels.SysDictDataVo) {
	whereSql := ` where status = '0' and dict_type = ? order by dict_sort asc`

	SysDictDataList = make([]*systemModels.SysDictDataVo, 0, 0)

	err := db.SelectContext(ctx, &SysDictDataList, sysDictDataDao.dictDataSql+whereSql, dictType)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictDataDao *SysDictDataDao) SelectDictDataList(ctx context.Context, db sqly.SqlyContext, dictData *systemModels.SysDictDataDQL) (list []*systemModels.SysDictDataVo, total int64) {
	whereSql := ``
	if dictData.DictType != "" {
		whereSql += " AND dict_type = :dict_type"
	}
	if dictData.Status != "" {
		whereSql += " AND  status = :status"
	}
	if dictData.DictLabel != "" {
		whereSql += " AND dict_label like concat('%', :dict_label, '%')"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	err := db.NamedSelectPageContext(ctx, &list, &total, sysDictDataDao.dictDataSql+whereSql, dictData)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictDataDao *SysDictDataDao) SelectDictDataById(ctx context.Context, db sqly.SqlyContext, dictCode int64) (dictData *systemModels.SysDictDataVo) {

	dictData = new(systemModels.SysDictDataVo)
	err := db.GetContext(ctx, dictData, sysDictDataDao.dictDataSql+" where dict_code = ?", dictCode)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return
}

func (sysDictDataDao *SysDictDataDao) InsertDictData(ctx context.Context, db sqly.SqlyContext, dictData *systemModels.SysDictDataVo) {
	insertSQL := `insert into sys_dict_data(dict_code,dict_sort,dict_label,dict_value,dict_type,css_class,list_class,is_default,status,remark,create_by,create_time,update_by,update_time )
					values(:dict_code,:dict_sort,:dict_label,:dict_value,:dict_type,:css_class,:list_class,:is_default,:status,:remark,:create_by,now(),:update_by,now() )`
	_, err := db.NamedExecContext(ctx, insertSQL, dictData)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictDataDao *SysDictDataDao) UpdateDictData(ctx context.Context, db sqly.SqlyContext, dictData *systemModels.SysDictDataVo) {
	updateSQL := `update sys_dict_data set update_time = now() , update_by = :update_by`

	if dictData.DictSort != 0 {
		updateSQL += ",dict_sort = :dict_sort"
	}

	if dictData.DictLabel != "" {
		updateSQL += ",dict_label = :dict_label"
	}
	if dictData.DictValue != "" {
		updateSQL += ",dict_value = :dict_value"
	}
	if dictData.DictType != "" {
		updateSQL += ",dict_type = :dict_type"
	}
	if dictData.CssClass != "" {
		updateSQL += ",css_class = :css_class"
	}
	if dictData.ListClass != "" {
		updateSQL += ",list_class = :list_class"
	}
	if dictData.IsDefault != "" {
		updateSQL += ",is_default = :is_default"
	}
	if dictData.Status != "" {
		updateSQL += ",status = :status"
	}
	if dictData.Remark != "" {
		updateSQL += ",remark = :remark"
	}

	updateSQL += " where dict_code = :dict_code"

	_, err := db.NamedExecContext(ctx, updateSQL, dictData)
	if err != nil {
		panic(err)
	}
	return
}
func (sysDictDataDao *SysDictDataDao) SelectDictTypesByDictCodes(ctx context.Context, db sqly.SqlyContext, dictCodes []int64) []string {
	query, i, err := sqly.In("select dict_type from sys_dict_data where dict_code in(?)", dictCodes)
	if err != nil {
		panic(err)
	}
	list := make([]string, 0)
	err = db.SelectContext(ctx, &list, query, i...)
	if err != nil {
		panic(err)
	}
	return list
}

func (sysDictDataDao *SysDictDataDao) DeleteDictDataByIds(ctx context.Context, db sqly.SqlyContext, dictCodes []int64) {
	query, i, err := sqly.In("delete from sys_dict_data where dict_code in (?)", dictCodes)
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func (sysDictDataDao *SysDictDataDao) CountDictDataByTypes(ctx context.Context, db sqly.SqlyContext, dictType []string) int {
	var count = 0
	query, i, err := sqly.In("SELECT EXISTS ( SELECT * FROM sys_dict_data where dict_type in(?))", dictType)
	if err != nil {
		panic(err)
	}
	err = db.GetContext(ctx, &count, query, i...)
	if err != nil {
		panic(err)
	}
	return count
}
