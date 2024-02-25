package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IRoleMenuDao interface {
	BatchRoleMenu(ctx context.Context, db sqly.SqlyContext, list []*systemModels.SysRoleMenu)
	DeleteRoleMenu(ctx context.Context, db sqly.SqlyContext, ids []int64)
	DeleteRoleMenuByRoleId(ctx context.Context, db sqly.SqlyContext, roleId int64)
	CheckMenuExistRole(ctx context.Context, db sqly.SqlyContext, menuId int64) int
}
