package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IRoleDao interface {
	SelectRoleList(ctx context.Context, role *systemModels.SysRoleDQL) (roleList []*systemModels.SysRoleVo, total int64)
	SelectRoleAll(ctx context.Context, role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo)
	SelectRoleById(ctx context.Context, roleId string) (role *systemModels.SysRoleVo)
	SelectBasicRolesByUserId(ctx context.Context, string string) (roles []*systemModels.SysRole)
	SelectRoleListByUserId(ctx context.Context, userId string) (list []string)
	InsertRole(ctx context.Context, sysRole *systemModels.SysRoleDML)
	UpdateRole(ctx context.Context, sysRole *systemModels.SysRoleDML)
	DeleteRoleByIds(ctx context.Context, ids []string)
	CheckRoleNameUnique(ctx context.Context, roleName string) string
	SelectAllocatedList(ctx context.Context, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total int64)
	SelectUnallocatedList(ctx context.Context, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total int64)
	SelectRoleIdAndNameAll(ctx context.Context) (list []*systemModels.SysRoleIdAndName)
	SelectRoleIdAndName(ctx context.Context, userId string, roleIds []string) (list []*systemModels.SysRoleIdAndName)
}
