package systemDaoImpl

import (
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/baizeplus/sqly"
)

type SysDeptDao struct {
	deptSql string
}

func NewSysDeptDao() *SysDeptDao {
	return &SysDeptDao{
		deptSql: `select d.dept_id, d.parent_id, d.ancestors, d.dept_name, d.order_num, d.leader, d.phone, d.email, d.status, d.del_flag, d.create_by, d.create_time from sys_dept d`,
	}
}

func (sysDeptDao *SysDeptDao) SelectDeptList(ctx context.Context, db sqly.SqlyContext, dept *systemModels.SysDeptDQL) (list []*systemModels.SysDeptVo) {
	whereSql := ` where d.del_flag = '0'`
	if dept.ParentId != 0 {
		whereSql += " AND parent_id = :parent_id"
	}
	if dept.DeptName != "" {
		whereSql += " AND dept_name like concat('%', :dept_name, '%')"
	}
	if dept.Status != "" {
		whereSql += " AND status = :status"
	}
	if dept.DataScope != "" {
		whereSql += " AND " + dept.DataScope
	}
	whereSql += " order by d.parent_id, d.order_num"
	list = make([]*systemModels.SysDeptVo, 0, 16)
	err := db.NamedSelectContext(ctx, &list, sysDeptDao.deptSql+whereSql, dept)
	if err != nil {
		panic(err)
	}
	return list

}
func (sysDeptDao *SysDeptDao) SelectDeptById(ctx context.Context, db sqly.SqlyContext, deptId int64) (dept *systemModels.SysDeptVo) {
	whereSql := ` where d.dept_id = ?`
	dept = new(systemModels.SysDeptVo)
	err := db.GetContext(ctx, dept, sysDeptDao.deptSql+whereSql, deptId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return dept
}

func (sysDeptDao *SysDeptDao) InsertDept(ctx context.Context, db sqly.SqlyContext, dept *systemModels.SysDeptVo) {
	insertSQL := `insert into sys_dept(dept_id,parent_id,dept_name,order_num,create_by,create_time,update_by,update_time %s)
					values(:dept_id,:parent_id,:dept_name,:order_num,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""
	if dept.Ancestors != "" {
		key += ",ancestors"
		value += ",:ancestors"
	}
	if dept.Leader != "" {
		key += ",leader"
		value += ",:leader"
	}
	if dept.Phone != "" {
		key += ",phone"
		value += ",:phone"
	}
	if dept.Email != "" {
		key += ",email"
		value += ",:email"
	}
	if dept.Status != "" {
		key += ",status"
		value += ",:status"
	}

	insertStr := fmt.Sprintf(insertSQL, key, value)
	_, err := db.NamedExecContext(ctx, insertStr, dept)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDeptDao *SysDeptDao) UpdateDept(ctx context.Context, db sqly.SqlyContext, dept *systemModels.SysDeptVo) {
	updateSQL := `update sys_dept set order_num=:order_num , update_time = now() , update_by = :update_by `
	if dept.ParentId != 0 {
		updateSQL += ",parent_id = :parent_id"
	}

	if dept.DeptName != "" {
		updateSQL += ",dept_name = :dept_name"
	}
	if dept.Ancestors != "" {
		updateSQL += ",ancestors = :ancestors"
	}

	if dept.Leader != "" {
		updateSQL += ",leader = :leader"
	}
	if dept.Phone != "" {
		updateSQL += ",phone = :phone"
	}
	if dept.Email != "" {
		updateSQL += ",email = :email"
	}
	if dept.Status != "" {
		updateSQL += ",status = :status"
	}

	updateSQL += " where dept_id = :dept_id"
	_, err := db.NamedExecContext(ctx, updateSQL, dept)
	if err != nil {
		panic(err)
	}
	return
}

func (sysDeptDao *SysDeptDao) DeleteDeptById(ctx context.Context, db sqly.SqlyContext, deptId int64) {
	_, err := db.ExecContext(ctx, "update sys_dept set del_flag = '2',dept_name = concat(dept_name,'(删除)')  where dept_id =?", deptId)
	if err != nil {
		panic(err)
	}
	return
}
func (sysDeptDao *SysDeptDao) CheckDeptNameUnique(ctx context.Context, db sqly.SqlyContext, deptName string, parentId int64) int64 {
	var roleId int64 = 0
	err := db.GetContext(ctx, &roleId, "select dept_id from sys_dept where dept_name=? and parent_id = ?", deptName, parentId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return roleId
}
func (sysDeptDao *SysDeptDao) HasChildByDeptId(ctx context.Context, db sqly.SqlyContext, deptId int64) int {
	var count = 0
	err := db.GetContext(ctx, &count, "select count(1) from sys_dept where parent_id = ?", deptId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return count
}
func (sysDeptDao *SysDeptDao) CheckDeptExistUser(ctx context.Context, db sqly.SqlyContext, deptId int64) int {
	var count = 0
	err := db.GetContext(ctx, &count, "select count(1) from sys_user where dept_id = ? and del_flag = '0'", deptId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return count
}
