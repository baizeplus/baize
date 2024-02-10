package daoImpl

import (
	"baize/app/baize"
	"baize/app/bzSystem/models"
	"context"
	"database/sql"
	"fmt"

	"github.com/baizeplus/sqly"
)

func NewSysRoleDao() *SysRoleDao {
	return &SysRoleDao{
		selectSql: ` select distinct r.role_id, r.role_name, r.role_key, r.role_sort, r.data_scope, r.menu_check_strictly, r.dept_check_strictly, r.status, r.del_flag, r.create_time, r.remark from sys_role r
	        left join sys_user_role ur on ur.role_id = r.role_id
	        left join sys_user u on u.user_id = ur.user_id
	        left join sys_dept d on u.dept_id = d.dept_id`,
	}
}

type SysRoleDao struct {
	selectSql string
}

func (rd *SysRoleDao) SelectRoleList(ctx context.Context, db sqly.SqlyContext, role *models.SysRoleDQL) (list []*models.SysRoleVo, total *int64) {
	whereSql := " where r.del_flag = '0'"
	if role.RoleName != "" {
		whereSql += " AND r.role_name like concat('%', :role_name, '%')"
	}
	if role.Status != "" {
		whereSql += " AND r.status = :status"
	}
	if role.RoleKey != "" {
		whereSql += " AND r.role_key like concat('%', :roleKey, '%')"
	}

	if role.BeginTime != "" {
		whereSql += " and date_format(r.create_time,'%y%m%d') &gt;= date_format(:begin_time,'%y%m%d')"
	}
	if role.EndTime != "" {
		whereSql += " and date_format(r.create_time,'%y%m%d') &lt;= date_format(:end_time,'%y%m%d')"
	}
	list = make([]*models.SysRoleVo, 0, 16)
	total = new(int64)
	err := db.NamedSelectPageContext(ctx, list, total, rd.selectSql+whereSql, role, role.ToPage())
	if err != nil {
		panic(err)
	}
	return
}
func (rd *SysRoleDao) SelectRoleAll(ctx context.Context, db sqly.SqlyContext) (list []*models.SysRoleVo) {
	whereSql := " where r.del_flag = '0'"
	list = make([]*models.SysRoleVo, 0, 16)
	err := db.NamedSelectContext(ctx, &list, rd.selectSql+whereSql, struct{}{})
	if err != nil {
		panic(err)
	}
	return
}
func (rd *SysRoleDao) SelectRoleById(ctx context.Context, db sqly.SqlyContext, roleId int64) (role *models.SysRoleVo) {
	whereSql := ` where r.role_id = ?`
	role = new(models.SysRoleVo)
	err := db.GetContext(ctx, role, rd.selectSql+whereSql, roleId)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}

func (rd *SysRoleDao) SelectBasicRolesByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (roles []*models.SysRole) {
	sqlStr := `select  r.role_id, r.role_name, r.role_key,r.data_scope
				from sys_role r
				left join sys_user_role ur  on r.role_id = ur.role_id
				where  ur.user_id = ?`
	roles = make([]*models.SysRole, 0, 2)
	err := db.SelectContext(ctx, &roles, sqlStr, userId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return
}

func (rd *SysRoleDao) SelectRolePermissionByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (roles []string) {
	sqlStr := `select   r.role_key
				from sys_role r
				left join sys_user_role ur  on r.role_id = ur.role_id
				where  ur.user_id = ?`
	roles = make([]string, 0, 1)
	err := db.SelectContext(ctx, &roles, sqlStr, userId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return
}
func (rd *SysRoleDao) SelectRoleIdAndDataScopeByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (roles []*baize.Role) {
	sqlStr := `select  r.role_id, r.data_scope
				from sys_role r
				left join sys_user_role ur  on r.role_id = ur.role_id
				where  ur.user_id = ?`
	roles = make([]*baize.Role, 0, 2)
	err := db.SelectContext(ctx, &roles, sqlStr, userId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return
}

func (rd *SysRoleDao) SelectRoleListByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (list []int64) {
	sqlStr := `select r.role_id
        from sys_role r
	        left join sys_user_role ur on ur.role_id = r.role_id
	        left join sys_user u on u.user_id = ur.user_id
		where u.user_id = ?`
	list = make([]int64, 0, 1)
	err := db.SelectContext(ctx, &list, sqlStr, userId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return
}

func (rd *SysRoleDao) InsertRole(ctx context.Context, db sqly.SqlyContext, sysRole *models.SysRoleDML) {
	insertSQL := `insert into sys_role(role_id,role_name,role_key,role_sort,create_by,create_time,update_by,update_time %s)
					values(:role_id,:role_name,:role_key,:role_sort,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""
	if sysRole.DataScope != "" {
		key += ",data_scope"
		value += ",:data_scope"
	}

	if sysRole.Status != "" {
		key += ",status"
		value += ",:status"
	}
	if sysRole.Remake != "" {
		key += ",remake"
		value += ",:remake"
	}
	insertStr := fmt.Sprintf(insertSQL, key, value)

	_, err := db.NamedExecContext(ctx, insertStr, sysRole)
	if err != nil {
		panic(err)
	}
	return
}

func (rd *SysRoleDao) UpdateRole(ctx context.Context, db sqly.SqlyContext, sysRole *models.SysRoleDML) {
	updateSQL := `update sys_role set update_time = now() , update_by = :update_by`

	if sysRole.RoleName != "" {
		updateSQL += ",role_name = :role_name"
	}
	if sysRole.RoleKey != "" {
		updateSQL += ",role_key = :role_key"
	}
	if sysRole.RoleSort != -1 {
		updateSQL += ",role_sort = :role_sort"
	}
	if sysRole.DataScope != "" {
		updateSQL += ",data_scope = :data_scope"
	}
	//if sysRole.PermissionCheckStrictly != nil {
	//	updateSQL += ",permission_check_strictly = :permission_check_strictly"
	//}
	//if sysRole.DeptCheckStrictly != nil {
	//	updateSQL += ",dept_check_strictly = :dept_check_strictly"
	//}
	if sysRole.Remake != "" {
		updateSQL += ",remake = :remake"
	}
	if sysRole.Status != "" {
		updateSQL += ",status = :status"
	}
	updateSQL += " where role_id = :role_id"
	_, err := db.NamedExecContext(ctx, updateSQL, sysRole)
	if err != nil {
		panic(err)
	}
	return
}

func (rd *SysRoleDao) DeleteRoleByIds(ctx context.Context, db sqly.SqlyContext, ids []int64) {
	query, i, err := sqly.In("update sys_role set del_flag = '2',role_name = concat(role_name,'(delete)')  where role_id in(?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
	return
}
func (rd *SysRoleDao) CheckRoleNameUnique(ctx context.Context, db sqly.SqlyContext, roleName string) int64 {
	var roleId int64 = 0
	err := db.GetContext(ctx, &roleId, "select role_id from sys_role where role_name = ?", roleName)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return roleId
}

func (rd *SysRoleDao) CheckRoleKeyUnique(ctx context.Context, db sqly.SqlyContext, roleKey string) int64 {
	var roleId int64 = 0
	err := db.GetContext(ctx, &roleId, "select role_id from sys_role where role_key = ?", roleKey)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return roleId
}
func (rd *SysRoleDao) SelectAllocatedList(ctx context.Context, db sqly.SqlyContext, user *models.SysRoleAndUserDQL) (list []*models.SysUserVo, total *int64) {
	selectStr := ` select distinct u.user_id, u.dept_id, u.user_name, u.nick_name, u.email, u.phonenumber, u.status, u.create_time`

	whereSql := ` from sys_user u
			 left join sys_dept d on u.dept_id = d.dept_id
			 left join sys_user_role ur on u.user_id = ur.user_id
			 left join sys_role r on r.role_id = ur.role_id where u.del_flag = '0' and r.role_id =:role_id`
	if user.UserName != "" {
		whereSql += " AND u.user_name like concat('%', :user_name, '%')"
	}
	if user.Phonenumber != "" {
		whereSql += " AND u.phonenumber like concat('%', :phonenumber, '%')"
	}
	list = make([]*models.SysUserVo, 0, 16)
	total = new(int64)
	err := db.NamedSelectPageContext(ctx, list, total, selectStr+whereSql, user, user.ToPage())
	if err != nil {
		panic(err)
	}
	return

}

func (rd *SysRoleDao) SelectUnallocatedList(ctx context.Context, db sqly.SqlyContext, user *models.SysRoleAndUserDQL) (list []*models.SysUserVo, total *int64) {
	selectStr := ` select distinct u.user_id, u.dept_id, u.user_name, u.nick_name, u.email, u.phonenumber, u.status, u.create_time`

	whereSql := `  from sys_user u
			 left join sys_dept d on u.dept_id = d.dept_id
			 left join sys_user_role ur on u.user_id = ur.user_id
			 left join sys_role r on r.role_id = ur.role_id
	    where u.del_flag = '0' and (r.role_id != :role_id or r.role_id IS NULL)
	    and u.user_id not in (select u.user_id from sys_user u inner join sys_user_role ur on u.user_id = ur.user_id and ur.role_id = :role_id)`
	if user.UserName != "" {
		whereSql += " AND u.user_name like concat('%', :user_name, '%')"
	}
	if user.Phonenumber != "" {
		whereSql += " AND u.phonenumber like concat('%', :phonenumber, '%')"
	}
	list = make([]*models.SysUserVo, 0, 16)
	total = new(int64)
	err := db.NamedSelectPageContext(ctx, list, total, selectStr+whereSql, user, user.ToPage())
	if err != nil {
		panic(err)
	}
	return

}
