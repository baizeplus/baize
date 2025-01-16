package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IRoleMenuDao interface {
	BatchRoleMenu(ctx context.Context, list []*systemModels.SysRoleMenu)
	DeleteRoleMenu(ctx context.Context, ids []int64)
	DeleteRoleMenuByRoleId(ctx context.Context, roleId int64)
	CheckMenuExistRole(ctx context.Context, menuId int64) int
}
