package systemDaoImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"errors"
	"github.com/baizeplus/sqly"
)

type sysDictDataDao struct {
	ms          sqly.SqlyContext
	dictDataSql string
}

func NewSysDictDataDao(ms sqly.SqlyContext) systemDao.IDictDataDao {
	return &sysDictDataDao{
		ms:          ms,
		dictDataSql: `select dict_code, dict_sort, dict_label, dict_value, dict_type, css_class, list_class, is_default , status , create_by, create_time, remark  from sys_dict_data`,
	}
}

func (sysDictDataDao *sysDictDataDao) SelectDictDataByType(ctx context.Context, dictType string) (SysDictDataList []*systemModels.SysDictDataVo) {
	whereSql := ` where status = '0' and dict_type = ? order by dict_sort asc`

	SysDictDataList = make([]*systemModels.SysDictDataVo, 0, 0)

	err := sysDictDataDao.ms.SelectContext(ctx, &SysDictDataList, sysDictDataDao.dictDataSql+whereSql, dictType)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictDataDao *sysDictDataDao) SelectDictDataList(ctx context.Context, dictData *systemModels.SysDictDataDQL) (list []*systemModels.SysDictDataVo, total int64) {
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
	err := sysDictDataDao.ms.NamedSelectPageContext(ctx, &list, &total, sysDictDataDao.dictDataSql+whereSql, dictData)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictDataDao *sysDictDataDao) SelectDictDataById(ctx context.Context, dictCode int64) (dictData *systemModels.SysDictDataVo) {

	dictData = new(systemModels.SysDictDataVo)
	err := sysDictDataDao.ms.GetContext(ctx, dictData, sysDictDataDao.dictDataSql+" where dict_code = ?", dictCode)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return
}

func (sysDictDataDao *sysDictDataDao) InsertDictData(ctx context.Context, dictData *systemModels.SysDictDataVo) {
	insertSQL := `insert into sys_dict_data(dict_code,dict_sort,dict_label,dict_value,dict_type,css_class,list_class,is_default,status,remark,create_by,create_time,update_by,update_time )
					values(:dict_code,:dict_sort,:dict_label,:dict_value,:dict_type,:css_class,:list_class,:is_default,:status,:remark,:create_by,:create_time,:update_by,:update_time )`
	_, err := sysDictDataDao.ms.NamedExecContext(ctx, insertSQL, dictData)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictDataDao *sysDictDataDao) UpdateDictData(ctx context.Context, dictData *systemModels.SysDictDataVo) {
	updateSQL := `update sys_dict_data set update_time = :update_time , update_by = :update_by`

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

	_, err := sysDictDataDao.ms.NamedExecContext(ctx, updateSQL, dictData)
	if err != nil {
		panic(err)
	}
	return
}
func (sysDictDataDao *sysDictDataDao) SelectDictTypesByDictCodes(ctx context.Context, dictCodes []int64) []string {
	query, i, err := sqly.In("select dict_type from sys_dict_data where dict_code in(?)", dictCodes)
	if err != nil {
		panic(err)
	}
	list := make([]string, 0)
	err = sysDictDataDao.ms.SelectContext(ctx, &list, query, i...)
	if err != nil {
		panic(err)
	}
	return list
}

func (sysDictDataDao *sysDictDataDao) DeleteDictDataByIds(ctx context.Context, dictCodes []int64) {
	query, i, err := sqly.In("delete from sys_dict_data where dict_code in (?)", dictCodes)
	if err != nil {
		panic(err)
	}
	_, err = sysDictDataDao.ms.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func (sysDictDataDao *sysDictDataDao) CountDictDataByTypes(ctx context.Context, dictType []string) int {
	var count = 0
	query, i, err := sqly.In("SELECT EXISTS ( SELECT * FROM sys_dict_data where dict_type in(?))", dictType)
	if err != nil {
		panic(err)
	}
	err = sysDictDataDao.ms.GetContext(ctx, &count, query, i...)
	if err != nil {
		panic(err)
	}
	return count
}
