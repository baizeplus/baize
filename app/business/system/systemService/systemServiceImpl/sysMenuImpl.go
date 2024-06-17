package systemServiceImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemDao/systemDaoImpl"
	"baize/app/business/system/systemModels"
	"baize/app/constant/userConstants"
	"baize/app/utils/baizeContext"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"

	"baize/app/utils/snowflake"
	"baize/app/utils/stringUtils"
)

type MenuService struct {
	data        *sqly.DB
	menuDao     systemDao.IMenuDao
	roleMenuDao systemDao.IRoleMenuDao
	roleDao     systemDao.IRoleDao
}

func NewMenuService(data *sqly.DB, md *systemDaoImpl.SysMenuDao, rmd *systemDaoImpl.SysRoleMenuDao, rd *systemDaoImpl.SysRoleDao) *MenuService {
	return &MenuService{
		data:        data,
		menuDao:     md,
		roleMenuDao: rmd,
		roleDao:     rd,
	}
}

func (menuService *MenuService) SelectMenuList(c *gin.Context, menu *systemModels.SysMenuDQL) (list []*systemModels.SysMenuVo) {
	if baizeContext.IsAdmin(c) {
		list = menuService.menuDao.SelectMenuList(c, menuService.data, menu)
	} else {
		list = menuService.menuDao.SelectMenuListByUserId(c, menuService.data, menu)
	}
	return
}

func (menuService *MenuService) SelectMenuById(c *gin.Context, menuId int64) (menu *systemModels.SysMenuVo) {
	return menuService.menuDao.SelectMenuById(c, menuService.data, menuId)

}
func (menuService *MenuService) InsertMenu(c *gin.Context, menu *systemModels.SysMenuVo) {
	menu.MenuId = snowflake.GenID()
	menuService.menuDao.InsertMenu(c, menuService.data, menu)
}
func (menuService *MenuService) UpdateMenu(c *gin.Context, menu *systemModels.SysMenuVo) {
	menuService.menuDao.UpdateMenu(c, menuService.data, menu)
}
func (menuService *MenuService) DeleteMenuById(c *gin.Context, menuId int64) {
	menuService.menuDao.DeleteMenuById(c, menuService.data, menuId)
}

func (menuService *MenuService) SelectMenuTreeByUserId(c *gin.Context, userId int64) (sysMenu []*systemModels.SysMenuVo) {
	if baizeContext.IsAdmin(c) {
		sysMenu = menuService.menuDao.SelectMenuTreeAll(c, menuService.data)
	} else {
		sysMenu = menuService.menuDao.SelectMenuTreeByUserId(c, menuService.data, userId)
	}
	sysMenu = getChildPerms(sysMenu, 0)
	return
}

func (menuService *MenuService) BuildMenus(c *gin.Context, sysMenus []*systemModels.SysMenuVo) []*systemModels.RouterVo {
	routerVo := make([]*systemModels.RouterVo, 0, 2)
	for _, m := range sysMenus {
		r := new(systemModels.RouterVo)
		r.Hidden = m.Visible == "1"
		r.Name = m.GetRouteName()
		r.Path = m.GetRouterPath()
		r.Component = m.GetComponent()
		r.Meta.Title = m.MenuName
		r.Meta.Icon = m.Icon
		r.Meta.NoCache = m.IsCache == "1"
		cMenus := m.Children
		if cMenus != nil && len(cMenus) > 0 && m.MenuType == userConstants.TypeDir {
			r.AlwaysShow = true
			r.Redirect = "noRedirect"
			r.Children = menuService.BuildMenus(c, cMenus)
		} else if m.IsMenuFrame() {
			childrenList := make([]*systemModels.RouterVo, 0, 2)
			children := new(systemModels.RouterVo)
			children.Path = m.Path
			children.Component = m.Component
			children.Name = stringUtils.Capitalize(m.Path)
			children.Meta.Title = m.MenuName
			r.Meta.Icon = m.Icon
			r.Meta.NoCache = m.IsCache == "1"
			childrenList = append(childrenList, children)
			r.Children = childrenList
		}
		routerVo = append(routerVo, r)
	}
	return routerVo
}

func (menuService *MenuService) CheckMenuNameUnique(c *gin.Context, menu *systemModels.SysMenuVo) bool {
	RoleId := menuService.menuDao.CheckMenuNameUnique(c, menuService.data, menu.MenuName, menu.ParentId)
	if RoleId == menu.MenuId || RoleId == 0 {
		return false
	}
	return true
}

func (menuService *MenuService) HasChildByMenuId(c *gin.Context, menuId int64) bool {
	return menuService.menuDao.HasChildByMenuId(c, menuService.data, menuId) > 0
}

func (menuService *MenuService) CheckMenuExistRole(c *gin.Context, menuId int64) bool {
	return menuService.roleMenuDao.CheckMenuExistRole(c, menuService.data, menuId) > 0
}
func (menuService *MenuService) SelectMenuListByRoleId(c *gin.Context, roleId int64) []string {
	// TODO
	//role := menuService.roleDao.SelectRoleById(menuService.data,roleId)
	//return menuService.menuDao.SelectMenuListByRoleId(menuService.data,roleId, role.MenuCheckStrictly)
	return menuService.menuDao.SelectMenuListByRoleId(c, menuService.data, roleId, false)
}

func getChildPerms(menu []*systemModels.SysMenuVo, parentId int64) []*systemModels.SysMenuVo {
	sysMenus := make([]*systemModels.SysMenuVo, 0, 2)
	for _, sysMenu := range menu {
		if sysMenu.ParentId == parentId {
			recursionFn(menu, sysMenu)
			sysMenus = append(sysMenus, sysMenu)
		}
	}
	return sysMenus
}

/**
 * 递归列表
 *
 * @param menu
 * @param s
 */
func recursionFn(menu []*systemModels.SysMenuVo, s *systemModels.SysMenuVo) {
	childList := getChildList(menu, s)
	s.Children = childList
	for _, sysMenu := range childList {
		if hasChild(menu, sysMenu) {
			recursionFn(menu, sysMenu)
		}
	}
}

/**
 * 判断是否有子节点
 */
func hasChild(list []*systemModels.SysMenuVo, m *systemModels.SysMenuVo) bool {
	return len(getChildList(list, m)) > 0
}

/**
 * 得到子节点列表
 */
func getChildList(menu []*systemModels.SysMenuVo, s *systemModels.SysMenuVo) []*systemModels.SysMenuVo {
	tlist := make([]*systemModels.SysMenuVo, 0, 2)
	for _, sysMenu := range menu {
		if sysMenu.ParentId == s.MenuId {
			tlist = append(tlist, sysMenu)
		}

	}
	return tlist
}
