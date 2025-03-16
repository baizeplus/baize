package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IPermissionDao interface {
	SelectPermissionByUserId(ctx context.Context, userId int64) []string
	SelectPermissionList(ctx context.Context, permission *systemModels.SysPermissionDQL) (list []*systemModels.SysPermissionVo)
	SelectPermissionById(ctx context.Context, permissionId int64) *systemModels.SysPermissionVo
	SelectPermissionListByParentId(ctx context.Context, parentId int64) (list []*systemModels.SysPermissionVo)
	SelectPermissionListByRoleIds(ctx context.Context, roleIds []int64) (list []*systemModels.SysPermissionVo)
	InsertPermission(ctx context.Context, permission *systemModels.SysPermissionAdd)
	UpdatePermission(ctx context.Context, permission *systemModels.SysPermissionEdit)
	DeletePermissionById(ctx context.Context, permissionId int64)
	HasChildByPermissionId(ctx context.Context, permissionId int64) int
	SelectPermissionAll(ctx context.Context) []string
	SelectPermissionListSelectBoxByPerm(ctx context.Context, perm []string) (list []*systemModels.SelectPermission)
}
