package systemDao

import (
	"baize/app/baize"
	"baize/app/bzSystem/models"
	"context"
	"github.com/baizeplus/sqly"
)

type IRoleDao interface {
	SelectRoleList(ctx context.Context, db sqly.SqlyContext, role *models.SysRoleDQL) (roleList []*models.SysRoleVo, total *int64)
	SelectRoleAll(ctx context.Context, db sqly.SqlyContext) (list []*models.SysRoleVo)
	SelectRoleById(ctx context.Context, db sqly.SqlyContext, roleId int64) (role *models.SysRoleVo)
	SelectBasicRolesByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (roles []*models.SysRole)
	SelectRolePermissionByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (roles []string)
	SelectRoleIdAndDataScopeByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (roles []*baize.Role)
	SelectRoleListByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (list []int64)
	InsertRole(ctx context.Context, db sqly.SqlyContext, sysRole *models.SysRoleDML)
	UpdateRole(ctx context.Context, db sqly.SqlyContext, sysRole *models.SysRoleDML)
	DeleteRoleByIds(ctx context.Context, db sqly.SqlyContext, ids []int64)
	CheckRoleNameUnique(ctx context.Context, db sqly.SqlyContext, roleName string) int64
	CheckRoleKeyUnique(ctx context.Context, db sqly.SqlyContext, roleKey string) int64
	SelectAllocatedList(ctx context.Context, db sqly.SqlyContext, user *models.SysRoleAndUserDQL) (list []*models.SysUserVo, total *int64)
	SelectUnallocatedList(ctx context.Context, db sqly.SqlyContext, user *models.SysRoleAndUserDQL) (list []*models.SysUserVo, total *int64)
}
