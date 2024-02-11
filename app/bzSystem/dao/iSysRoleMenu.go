package dao

import (
	"baize/app/bzSystem/models"
	"context"
	"github.com/baizeplus/sqly"
)

type IRoleMenuDao interface {
	BatchRoleMenu(ctx context.Context, db sqly.SqlyContext, list []*models.SysRoleMenu)
	DeleteRoleMenu(ctx context.Context, db sqly.SqlyContext, ids []int64)
	DeleteRoleMenuByRoleId(ctx context.Context, db sqly.SqlyContext, roleId int64)
	CheckMenuExistRole(ctx context.Context, db sqly.SqlyContext, menuId int64) int
}
