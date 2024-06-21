package systemDaoImpl

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type SysUserRoleDao struct {
}

func NewSysUserRoleDao() *SysUserRoleDao {
	return &SysUserRoleDao{}
}

func (sysUserRoleDao *SysUserRoleDao) DeleteUserRole(ctx context.Context, db sqly.SqlyContext, ids []int64) {
	query, i, err := sqly.In("delete from sys_user_role where user_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}

func (sysUserRoleDao *SysUserRoleDao) BatchUserRole(ctx context.Context, db sqly.SqlyContext, users []*systemModels.SysUserRole) {

	_, err := db.NamedExecContext(ctx, "insert into sys_user_role(user_id, role_id) values (:user_id,:role_id)", users)
	if err != nil {
		panic(err)
	}
}

func (sysUserRoleDao *SysUserRoleDao) DeleteUserRoleByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) {

	_, err := db.ExecContext(ctx, "delete from sys_user_role where user_id= ?", userId)
	if err != nil {
		panic(err)
	}
}
func (sysUserRoleDao *SysUserRoleDao) CountUserRoleByRoleId(ctx context.Context, db sqly.SqlyContext, ids []int64) int {
	var count = 0
	query, i, err := sqly.In("SELECT EXISTS( SELECT 1  from sys_user_role where role_id in(?))", ids)
	if err != nil {
		panic(err)
	}
	err = db.GetContext(ctx, &count, query, i...)
	if err != nil {
		panic(err)
	}
	return count
}
func (sysUserRoleDao *SysUserRoleDao) DeleteUserRoleInfo(ctx context.Context, db sqly.SqlyContext, userRole *systemModels.SysUserRole) {
	_, err := db.NamedExecContext(ctx, "delete from sys_user_role where user_id=:user_id and role_id=:role_id", userRole)
	if err != nil {
		panic(err)
	}
}
func (sysUserRoleDao *SysUserRoleDao) DeleteUserRoleInfos(ctx context.Context, db sqly.SqlyContext, roleId int64, userIds []int64) {
	query, i, err := sqly.In("delete from sys_user_role where role_id=(?) and user_id in (?)", roleId, userIds)
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}
