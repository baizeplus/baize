package systemDaoImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type sysUserRoleDao struct {
	ms sqly.SqlyContext
}

func NewSysUserRoleDao(ms sqly.SqlyContext) systemDao.IUserRoleDao {
	return &sysUserRoleDao{ms: ms}
}

func (sysUserRoleDao *sysUserRoleDao) DeleteUserRole(ctx context.Context, ids []int64) {
	query, i, err := sqly.In("delete from sys_user_role where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = sysUserRoleDao.ms.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}

func (sysUserRoleDao *sysUserRoleDao) BatchUserRole(ctx context.Context, users []*systemModels.SysUserRole) {

	_, err := sysUserRoleDao.ms.NamedExecContext(ctx, "insert into sys_user_role(user_id, role_id) values (:user_id,:role_id)", users)
	if err != nil {
		panic(err)
	}
}

func (sysUserRoleDao *sysUserRoleDao) DeleteUserRoleByUserId(ctx context.Context, userId int64) {

	_, err := sysUserRoleDao.ms.ExecContext(ctx, "delete from sys_user_role where user_id= ?", userId)
	if err != nil {
		panic(err)
	}
}
func (sysUserRoleDao *sysUserRoleDao) CountUserRoleByRoleId(ctx context.Context, ids []int64) int {
	var count = 0
	query, i, err := sqly.In("SELECT EXISTS( SELECT 1  from sys_user_role where role_id in(?))", ids)
	if err != nil {
		panic(err)
	}
	err = sysUserRoleDao.ms.GetContext(ctx, &count, query, i...)
	if err != nil {
		panic(err)
	}
	return count
}
func (sysUserRoleDao *sysUserRoleDao) DeleteUserRoleInfo(ctx context.Context, userRole *systemModels.SysUserRole) {
	_, err := sysUserRoleDao.ms.NamedExecContext(ctx, "delete from sys_user_role where user_id=:user_id and role_id=:role_id", userRole)
	if err != nil {
		panic(err)
	}
}
func (sysUserRoleDao *sysUserRoleDao) DeleteUserRoleInfos(ctx context.Context, roleId int64, userIds []int64) {
	query, i, err := sqly.In("delete from sys_user_role where role_id=(?) and user_id in (?)", roleId, userIds)
	if err != nil {
		panic(err)
	}
	_, err = sysUserRoleDao.ms.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}
