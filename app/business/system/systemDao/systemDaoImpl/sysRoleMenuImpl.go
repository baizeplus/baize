package systemDaoImpl

import (
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"errors"
	"github.com/baizeplus/sqly"
)

type SysRoleMenuDao struct {
}

func NewSysRoleMenuDao() *SysRoleMenuDao {
	return &SysRoleMenuDao{}
}

func (sysRoleMenuDao *SysRoleMenuDao) BatchRoleMenu(ctx context.Context, db sqly.SqlyContext, list []*systemModels.SysRoleMenu) {
	_, err := db.NamedExecContext(ctx, "insert into sys_role_menu(role_id, menu_id) values (:role_id,:menu_id)", list)
	if err != nil {
		panic(err)
	}
}

func (sysRoleMenuDao *SysRoleMenuDao) DeleteRoleMenu(ctx context.Context, db sqly.SqlyContext, ids []int64) {
	query, i, err := sqly.In("delete from sys_role_menu where role_id in (?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}

func (sysRoleMenuDao *SysRoleMenuDao) DeleteRoleMenuByRoleId(ctx context.Context, db sqly.SqlyContext, roleId int64) {
	_, err := db.ExecContext(ctx, "delete from sys_role_menu where role_id=?", roleId)
	if err != nil {
		panic(err)
	}
}

func (sysRoleMenuDao *SysRoleMenuDao) CheckMenuExistRole(ctx context.Context, db sqly.SqlyContext, menuId int64) int {
	var count = 0
	err := db.GetContext(ctx, &count, "select count(1) from sys_role_menu where menu_id = ?", menuId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return count
}
