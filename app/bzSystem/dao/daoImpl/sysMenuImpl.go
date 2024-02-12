package daoImpl

import (
	"baize/app/bzSystem/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/baizeplus/sqly"
)

type SysMenuDao struct {
	selectMenuSql string
}

func NewSysMenuDao() *SysMenuDao {
	return &SysMenuDao{
		selectMenuSql: `select distinct m.menu_id, m.parent_id, m.menu_name, m.path, m.component, m.visible, m.status, ifnull(m.perms,'') as perms, m.is_frame, m.is_cache, m.menu_type, m.icon, m.order_num, m.create_time from sys_menu m`,
	}
}

func (sysMenuDao *SysMenuDao) SelectMenuById(ctx context.Context, db sqly.SqlyContext, menuId int64) (menu *models.SysMenuVo) {
	whereSql := ` where menu_id = ?`
	menu = new(models.SysMenuVo)
	err := db.GetContext(ctx, menu, sysMenuDao.selectMenuSql+whereSql, menuId)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}

func (sysMenuDao *SysMenuDao) SelectMenuList(ctx context.Context, db sqly.SqlyContext, menu *models.SysMenuDQL) (list []*models.SysMenuVo) {
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

	list = make([]*models.SysMenuVo, 0, 16)
	err := db.NamedSelectContext(ctx, &list, sysMenuDao.selectMenuSql+whereSql, menu)

	if err != nil {
		panic(err)
	}

	return list
}

func (sysMenuDao *SysMenuDao) SelectMenuListByUserId(ctx context.Context, db sqly.SqlyContext, menu *models.SysMenuDQL) (list []*models.SysMenuVo) {
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

	list = make([]*models.SysMenuVo, 0, 16)
	err := db.NamedSelectContext(ctx, &list, sysMenuDao.selectMenuSql+whereSql, menu)

	if err != nil {
		panic(err)
	}

	return list
}

func (sysMenuDao *SysMenuDao) InsertMenu(ctx context.Context, db sqly.SqlyContext, menu *models.SysMenuVo) {
	insertSQL := `insert into sys_menu(menu_id,menu_name,parent_id,create_by,create_time,update_by,update_time %s)
					values(:menu_id,:menu_name,:parent_id,:create_by,now(),:update_by,now() %s)`
	key := ""
	value := ""
	if menu.OrderNum != 0 {
		key += ",order_num"
		value += ",:order_num"
	}
	if menu.Path != "" {
		key += ",path"
		value += ",:path"
	}
	if menu.Component != "" {
		key += ",component"
		value += ",:component"
	}
	if menu.IsFrame != "" {
		key += ",is_frame"
		value += ",:is_frame"
	}
	if menu.IsCache != "" {
		key += ",is_cache"
		value += ",:is_cache"
	}
	if menu.MenuType != "" {
		key += ",menu_type"
		value += ",:menu_type"
	}
	if menu.Visible != "" {
		key += ",visible"
		value += ",:visible"
	}
	if menu.Status != "" {
		key += ",status"
		value += ",:status"
	}
	if menu.Perms != "" {
		key += ",Perms"
		value += ",:perms"
	}
	if menu.Icon != "" {
		key += ",icon"
		value += ",:icon"
	}
	if menu.Remark != "" {
		key += ",remark"
		value += ",:remark"
	}

	insertStr := fmt.Sprintf(insertSQL, key, value)
	_, err := db.NamedExecContext(ctx, insertStr, menu)
	if err != nil {
		panic(err)
	}
	return
}

func (sysMenuDao *SysMenuDao) UpdateMenu(ctx context.Context, db sqly.SqlyContext, menu *models.SysMenuVo) {
	updateSQL := `update sys_menu set update_time = now() , update_by = :update_by`

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

	_, err := db.NamedExecContext(ctx, updateSQL, menu)
	if err != nil {
		panic(err)
	}
	return
}

func (sysMenuDao *SysMenuDao) DeleteMenuById(ctx context.Context, db sqly.SqlyContext, menuId int64) {
	_, err := db.ExecContext(ctx, "delete from sys_menu where menu_id = ?", menuId)
	if err != nil {
		panic(err)
	}
	return
}

func (sysMenuDao *SysMenuDao) SelectMenuPermsByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (perms []string) {
	sqlStr := `	select distinct m.perms
				from sys_menu m
				left join sys_role_menu rm on m.menu_id = rm.menu_id
				left join sys_user_role ur on rm.role_id = ur.role_id
				left join sys_role r on r.role_id = ur.role_id
				where m.status = '0' and r.status = '0' and ur.user_id =  ?`
	perms = make([]string, 0, 2)
	err := db.SelectContext(ctx, &perms, sqlStr, userId)
	if err != nil {
		panic(err)
	}
	return
}

func (sysMenuDao *SysMenuDao) SelectMenuTreeAll(ctx context.Context, db sqly.SqlyContext) (sysMenus []*models.SysMenuVo) {
	whereSql := ` where m.menu_type in ('M', 'C') and m.status = 0
		order by m.parent_id, m.order_num`
	sysMenus = make([]*models.SysMenuVo, 0, 2)
	err := db.SelectContext(ctx, &sysMenus, sysMenuDao.selectMenuSql+whereSql)
	if err != nil {
		panic(err)
	}
	return
}
func (sysMenuDao *SysMenuDao) SelectMenuTreeByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) (sysMenus []*models.SysMenuVo) {
	whereSql := ` left join sys_role_menu rm on m.menu_id = rm.menu_id
			 left join sys_user_role ur on rm.role_id = ur.role_id
			 left join sys_role ro on ur.role_id = ro.role_id
			 left join sys_user u on ur.user_id = u.user_id
		where u.user_id = ? and m.menu_type in ('M', 'C') and m.status = 0  AND ro.status = 0
		order by m.parent_id, m.order_num`
	sysMenus = make([]*models.SysMenuVo, 0, 2)
	err := db.SelectContext(ctx, &sysMenus, sysMenuDao.selectMenuSql+whereSql, userId)
	if err != nil {
		panic(err)
	}
	return
}

func (sysMenuDao *SysMenuDao) CheckMenuNameUnique(ctx context.Context, db sqly.SqlyContext, menuName string, parentId int64) int64 {
	var roleId int64 = 0
	err := db.GetContext(ctx, &roleId, "select menu_id from sys_menu where menu_name=? and parent_id = ?", menuName, parentId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return roleId
}

func (sysMenuDao *SysMenuDao) HasChildByMenuId(ctx context.Context, db sqly.SqlyContext, menuId int64) int {
	var count = 0
	err := db.GetContext(ctx, &count, "select count(*) from sys_menu where parent_id = ?", menuId)
	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}
	return count
}
func (sysMenuDao *SysMenuDao) SelectMenuListByRoleId(ctx context.Context, db sqly.SqlyContext, roleId int64, menuCheckStrictly bool) (roleIds []string) {
	var err error
	roleIds = make([]string, 0, 2)
	sqlstr := `select m.menu_id
		from sys_menu m
            left join sys_role_menu rm on m.menu_id = rm.menu_id
        where rm.role_id = ?`
	if menuCheckStrictly {
		sqlstr += " and m.menu_id not in (select m.parent_id from sys_menu m inner join sys_role_menu rm on m.menu_id = rm.menu_id and rm.role_id = ?)"
		err = db.SelectContext(ctx, &roleIds, sqlstr, roleId, roleId)
	} else {
		err = db.SelectContext(ctx, &roleIds, sqlstr, roleId)
	}

	if err != nil {
		panic(err)
	}
	return
}
