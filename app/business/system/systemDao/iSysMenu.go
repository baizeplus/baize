package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IMenuDao interface {
	SelectMenuById(ctx context.Context, menuId int64) (menu *systemModels.SysMenuVo)
	SelectMenuList(ctx context.Context, menu *systemModels.SysMenuDQL) (list []*systemModels.SysMenuVo)
	SelectMenuListByUserId(ctx context.Context, menu *systemModels.SysMenuDQL) (list []*systemModels.SysMenuVo)
	InsertMenu(ctx context.Context, menu *systemModels.SysMenuVo)
	UpdateMenu(ctx context.Context, menu *systemModels.SysMenuVo)
	DeleteMenuById(ctx context.Context, menuId int64)
	SelectMenuPermsByUserId(ctx context.Context, userId int64) (perms []string)
	SelectMenuPermsAll(ctx context.Context) (perms []string)
	SelectMenuTreeAll(ctx context.Context) (sysMenus []*systemModels.SysMenuVo)
	SelectMenuTreeByUserId(ctx context.Context, userId int64) (sysMenus []*systemModels.SysMenuVo)
	CheckMenuNameUnique(ctx context.Context, menuName string, parentId int64) int64
	HasChildByMenuId(ctx context.Context, menuId int64) int
	SelectMenuListByRoleId(ctx context.Context, roleId int64, menuCheckStrictly bool) (roleIds []string)
}
