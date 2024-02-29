package systemDaoImpl

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type SysUserDeptScopeDao struct {
}

func NewSysUserDeptScopeDao() *SysUserDeptScopeDao {
	return &SysUserDeptScopeDao{}
}

func (uds *SysUserDeptScopeDao) DeleteUserDeptScope(ctx context.Context, db sqly.SqlyContext, ids []int64) {
	query, i, err := sqly.In("delete from sys_user_dept_scope where user_id in (?)", ids)
	if err != nil {
		panic(err)
	}

	_, err = db.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
	return
}

func (uds *SysUserDeptScopeDao) DeleteUserDeptScopeByUserId(ctx context.Context, db sqly.SqlyContext, id int64) {

	_, err := db.ExecContext(ctx, "delete from sys_user_dept_scope where role_id=? ", id)
	if err != nil {
		panic(err)
	}

}
func (uds *SysUserDeptScopeDao) BatchUserDeptScope(ctx context.Context, db sqly.SqlyContext, list []*systemModels.SysUserDeptScope) {

	_, err := db.NamedExecContext(ctx, "insert into sys_user_dept_scope(user_id, dept_id) values (:user_id,:dept_id)", list)
	if err != nil {
		panic(err)
	}

}
