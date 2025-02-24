package systemDaoImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"errors"
	"github.com/baizeplus/sqly"
)

type sysMenuDao struct {
	ms            sqly.SqlyContext
	selectMenuSql string
}

func NewSysMenuDao(ms sqly.SqlyContext) systemDao.IMenuDao {
	return &sysMenuDao{
		ms:            ms,
		selectMenuSql: `select distinct m.menu_id, m.parent_id, m.menu_name, m.path, m.component, m.visible, m.status, ifnull(m.perms,'') as perms, m.is_frame, m.is_cache, m.menu_type, m.icon, m.order_num, m.create_time from sys_menu m`,
	}
}

func (sysMenuDao *sysMenuDao) SelectMenuById(ctx context.Context, menuId int64) (menu *systemModels.SysMenuVo) {
	whereSql := ` where menu_id = ?`
	menu = new(systemModels.SysMenuVo)
	err := sysMenuDao.ms.GetContext(ctx, menu, sysMenuDao.selectMenuSql+whereSql, menuId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return
}

func (sysMenuDao *sysMenuDao) SelectMenuList(ctx context.Context, menu *systemModels.SysMenuDQL) (list []*systemModels.SysMenuVo) {
	whereSql := ``
	if menu.MenuName != "" {
		whereSql += " AND menu_name like concat('%', :menu_name, '%')"
	}
	if menu.Visible != "" {
		whereSql += " AND Visible = :visible"
	}
	if menu.MenuName != "" {
		whereSql += "  AND status = :status"
	}
	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	whereSql += " order by m.parent_id, m.order_num"

	list = make([]*systemModels.SysMenuVo, 0, 16)
	err := sysMenuDao.ms.NamedSelectContext(ctx, &list, sysMenuDao.selectMenuSql+whereSql, menu)

	if err != nil {
		panic(err)
	}

	return list
}

func (sysMenuDao *sysMenuDao) SelectMenuListByUserId(ctx context.Context, menu *systemModels.SysMenuDQL) (list []*systemModels.SysMenuVo) {
	whereSql := ` left join sys_role_menu rm on m.menu_id = rm.menu_id
		left join sys_user_role ur on rm.role_id = ur.role_id
		left join sys_role ro on ur.role_id = ro.role_id
		where ur.user_id = :user_id`
	if menu.MenuName != "" {
		whereSql += " AND menu_name like concat('%', :menu_name, '%')"
	}
	if menu.Visible != "" {
		whereSql += " AND Visible = :visible"
	}
	if menu.MenuName != "" {
		whereSql += "  AND status = :status"
	}
	whereSql += " m.parent_id, m.order_num"

	list = make([]*systemModels.SysMenuVo, 0, 16)
	err := sysMenuDao.ms.NamedSelectContext(ctx, &list, sysMenuDao.selectMenuSql+whereSql, menu)

	if err != nil {
		panic(err)
	}

	return list
}

func (sysMenuDao *sysMenuDao) InsertMenu(ctx context.Context, menu *systemModels.SysMenuVo) {
	insertSQL := `insert into sys_menu(menu_id,menu_name,parent_id,create_by,create_time,update_by,update_time,order_num,path,component,is_frame,is_cache,menu_type,icon,status,perms,visible,remark)
					values(:menu_id,:menu_name,:parent_id,:create_by,:create_time,:update_by,:update_time ,:order_num,:path,:component,:is_frame,:is_cache,:menu_type,:icon,:status,:perms,:visible,:remark)`
	_, err := sysMenuDao.ms.NamedExecContext(ctx, insertSQL, menu)
	if err != nil {
		panic(err)
	}
	return
}

func (sysMenuDao *sysMenuDao) UpdateMenu(ctx context.Context, menu *systemModels.SysMenuVo) {
	updateSQL := `update sys_menu set update_time = :update_time , update_by = :update_by`

	if menu.ParentId != 0 {
		updateSQL += ",parent_id = :parent_id"
	}
	if menu.MenuName != "" {
		updateSQL += ",menu_name = :menu_name"
	}
	if menu.OrderNum != 0 {
		updateSQL += ",order_num = :order_num"
	}
	if menu.Path != "" {
		updateSQL += ",path = :path"
	}
	if menu.Component != "" {
		updateSQL += ",component = :component"
	}
	if menu.IsFrame != "" {
		updateSQL += ",is_frame = :is_frame"
	}
	if menu.IsCache != "" {
		updateSQL += ",is_cache = :is_cache"
	}
	if menu.MenuType != "" {
		updateSQL += ",menu_type = :menu_type"
	}
	if menu.Visible != "" {
		updateSQL += ",visible = :visible"
	}
	if menu.Status != "" {
		updateSQL += ",status = :status"
	}
	if menu.Perms != "" {
		updateSQL += ",perms = :perms"
	}
	if menu.Icon != "" {
		updateSQL += ",icon = :icon"
	}
	if menu.Status != "" {
		updateSQL += ",status = :status"
	}
	updateSQL += " where menu_id = :menu_id"

	_, err := sysMenuDao.ms.NamedExecContext(ctx, updateSQL, menu)
	if err != nil {
		panic(err)
	}
	return
}

func (sysMenuDao *sysMenuDao) DeleteMenuById(ctx context.Context, menuId int64) {
	_, err := sysMenuDao.ms.ExecContext(ctx, "delete from sys_menu where menu_id = ?", menuId)
	if err != nil {
		panic(err)
	}
	return
}

func (sysMenuDao *sysMenuDao) SelectMenuPermsByUserId(ctx context.Context, userId int64) (perms []string) {
	sqlStr := `	select distinct m.perms
				from sys_menu m
				left join sys_role_menu rm on m.menu_id = rm.menu_id
				left join sys_user_role ur on rm.role_id = ur.role_id
				left join sys_role r on r.role_id = ur.role_id
				where m.status = '0' and r.status = '0' and ur.user_id =  ?`
	perms = make([]string, 0, 2)
	err := sysMenuDao.ms.SelectContext(ctx, &perms, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
}
func (sysMenuDao *sysMenuDao) SelectMenuPermsAll(ctx context.Context) (perms []string) {
	sqlStr := `	select distinct m.perms
				from sys_menu m`
	perms = make([]string, 0, 2)
	err := sysMenuDao.ms.SelectContext(ctx, &perms, sqlStr)
	if err != nil {
		panic(err)
	}
	return
}

func (sysMenuDao *sysMenuDao) SelectMenuTreeAll(ctx context.Context) (sysMenus []*systemModels.SysMenuVo) {
	whereSql := ` where m.menu_type in ('M', 'C') and m.status = 0
		order by m.parent_id, m.order_num`
	sysMenus = make([]*systemModels.SysMenuVo, 0, 2)
	err := sysMenuDao.ms.SelectContext(ctx, &sysMenus, sysMenuDao.selectMenuSql+whereSql)
	if err != nil {
		panic(err)
	}
	return
}
func (sysMenuDao *sysMenuDao) SelectMenuTreeByUserId(ctx context.Context, userId int64) (sysMenus []*systemModels.SysMenuVo) {
	whereSql := ` left join sys_role_menu rm on m.menu_id = rm.menu_id
			 left join sys_user_role ur on rm.role_id = ur.role_id
			 left join sys_role ro on ur.role_id = ro.role_id
			 left join sys_user u on ur.user_id = u.user_id
		where u.user_id = ? and m.menu_type in ('M', 'C') and m.status = 0  AND ro.status = 0
		order by m.parent_id, m.order_num`
	sysMenus = make([]*systemModels.SysMenuVo, 0, 2)
	err := sysMenuDao.ms.SelectContext(ctx, &sysMenus, sysMenuDao.selectMenuSql+whereSql, userId)
	if err != nil {
		panic(err)
	}
	return
}

func (sysMenuDao *sysMenuDao) CheckMenuNameUnique(ctx context.Context, menuName string, parentId int64) int64 {
	var roleId int64 = 0
	err := sysMenuDao.ms.GetContext(ctx, &roleId, "select menu_id from sys_menu where menu_name=? and parent_id = ?", menuName, parentId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return roleId
}

func (sysMenuDao *sysMenuDao) HasChildByMenuId(ctx context.Context, menuId int64) int {
	var count = 0
	err := sysMenuDao.ms.GetContext(ctx, &count, "SELECT EXISTS( SELECT 1 FROM sys_menu where parent_id = ?)", menuId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return count
}
func (sysMenuDao *sysMenuDao) SelectMenuListByRoleId(ctx context.Context, roleId int64, menuCheckStrictly bool) (roleIds []string) {
	var err error
	roleIds = make([]string, 0, 2)
	sqlstr := `select m.menu_id
		from sys_menu m
            left join sys_role_menu rm on m.menu_id = rm.menu_id
        where rm.role_id = ?`
	if menuCheckStrictly {
		sqlstr += " and m.menu_id not in (select m.parent_id from sys_menu m inner join sys_role_menu rm on m.menu_id = rm.menu_id and rm.role_id = ?)"
		err = sysMenuDao.ms.SelectContext(ctx, &roleIds, sqlstr, roleId, roleId)
	} else {
		err = sysMenuDao.ms.SelectContext(ctx, &roleIds, sqlstr, roleId)
	}

	if err != nil {
		panic(err)
	}
	return
}
