package systemDao

import (
	"baize/app/bzSystem/models"
	"context"
	"github.com/baizeplus/sqly"
)

type IMenuDao interface {
	SelectMenuById(ctx context.Context, db sqly.SqlyContext, menuId int64) (menu *models.SysMenuVo)
	SelectMenuList(ctx context.Context, db sqly.SqlyContext, menu *models.SysMenuDQL) (list []*models.SysMenuVo)
	SelectMenuListByUserId(ctx context.Context, db sqly.SqlyContext, menu *models.SysMenuDQL) (list []*models.SysMenuVo)
	InsertMenu(ctx context.Context, db sqly.SqlyContext, menu *models.SysMenuVo)
	UpdateMenu(ctx context.Context, db sqly.SqlyContext, menu *models.SysMenuVo)
	DeleteMenuById(ctx context.Context, db sqly.SqlyContext, menuId int64)
	SelectMenuPermsByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (perms []string)
	SelectMenuTreeAll(ctx context.Context, db sqly.SqlyContext) (sysMenus []*models.SysMenuVo)
	SelectMenuTreeByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (sysMenus []*models.SysMenuVo)
	CheckMenuNameUnique(ctx context.Context, db sqly.SqlyContext, menuName string, parentId int64) int64
	HasChildByMenuId(ctx context.Context, db sqly.SqlyContext, menuId int64) int
	SelectMenuListByRoleId(ctx context.Context, db sqly.SqlyContext, roleId int64, menuCheckStrictly bool) (roleIds []string)
}
