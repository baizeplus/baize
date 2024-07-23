package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IRoleDao interface {
	SelectRoleList(ctx context.Context, db sqly.SqlyContext, role *systemModels.SysRoleDQL) (roleList []*systemModels.SysRoleVo, total int64)
	SelectRoleAll(ctx context.Context, db sqly.SqlyContext, role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo)
	SelectRoleById(ctx context.Context, db sqly.SqlyContext, roleId int64) (role *systemModels.SysRoleVo)
	SelectBasicRolesByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (roles []*systemModels.SysRole)
	SelectRolePermissionByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (roles []string)
	SelectRoleListByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (list []int64)
	InsertRole(ctx context.Context, db sqly.SqlyContext, sysRole *systemModels.SysRoleDML)
	UpdateRole(ctx context.Context, db sqly.SqlyContext, sysRole *systemModels.SysRoleDML)
	DeleteRoleByIds(ctx context.Context, db sqly.SqlyContext, ids []int64)
	CheckRoleNameUnique(ctx context.Context, db sqly.SqlyContext, roleName string) int64
	CheckRoleKeyUnique(ctx context.Context, db sqly.SqlyContext, roleKey string) int64
	SelectAllocatedList(ctx context.Context, db sqly.SqlyContext, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total int64)
	SelectUnallocatedList(ctx context.Context, db sqly.SqlyContext, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total int64)
}
