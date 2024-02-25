package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IMenuDao interface {
	SelectMenuById(ctx context.Context, db sqly.SqlyContext, menuId int64) (menu *systemModels.SysMenuVo)
	SelectMenuList(ctx context.Context, db sqly.SqlyContext, menu *systemModels.SysMenuDQL) (list []*systemModels.SysMenuVo)
	SelectMenuListByUserId(ctx context.Context, db sqly.SqlyContext, menu *systemModels.SysMenuDQL) (list []*systemModels.SysMenuVo)
	InsertMenu(ctx context.Context, db sqly.SqlyContext, menu *systemModels.SysMenuVo)
	UpdateMenu(ctx context.Context, db sqly.SqlyContext, menu *systemModels.SysMenuVo)
	DeleteMenuById(ctx context.Context, db sqly.SqlyContext, menuId int64)
	SelectMenuPermsByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (perms []string)
	SelectMenuPermsAll(ctx context.Context, db sqly.SqlyContext) (perms []string)
	SelectMenuTreeAll(ctx context.Context, db sqly.SqlyContext) (sysMenus []*systemModels.SysMenuVo)
	SelectMenuTreeByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (sysMenus []*systemModels.SysMenuVo)
	CheckMenuNameUnique(ctx context.Context, db sqly.SqlyContext, menuName string, parentId int64) int64
	HasChildByMenuId(ctx context.Context, db sqly.SqlyContext, menuId int64) int
	SelectMenuListByRoleId(ctx context.Context, db sqly.SqlyContext, roleId int64, menuCheckStrictly bool) (roleIds []string)
}
