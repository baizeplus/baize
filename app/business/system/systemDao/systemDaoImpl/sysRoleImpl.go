package systemDaoImpl

import (
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"errors"
	"github.com/baizeplus/sqly"
)

func NewSysRoleDao() *SysRoleDao {
	return &SysRoleDao{
		selectSql: ` select distinct r.role_id, r.role_name, r.role_key, r.role_sort,  r.status, r.del_flag, r.create_time, r.remark from sys_role r
	        left join sys_user_role ur on ur.role_id = r.role_id
	        left join sys_user u on u.user_id = ur.user_id
	        left join sys_dept d on u.dept_id = d.dept_id`,
	}
}

type SysRoleDao struct {
	selectSql string
}

func (rd *SysRoleDao) SelectRoleList(ctx context.Context, db sqly.SqlyContext, role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo, total int64) {
	whereSql := " where r.del_flag = '0'"
	if role.RoleName != "" {
		whereSql += " AND r.role_name like concat('%', :role_name, '%')"
	}
	if role.Status != "" {
		whereSql += " AND r.status = :status"
	}
	if role.RoleKey != "" {
		whereSql += " AND r.role_key like concat('%', :role_key, '%')"
	}
	if role.BeginTime != "" {
		whereSql += " and date_format(r.create_time,'%y%m%d') >= date_format(:begin_time,'%y%m%d')"
	}
	if role.EndTime != "" {
		whereSql += " and date_format(r.create_time,'%y%m%d') <= date_format(:end_time,'%y%m%d')"
	}
	err := db.NamedSelectPageContext(ctx, &list, &total, rd.selectSql+whereSql, role)
	if err != nil {
		panic(err)
	}
	return
}
func (rd *SysRoleDao) SelectRoleAll(ctx context.Context, db sqly.SqlyContext, role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo) {
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
		whereSql += " and date_format(r.create_time,'%y%m%d') >= date_format(:begin_time,'%y%m%d')"
	}
	if role.EndTime != "" {
		whereSql += " and date_format(r.create_time,'%y%m%d') <= date_format(:end_time,'%y%m%d')"
	}
	if role.CreateBy != 0 {
		whereSql += " and r.create_by = :create_by"
	}
	list = make([]*systemModels.SysRoleVo, 0)
	err := db.NamedSelectContext(ctx, &list, rd.selectSql+whereSql, role)
	if err != nil {
		panic(err)
	}
	return
}
func (rd *SysRoleDao) SelectRoleById(ctx context.Context, db sqly.SqlyContext, roleId int64) (role *systemModels.SysRoleVo) {
	whereSql := ` where r.role_id = ?`
	role = new(systemModels.SysRoleVo)
	err := db.GetContext(ctx, role, rd.selectSql+whereSql, roleId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return
}

func (rd *SysRoleDao) SelectBasicRolesByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (roles []*systemModels.SysRole) {
	sqlStr := `select  r.role_id, r.role_name, r.role_key
				from sys_role r
				left join sys_user_role ur  on r.role_id = ur.role_id
				where  ur.user_id = ?`
	roles = make([]*systemModels.SysRole, 0, 2)
	err := db.SelectContext(ctx, &roles, sqlStr, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
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
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
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
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return
}

func (rd *SysRoleDao) InsertRole(ctx context.Context, db sqly.SqlyContext, sysRole *systemModels.SysRoleDML) {
	insertSQL := `insert into sys_role(role_id,role_name,role_key,role_sort,status,remark,create_by,create_time,update_by,update_time )
					values(:role_id,:role_name,:role_key,:role_sort,:status,:remark,:create_by,now(),:update_by,now())`
	_, err := db.NamedExecContext(ctx, insertSQL, sysRole)
	if err != nil {
		panic(err)
	}
	return
}

func (rd *SysRoleDao) UpdateRole(ctx context.Context, db sqly.SqlyContext, sysRole *systemModels.SysRoleDML) {
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

	if sysRole.Remake != "" {
		updateSQL += ",remark = :remark"
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
	query, i, err := sqly.In("update sys_role set del_flag = '2',role_name = concat(role_name,'(删除)')  where role_id in(?)", ids)
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
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return roleId
}

func (rd *SysRoleDao) CheckRoleKeyUnique(ctx context.Context, db sqly.SqlyContext, roleKey string) int64 {
	var roleId int64 = 0
	err := db.GetContext(ctx, &roleId, "select role_id from sys_role where role_key = ?", roleKey)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return roleId
}
func (rd *SysRoleDao) SelectAllocatedList(ctx context.Context, db sqly.SqlyContext, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total int64) {
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
	err := db.NamedSelectPageContext(ctx, &list, &total, selectStr+whereSql, user)
	if err != nil {
		panic(err)
	}
	return

}

func (rd *SysRoleDao) SelectUnallocatedList(ctx context.Context, db sqly.SqlyContext, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total int64) {
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
	err := db.NamedSelectPageContext(ctx, &list, &total, selectStr+whereSql, user)
	if err != nil {
		panic(err)
	}
	return

}
