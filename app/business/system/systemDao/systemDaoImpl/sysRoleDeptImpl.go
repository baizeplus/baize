package systemDaoImpl

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type SysRoleDeptDao struct {
}

func NewSysRoleDeptDao() *SysRoleDeptDao {
	return &SysRoleDeptDao{}
}

func (sysRoleDeptDao *SysRoleDeptDao) DeleteRoleDept(ctx context.Context, db sqly.SqlyContext, ids []int64) {
	query, i, err := sqly.In("delete from sys_role_dept where role_id in (?)", ids)
	if err != nil {
		panic(err)
	}

	_, err = db.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
	return
}

func (sysRoleDeptDao *SysRoleDeptDao) DeleteRoleDeptByRoleId(ctx context.Context, db sqly.SqlyContext, id int64) {

	_, err := db.ExecContext(ctx, "delete from sys_role_dept where role_id=? ", id)
	if err != nil {
		panic(err)
	}

}
func (sysRoleDeptDao *SysRoleDeptDao) BatchRoleDept(ctx context.Context, db sqly.SqlyContext, list []*systemModels.SysRoleDept) {

	_, err := db.NamedExecContext(ctx, "insert into sys_role_dept(role_id, dept_id) values (:role_id,:dept_id)", list)
	if err != nil {
		panic(err)
	}

}
