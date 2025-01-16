package systemDaoImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type sysUserDeptScopeDao struct {
	ms sqly.SqlyContext
}

func NewSysUserDeptScopeDao(ms sqly.SqlyContext) systemDao.IUserDeptScopeDao {
	return &sysUserDeptScopeDao{ms: ms}
}

func (uds *sysUserDeptScopeDao) DeleteUserDeptScope(ctx context.Context, ids []int64) {
	query, i, err := sqly.In("delete from sys_user_dept_scope where user_id in (?)", ids)
	if err != nil {
		panic(err)
	}

	_, err = uds.ms.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func (uds *sysUserDeptScopeDao) SelectUserDeptScopeDeptIdByUserId(ctx context.Context, id int64) []string {
	ids := make([]string, 0)
	err := uds.ms.SelectContext(ctx, &ids, "select dept_id from sys_user_dept_scope where user_id = ?", id)
	if err != nil {
		panic(err)
	}
	return ids
}

func (uds *sysUserDeptScopeDao) DeleteUserDeptScopeByUserId(ctx context.Context, id int64) {

	_, err := uds.ms.ExecContext(ctx, "delete from sys_user_dept_scope where user_id=? ", id)
	if err != nil {
		panic(err)
	}

}
func (uds *sysUserDeptScopeDao) BatchUserDeptScope(ctx context.Context, list []*systemModels.SysUserDeptScope) {
	_, err := uds.ms.NamedExecContext(ctx, "insert into sys_user_dept_scope(user_id, dept_id) values (:user_id,:dept_id)", list)
	if err != nil {
		panic(err)
	}

}
