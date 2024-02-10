package daoImpl

import (
	"baize/app/bzSystem/models"
	"context"
	"database/sql"
	"fmt"

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

func (sysDictTypeDao *SysDictTypeDao) SelectDictTypeList(ctx context.Context, db sqly.SqlyContext, dictType *models.SysDictTypeDQL) (list []*models.SysDictTypeVo, total *int64) {
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
	list = make([]*models.SysDictTypeVo, 0, 16)
	total = new(int64)
	err := db.NamedSelectPageContext(ctx, list, total, sysDictTypeDao.dictTypeSql+whereSql, dictType, dictType.ToPage())
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *SysDictTypeDao) SelectDictTypeAll(ctx context.Context, db sqly.SqlyContext) (list []*models.SysDictTypeVo) {

	list = make([]*models.SysDictTypeVo, 0)
	err := db.SelectContext(ctx, &list, sysDictTypeDao.dictTypeSql)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *SysDictTypeDao) SelectDictTypeById(ctx context.Context, db sqly.SqlyContext, dictId int64) (dictType *models.SysDictTypeVo) {

	dictType = new(models.SysDictTypeVo)
	err := db.GetContext(ctx, dictType, sysDictTypeDao.dictTypeSql+" where dict_id = ?", dictId)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
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

func (sysDictTypeDao *SysDictTypeDao) InsertDictType(ctx context.Context, db sqly.SqlyContext, dictType *models.SysDictTypeVo) {
	insertSQL := `insert into sys_dict_type(dict_id,dict_name,dict_type,create_by,create_time,update_by,update_time %s)
					values(:dict_id,:dict_name,:dict_type,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""

	if dictType.Status != "" {
		key += ",status"
		value += ",:status"
	}

	if dictType.Remark != "" {
		key += ",remark"
		value += ",:remark"
	}

	insertStr := fmt.Sprintf(insertSQL, key, value)
	_, err := db.NamedExecContext(ctx, insertStr, dictType)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDictTypeDao *SysDictTypeDao) UpdateDictType(ctx context.Context, db sqly.SqlyContext, dictType *models.SysDictTypeVo) {
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

func (sysDictTypeDao *SysDictTypeDao) DeleteDictTypeByIds(ctx context.Context, db sqly.SqlyContext, dictIds []int64) (err error) {
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
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return dictId
}
