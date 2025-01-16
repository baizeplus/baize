package systemDaoImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"errors"
	"github.com/baizeplus/sqly"
)

type sysRoleMenuDao struct {
	ms sqly.SqlyContext
}

func NewSysRoleMenuDao(ms sqly.SqlyContext) systemDao.IRoleMenuDao {
	return &sysRoleMenuDao{ms: ms}
}

func (sysRoleMenuDao *sysRoleMenuDao) BatchRoleMenu(ctx context.Context, list []*systemModels.SysRoleMenu) {
	_, err := sysRoleMenuDao.ms.NamedExecContext(ctx, "insert into sys_role_menu(role_id, menu_id) values (:role_id,:menu_id)", list)
	if err != nil {
		panic(err)
	}
}

func (sysRoleMenuDao *sysRoleMenuDao) DeleteRoleMenu(ctx context.Context, ids []int64) {
	query, i, err := sqly.In("delete from sys_role_menu where role_id in (?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = sysRoleMenuDao.ms.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}

func (sysRoleMenuDao *sysRoleMenuDao) DeleteRoleMenuByRoleId(ctx context.Context, roleId int64) {
	_, err := sysRoleMenuDao.ms.ExecContext(ctx, "delete from sys_role_menu where role_id=?", roleId)
	if err != nil {
		panic(err)
	}
}

func (sysRoleMenuDao *sysRoleMenuDao) CheckMenuExistRole(ctx context.Context, menuId int64) int {
	var count = 0
	err := sysRoleMenuDao.ms.GetContext(ctx, &count, "select count(1) from sys_role_menu where menu_id = ?", menuId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return count
}
