package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IRoleDao interface {
	SelectRoleList(ctx context.Context, role *systemModels.SysRoleDQL) (roleList []*systemModels.SysRoleVo, total int64)
	SelectRoleAll(ctx context.Context, role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo)
	SelectRoleById(ctx context.Context, roleId int64) (role *systemModels.SysRoleVo)
	SelectBasicRolesByUserId(ctx context.Context, userId int64) (roles []*systemModels.SysRole)
	SelectRolePermissionByUserId(ctx context.Context, userId int64) (roles []string)
	SelectRoleListByUserId(ctx context.Context, userId int64) (list []int64)
	InsertRole(ctx context.Context, sysRole *systemModels.SysRoleDML)
	UpdateRole(ctx context.Context, sysRole *systemModels.SysRoleDML)
	DeleteRoleByIds(ctx context.Context, ids []int64)
	CheckRoleNameUnique(ctx context.Context, roleName string) int64
	CheckRoleKeyUnique(ctx context.Context, roleKey string) int64
	SelectAllocatedList(ctx context.Context, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total int64)
	SelectUnallocatedList(ctx context.Context, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total int64)
}
