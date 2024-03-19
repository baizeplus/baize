package systemController

import (
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/business/system/systemService/systemServiceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Menu struct {
	ms systemService.IMenuService
}

func NewMenu(ms *systemServiceImpl.MenuService) *Menu {
	return &Menu{ms: ms}
}

// MenuList 查询菜单列表查询
// @Summary 查询菜单列表查询
// @Description 查询菜单列表查询
// @Tags 菜单相关
// @Param  object query systemModels.SysMenuDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ListData{Rows=[]systemModels.SysMenuVo}}  "成功"
// @Router /system/menu/list  [get]
func (mc *Menu) MenuList(c *gin.Context) {
	menu := new(systemModels.SysMenuDQL)
	_ = c.ShouldBind(menu)
	menu.UserId = baizeContext.GetUserId(c)
	list := mc.ms.SelectMenuList(c, menu)
	baizeContext.SuccessData(c, list)
}

// MenuGetInfo 根据菜单ID获取菜单信息
// @Summary 根据菜单ID获取菜单信息
// @Description 根据菜单ID获取菜单信息
// @Tags 菜单相关
// @Param id path string true "menuId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=systemModels.SysMenuVo}  "成功"
// @Router /system/menu/{menuId}  [get]
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

// MenuTreeSelect 根据菜单ID获取菜单信息
// @Summary 根据菜单ID获取菜单信息
// @Description 根据菜单ID获取菜单信息
// @Tags 菜单相关
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=systemModels.SysMenuVo}  "成功"
// @Router /system/menu/treeSelect  [get]
func (mc *Menu) MenuTreeSelect(c *gin.Context) {
	userId := baizeContext.GetUserId(c)
	menu := new(systemModels.SysMenuDQL)
	menu.UserId = userId
	baizeContext.SuccessData(c, mc.ms.SelectMenuList(c, menu))
}

// MenuAdd 添加菜单
// @Summary 添加菜单
// @Description 添加菜单
// @Tags 菜单相关
// @Param  object body systemModels.SysMenuVo true "菜单信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/menu  [post]
func (mc *Menu) MenuAdd(c *gin.Context) {
	sysMenu := new(systemModels.SysMenuVo)
	_ = c.ShouldBind(sysMenu)
	if mc.ms.CheckMenuNameUnique(c, sysMenu) {
		baizeContext.Waring(c, "新增菜单'"+sysMenu.MenuName+"'失败，菜单名称已存在")
		return
	}
	sysMenu.SetCreateBy(baizeContext.GetUserId(c))
	mc.ms.InsertMenu(c, sysMenu)
	baizeContext.Success(c)
}

// MenuEdit 修改菜单
// @Summary 修改菜单
// @Description 修改菜单
// @Tags 菜单相关
// @Param  object body systemModels.SysMenuVo true "公司信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/menu  [put]
func (mc *Menu) MenuEdit(c *gin.Context) {
	sysMenu := new(systemModels.SysMenuVo)
	_ = c.ShouldBind(sysMenu)
	if mc.ms.CheckMenuNameUnique(c, sysMenu) {
		baizeContext.Waring(c, "修改菜单'"+sysMenu.MenuName+"'失败，菜单名称已存在")
		return
	}
	sysMenu.SetUpdateBy(baizeContext.GetUserId(c))
	mc.ms.UpdateMenu(c, sysMenu)
	baizeContext.Success(c)
}

// MenuRemove 删除菜单
// @Summary 删除菜单
// @Description 删除菜单
// @Tags 菜单相关
// @Param ids path string true "postId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/post/{menuId}  [delete]
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

// RoleMenuTreeSelect 删除菜单
// @Summary 删除菜单
// @Description 删除菜单
// @Tags 菜单相关
// @Param ids path string true "postId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=systemModels.MenusAndKeys}  "成功"
// @Router /system/post/roleMenuTreeSelect/{roleId}  [get]
func (mc *Menu) RoleMenuTreeSelect(c *gin.Context) {
	roleId := baizeContext.ParamInt64(c, "roleId")
	if roleId == 0 {
		zap.L().Debug("参数错误")
		baizeContext.ParameterError(c)
	}
	userId := baizeContext.GetUserId(c)
	mak := new(systemModels.MenusAndKeys)
	mak.CheckedKeys = mc.ms.SelectMenuListByRoleId(c, roleId)
	s := new(systemModels.SysMenuDQL)
	s.UserId = userId
	mak.Menus = mc.ms.SelectMenuList(c, s)
	baizeContext.SuccessData(c, mak)
}
