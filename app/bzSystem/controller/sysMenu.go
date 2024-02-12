package controller

import (
	"baize/app/bzSystem/models"
	"baize/app/bzSystem/service"
	"baize/app/bzSystem/service/serviceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Menu struct {
	ms service.IMenuService
}

func NewMenu(ms *serviceImpl.MenuService) *Menu {
	return &Menu{ms: ms}
}

func (mc *Menu) MenuList(c *gin.Context) {
	menu := new(models.SysMenuDQL)
	_ = c.ShouldBind(menu)
	list := mc.ms.SelectMenuList(c, menu, baizeContext.GetUserId(c))
	baizeContext.SuccessData(c, list)
}
func (mc *Menu) MenuGetInfo(c *gin.Context) {
	menuId := baizeContext.ParamInt64(c, "menuId")
	if menuId == 0 {
		zap.L().Debug("参数错误")
		baizeContext.ParameterError(c)
		return
	}
	menu := mc.ms.SelectMenuById(c, menuId)
	baizeContext.SuccessData(c, menu)
}
func (mc *Menu) MenuTreeSelect(c *gin.Context) {
	userId := baizeContext.GetUserId(c)
	baizeContext.SuccessData(c, mc.ms.SelectMenuList(c, new(models.SysMenuDQL), userId))
}
func (mc *Menu) MenuAdd(c *gin.Context) {
	sysMenu := new(models.SysMenuVo)
	_ = c.ShouldBind(sysMenu)
	if mc.ms.CheckMenuNameUnique(c, sysMenu) {
		baizeContext.Waring(c, "新增菜单'"+sysMenu.MenuName+"'失败，菜单名称已存在")
		return
	}
	sysMenu.SetCreateBy(baizeContext.GetUserId(c))
	mc.ms.InsertMenu(c, sysMenu)
	baizeContext.Success(c)
}
func (mc *Menu) MenuEdit(c *gin.Context) {
	sysMenu := new(models.SysMenuVo)
	_ = c.ShouldBind(sysMenu)
	if mc.ms.CheckMenuNameUnique(c, sysMenu) {
		baizeContext.Waring(c, "修改菜单'"+sysMenu.MenuName+"'失败，菜单名称已存在")
		return
	}

	sysMenu.SetUpdateBy(baizeContext.GetUserId(c))
	mc.ms.UpdateMenu(c, sysMenu)
	baizeContext.Success(c)
}
func (mc *Menu) MenuRemove(c *gin.Context) {
	menuId := baizeContext.ParamInt64(c, "menuId")
	if menuId == 0 {
		zap.L().Debug("参数错误")
		baizeContext.ParameterError(c)
		return
	}
	if mc.ms.HasChildByMenuId(c, menuId) {
		baizeContext.Waring(c, "存在子菜单,不允许删除")
		return
	}
	if mc.ms.CheckMenuExistRole(c, menuId) {
		baizeContext.Waring(c, "菜单已分配,不允许删除")
		return
	}
	mc.ms.DeleteMenuById(c, menuId)
	baizeContext.Success(c)
}
func (mc *Menu) RoleMenuTreeSelect(c *gin.Context) {
	roleId := baizeContext.ParamInt64(c, "roleId")
	if roleId == 0 {
		zap.L().Debug("参数错误")
		baizeContext.ParameterError(c)
	}
	userId := baizeContext.GetUserId(c)
	m := make(map[string]interface{})
	m["checkedKeys"] = mc.ms.SelectMenuListByRoleId(c, roleId)
	m["menus"] = mc.ms.SelectMenuList(c, new(models.SysMenuDQL), userId)
	baizeContext.SuccessData(c, m)
}
