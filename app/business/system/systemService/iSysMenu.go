package systemService

import (
	"baize/app/business/system/systemModels"
	"github.com/gin-gonic/gin"
)

type IMenuService interface {
	SelectMenuList(c *gin.Context, menu *systemModels.SysMenuDQL) (list []*systemModels.SysMenuVo)
	SelectMenuById(c *gin.Context, menuId int64) (menu *systemModels.SysMenuVo)
	InsertMenu(c *gin.Context, menu *systemModels.SysMenuVo)
	UpdateMenu(c *gin.Context, menu *systemModels.SysMenuVo)
	DeleteMenuById(c *gin.Context, menuId int64)
	SelectMenuTreeByUserId(c *gin.Context, userId int64) (sysMenu []*systemModels.SysMenuVo)
	BuildMenus(c *gin.Context, sysMenus []*systemModels.SysMenuVo) []*systemModels.RouterVo
	CheckMenuNameUnique(c *gin.Context, menu *systemModels.SysMenuVo) bool
	HasChildByMenuId(c *gin.Context, menuId int64) bool
	CheckMenuExistRole(c *gin.Context, menuId int64) bool
	SelectMenuListByRoleId(c *gin.Context, roleId int64) []string
}
