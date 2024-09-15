package systemDaoImpl

import (
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"errors"
	"github.com/baizeplus/sqly"
)

type SysDictTypeDao struct {
	dictTypeSql string
}

func NewSysDictTypeDao() *SysDictTypeDao {
	return &SysDictTypeDao{
		dictTypeSql: `select dict_id, dict_name, dict_type, status, create_by, create_time, remark   from sys_dict_type`,
	}
}

func (sysDictTypeDao *SysDictTypeDao) SelectDictTypeList(ctx context.Context, db sqly.SqlyContext, dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo, total int64) {
	whereSql := ``
	if dictType.DictName != "" {
		whereSql += " AND dict_name like concat('%', :dictName, '%')"
	}
	if dictType.Status != "" {
		whereSql += " AND  status = :status"
	}
	if dictType.DictType != "" {
		whereSql += " AND dict_type like concat('%', :dictType, '%')"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	err := db.NamedSelectPageContext(ctx, &list, &total, sysDictTypeDao.dictTypeSql+whereSql, dictType)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *SysDictTypeDao) SelectDictTypeAll(ctx context.Context, db sqly.SqlyContext, dictType *systemModels.SysDictTypeDQL) (list []*systemModels.SysDictTypeVo) {
	whereSql := ``
	if dictType.DictName != "" {
		whereSql += " AND dict_name like concat('%', :dictName, '%')"
	}
	if dictType.Status != "" {
		whereSql += " AND  status = :status"
	}
	if dictType.DictType != "" {
		whereSql += " AND dict_type like concat('%', :dictType, '%')"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	list = make([]*systemModels.SysDictTypeVo, 0)
	err := db.SelectContext(ctx, &list, sysDictTypeDao.dictTypeSql+whereSql)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *SysDictTypeDao) SelectDictTypeById(ctx context.Context, db sqly.SqlyContext, dictId int64) (dictType *systemModels.SysDictTypeVo) {

	dictType = new(systemModels.SysDictTypeVo)
	err := db.GetContext(ctx, dictType, sysDictTypeDao.dictTypeSql+" where dict_id = ?", dictId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return
}

func (sysDictTypeDao *SysDictTypeDao) SelectDictTypeByIds(ctx context.Context, db sqly.SqlyContext, dictId []int64) (dictTypes []string) {
	dictTypes = make([]string, 0)
	query, args, err := sqly.In("select dict_type from sys_dict_type where dict_id in(?)", dictId)
	if err != nil {
		panic(err)
	}

	err = db.SelectContext(ctx, &dictTypes, query, args...)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *SysDictTypeDao) InsertDictType(ctx context.Context, db sqly.SqlyContext, dictType *systemModels.SysDictTypeVo) {
	insertSQL := `insert into sys_dict_type(dict_id,dict_name,dict_type,status,remark,create_by,create_time,update_by,update_time )
					values(:dict_id,:dict_name,:dict_type,:status,:remark,:create_by,now(),:update_by,now() )`

	_, err := db.NamedExecContext(ctx, insertSQL, dictType)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *SysDictTypeDao) UpdateDictType(ctx context.Context, db sqly.SqlyContext, dictType *systemModels.SysDictTypeVo) {
	updateSQL := `update sys_dict_type set update_time = now() , update_by = :update_by`

	if dictType.DictName != "" {
		updateSQL += ",dict_name = :dict_name"
	}
	if dictType.DictType != "" {
		updateSQL += ",dict_type = :dict_type"
	}
	if dictType.Status != "" {
		updateSQL += ",status = :status"
	}
	if dictType.Remark != "" {
		updateSQL += ",remark = :remark"
	}

	updateSQL += " where dict_id = :dict_id"

	_, err := db.NamedExecContext(ctx, updateSQL, dictType)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *SysDictTypeDao) DeleteDictTypeByIds(ctx context.Context, db sqly.SqlyContext, dictIds []int64) {
	query, i, err := sqly.In("delete from sys_dict_type where dict_id in (?)", dictIds)
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func (sysDictTypeDao *SysDictTypeDao) CheckDictTypeUnique(ctx context.Context, db sqly.SqlyContext, dictType string) int64 {
	var dictId int64 = 0
	err := db.GetContext(ctx, &dictId, "select dict_id from sys_dict_type where dict_type = ?", dictType)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return dictId
}
