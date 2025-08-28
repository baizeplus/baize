package systemDaoImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"errors"

	"github.com/baizeplus/sqly"
)

type sysRolePermissionDao struct {
	ms sqly.SqlyContext
}

func NewSysRolePermissionDao(ms sqly.SqlyContext) systemDao.IRolePermissionDao {
	return &sysRolePermissionDao{ms: ms}
}

func (sysRolePermissionDao *sysRolePermissionDao) SelectPermissionIdsByRoleId(ctx context.Context, roleId string) []string {
	ids := make([]string, 0)
	err := sysRolePermissionDao.ms.SelectContext(ctx, &ids, "select permission_id from sys_role_permission where role_id = ?", roleId)
	if err != nil {
		panic(err)
	}
	return ids
}

func (sysRolePermissionDao *sysRolePermissionDao) BatchRolePermission(ctx context.Context, list []*systemModels.SysRolePermission) {
	_, err := sysRolePermissionDao.ms.NamedExecContext(ctx, "insert into sys_role_permission(role_id, permission_id) values (:role_id,:permission_id)", list)
	if err != nil {
		panic(err)
	}
}

func (sysRolePermissionDao *sysRolePermissionDao) DeleteRolePermission(ctx context.Context, ids []string) {
	query, i, err := sqly.In("delete from sys_role_permission where role_id in (?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = sysRolePermissionDao.ms.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}

func (sysRolePermissionDao *sysRolePermissionDao) DeleteRolePermissionByRoleId(ctx context.Context, roleId string) {
	_, err := sysRolePermissionDao.ms.ExecContext(ctx, "delete from sys_role_permission where role_id=?", roleId)
	if err != nil {
		panic(err)
	}
}

func (sysRolePermissionDao *sysRolePermissionDao) CheckPermissionExistRole(ctx context.Context, permissionId string) int {
	var count = 0
	err := sysRolePermissionDao.ms.GetContext(ctx, &count, "select count(1) from sys_role_permission where permission_id = ?", permissionId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return count
}
