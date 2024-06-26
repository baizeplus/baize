package systemModels

import (
	"baize/app/baize"
	"baize/app/constant/userConstants"
	"baize/app/utils/stringUtils"
)

type SysMenuDQL struct {
	UserId   int64  `db:"userId" swaggerignore:"true"`
	MenuName string `form:"menuName" db:"menu_name"`
	Visible  string `form:"visible" db:"visible"`
	Status   string `form:"status" db:"status"`
	baize.BaseEntityDQL
}

type SysMenuVo struct {
	MenuId     int64  `json:"menuId,string" db:"menu_id"`
	MenuName   string `json:"menuName" db:"menu_name"`
	ParentName string `json:"parentName" db:"parent_name"`
	ParentId   int64  `json:"parentId,string" db:"parent_id"`
	OrderNum   int64  `json:"orderNum" db:"order_num"`
	Path       string `json:"path" db:"path"`
	Component  string `json:"component" db:"component"`
	IsFrame    string `json:"isFrame" db:"is_frame"`
	IsCache    string `json:"isCache" db:"is_cache"`
	MenuType   string `json:"menuType" db:"menu_type"`
	Visible    string `json:"visible" db:"visible"`
	Status     string `json:"status" db:"status"`
	Perms      string `json:"perms" db:"perms"`
	Icon       string `json:"icon" db:"icon"`
	Remark     string `json:"remark" db:"remark"`
	Children   []*SysMenuVo
	baize.BaseEntity
}

type MenusAndKeys struct {
	CheckedKeys []string     `json:"checkedKeys"`
	Menus       []*SysMenuVo `json:"menus"`
}

/**
 * 获取组件信息
 *
 * @return 组件信息
 */

func (m SysMenuVo) GetComponent() (component string) {

	if len(m.Component) != 0 && !m.IsMenuFrame() {
		component = m.Component
	} else if len(m.Component) == 0 && m.IsParentView() {
		component = userConstants.ParentView
	} else {
		component = userConstants.Layout
	}
	return
}

/**
 * 是否为parent_view组件
 *
 * @return 结果
 */

func (m SysMenuVo) IsParentView() bool {
	return m.ParentId != 0 && userConstants.TypeDir == m.MenuType
}

/**
* 获取路由地址
* @return 路由地址
 */
func (m SysMenuVo) GetRouterPath() (routerPath string) {

	if 0 == m.ParentId && userConstants.TypeDir == m.MenuType && m.IsFrame == userConstants.NoFrame {
		routerPath = "/" + m.Path
	} else if m.IsMenuFrame() {
		routerPath = "/"
	} else {
		routerPath = m.Path
	}
	return
}
func (m SysMenuVo) GetRouteName() (routerName string) {

	if m.IsMenuFrame() {
		routerName = ""
	} else {
		routerName = stringUtils.Capitalize(m.Path)
	}
	return
}

func (m SysMenuVo) IsMenuFrame() bool {
	return m.ParentId == 0 && m.MenuType == userConstants.TypeMenu && m.IsFrame == userConstants.NoFrame
}
