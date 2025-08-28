package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IRolePermissionDao interface {
	SelectPermissionIdsByRoleId(ctx context.Context, roleId string) []string
	BatchRolePermission(ctx context.Context, list []*systemModels.SysRolePermission)
	DeleteRolePermission(ctx context.Context, ids []string)
	DeleteRolePermissionByRoleId(ctx context.Context, roleId string)
	CheckPermissionExistRole(ctx context.Context, permissionId string) int
}
