package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IPermissionDao interface {
	SelectPermissionByUserId(ctx context.Context, userId string) []string
	SelectPermissionList(ctx context.Context, permission *systemModels.SysPermissionDQL) (list []*systemModels.SysPermissionVo)
	SelectPermissionById(ctx context.Context, permissionId string) *systemModels.SysPermissionVo
	SelectPermissionListByParentId(ctx context.Context, parentId string) (list []*systemModels.SysPermissionVo)
	SelectPermissionListByRoleIds(ctx context.Context, roleIds []string) (list []*systemModels.SysPermissionVo)
	InsertPermission(ctx context.Context, permission *systemModels.SysPermissionAdd)
	UpdatePermission(ctx context.Context, permission *systemModels.SysPermissionEdit)
	DeletePermissionById(ctx context.Context, permissionId string)
	HasChildByPermissionId(ctx context.Context, permissionId string) int
	SelectPermissionAll(ctx context.Context) []string
	SelectPermissionListSelectBoxByPerm(ctx context.Context, perm []string) (list []*systemModels.SelectPermission)
}
