package service

import (
	"baize/app/bzSystem/models"
	"github.com/gin-gonic/gin"
)

type IMenuService interface {
	SelectMenuList(c *gin.Context, menu *models.SysMenuDQL, userId int64) (list []*models.SysMenuVo)
	SelectMenuById(c *gin.Context, menuId int64) (menu *models.SysMenuVo)
	InsertMenu(c *gin.Context, menu *models.SysMenuVo)
	UpdateMenu(c *gin.Context, menu *models.SysMenuVo)
	DeleteMenuById(c *gin.Context, menuId int64)
	SelectMenuTreeByUserId(c *gin.Context, userId int64) (sysMenu []*models.SysMenuVo)
	BuildMenus(c *gin.Context, sysMenus []*models.SysMenuVo) []*models.RouterVo
	CheckMenuNameUnique(c *gin.Context, menu *models.SysMenuVo) bool
	HasChildByMenuId(c *gin.Context, menuId int64) bool
	CheckMenuExistRole(c *gin.Context, menuId int64) bool
	SelectMenuListByRoleId(c *gin.Context, roleId int64) []string
}
