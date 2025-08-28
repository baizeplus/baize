package systemDaoImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"errors"

	"github.com/baizeplus/sqly"
)

type SysPermissionDao struct {
	ms                  sqly.SqlyContext
	selectPermissionSql string
	fromPermissionSql   string
}

func NewSysPermissionDao(ms sqly.SqlyContext) systemDao.IPermissionDao {
	return &SysPermissionDao{
		ms:                  ms,
		selectPermissionSql: `select distinctrow p.permission_id,p.permission_name,p.parent_id,p.permission,p.sort,p.status,p.create_by,p.create_time,p.update_by,p.update_time  `,
		fromPermissionSql:   ` from sys_permission p `,
	}
}

func (pd *SysPermissionDao) SelectPermissionById(ctx context.Context, permissionId string) *systemModels.SysPermissionVo {
	whereSql := ` where permission_id = ?`
	sp := new(systemModels.SysPermissionVo)
	err := pd.ms.GetContext(ctx, sp, pd.selectPermissionSql+pd.fromPermissionSql+whereSql, permissionId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	} else if err != nil {
		panic(err)
	}
	return sp
}
func (pd *SysPermissionDao) SelectPermissionList(ctx context.Context, permission *systemModels.SysPermissionDQL) (list []*systemModels.SysPermissionVo) {
	whereSql := ``
	if permission.Status != "" {
		whereSql += " AND p.status = :status"
	}
	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	whereSql += " order by p.sort"
	err := pd.ms.NamedSelectContext(ctx, &list, pd.selectPermissionSql+pd.fromPermissionSql+whereSql, permission)
	if err != nil {
		panic(err)
	}
	return

}
func (pd *SysPermissionDao) SelectPermissionListByParentId(ctx context.Context, parentId string) (list []*systemModels.SysPermissionVo) {
	list = make([]*systemModels.SysPermissionVo, 0)
	err := pd.ms.SelectContext(ctx, &list, pd.selectPermissionSql+pd.fromPermissionSql+"where parent_id = ? ", parentId)
	if err != nil {
		panic(err)
	}
	return list
}

func (pd *SysPermissionDao) SelectPermissionListByRoleIds(ctx context.Context, roleIds []string) (list []*systemModels.SysPermissionVo) {
	whereSql := `  left join sys_role_permission rp on rp.permission_id=p.permission_id
where rp.role_id in (?)	`
	list = make([]*systemModels.SysPermissionVo, 0)
	query, args, err := sqly.In(pd.selectPermissionSql+pd.fromPermissionSql+whereSql, roleIds)
	if err != nil {
		panic(err)
	}

	err = pd.ms.SelectContext(ctx, &list, query, args...)
	if err != nil {
		panic(err)
	}
	return
}

func (pd *SysPermissionDao) InsertPermission(ctx context.Context, permission *systemModels.SysPermissionAdd) {
	insertSQL := `insert into sys_permission(permission_id,permission_name,parent_id,permission,sort,status,create_by,create_time,update_by,update_time )
					values(:permission_id,:permission_name,:parent_id,:permission,:sort,:status,:create_by,:create_time,:update_by,:update_time)`
	_, err := pd.ms.NamedExecContext(ctx, insertSQL, permission)
	if err != nil {
		panic(err)
	}
	return
}

func (pd *SysPermissionDao) UpdatePermission(ctx context.Context, permission *systemModels.SysPermissionEdit) {
	updateSQL := `update sys_permission set update_time = :update_time, update_by = :update_by`

	if permission.PermissionName != "" {
		updateSQL += ",permission_name = :permission_name"

	}
	if permission.Sort != nil {
		updateSQL += ",sort = :sort"

	}
	if permission.ParentId != "" {
		updateSQL += ",parent_id = :parent_id"
	}

	if permission.Permission != "" {
		updateSQL += ",permission = :permission"
	}
	if permission.Status != "" {
		updateSQL += ",status = :status"
	}

	updateSQL += " where permission_id = :permission_id"

	_, err := pd.ms.NamedExecContext(ctx, updateSQL, permission)
	if err != nil {
		panic(err)
	}
	return
}

func (pd *SysPermissionDao) DeletePermissionById(ctx context.Context, permissionId string) {
	_, err := pd.ms.ExecContext(ctx, "delete from sys_permission where permission_id = ?", permissionId)
	if err != nil {
		panic(err)
	}
	return
}

func (pd *SysPermissionDao) HasChildByPermissionId(ctx context.Context, permissionId string) int {
	var count = 0
	err := pd.ms.GetContext(ctx, &count, "select exists(select * from sys_permission where parent_id = ?)", permissionId)
	if err != nil && !errors.Is(sql.ErrNoRows, err) {
		panic(err)
	}
	return count
}

func (pd *SysPermissionDao) SelectPermissionByUserId(ctx context.Context, userId string) []string {
	selectSql := `select distinct p.permission 
				from sys_permission p
				left join sys_role_permission rp on p.permission_id = rp.permission_id
				left join sys_role r on rp.role_id=r.role_id
				left join sys_user_role ur on r.role_id=ur.role_id
				where  r.status = '0' and ur.user_id =  ?`
	permission := make([]string, 0)
	err := pd.ms.SelectContext(ctx, &permission, selectSql, userId)
	if errors.Is(sql.ErrNoRows, err) {
		return permission
	} else if err != nil {
		panic(err)
	}
	return permission
}
func (pd *SysPermissionDao) SelectPermissionAll(ctx context.Context) []string {
	selectSql := `select distinct permission
				from sys_permission `
	permission := make([]string, 0)
	err := pd.ms.SelectContext(ctx, &permission, selectSql)
	if errors.Is(sql.ErrNoRows, err) {
		return permission
	} else if err != nil {
		panic(err)
	}
	return permission
}

func (pd *SysPermissionDao) SelectPermissionListSelectBoxByPerm(ctx context.Context, perm []string) (list []*systemModels.SelectPermission) {
	selectSql := `select permission_id ,permission_name ,parent_id 
				from sys_permission
				where status = '0' and permission in (?)  order by sort`
	query, args, err := sqly.In(selectSql, perm)
	if err != nil {
		panic(err)
	}
	err = pd.ms.SelectContext(ctx, &list, query, args...)
	if err != nil {
		panic(err)
	}
	return list
}
