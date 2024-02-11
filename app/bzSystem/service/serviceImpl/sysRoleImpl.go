package serviceImpl

import (
	"baize/app/baize"
	"baize/app/bzSystem/dao"
	"baize/app/bzSystem/dao/daoImpl"
	"baize/app/bzSystem/models"
	"baize/app/utils/snowflake"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type RoleService struct {
	data        *sqly.DB
	roleDao     dao.IRoleDao
	roleMenuDao dao.IRoleMenuDao
	roleDeptDao dao.IRoleDeptDao
	userRoleDao dao.IUserRoleDao
}

func NewRoleService(data *sqly.DB, rd *daoImpl.SysRoleDao, rmd *daoImpl.SysRoleMenuDao, rdd *daoImpl.SysRoleDeptDao, urd *daoImpl.SysUserRoleDao) *RoleService {
	return &RoleService{
		data:        data,
		roleDao:     rd,
		roleMenuDao: rmd,
		roleDeptDao: rdd,
		userRoleDao: urd,
	}
}

func (roleService *RoleService) SelectRoleList(c *gin.Context, role *models.SysRoleDQL) (list []*models.SysRoleVo, count *int64) {
	return roleService.roleDao.SelectRoleList(c, roleService.data, role)

}
func (roleService *RoleService) RoleExport(c *gin.Context, role *models.SysRoleDQL) (data []byte) {
	//list, _ := roleService.roleDao.SelectRoleList(roleService.data.GetSlaveDb(), role)
	//rows := models.SysRoleListToRows(list)
	//return exceLize.SetRows(rows)
	return nil
}

func (roleService *RoleService) SelectRoleById(c *gin.Context, roseId int64) (role *models.SysRoleVo) {
	return roleService.roleDao.SelectRoleById(c, roleService.data, roseId)

}

func (roleService *RoleService) InsertRole(c *gin.Context, sysRole *models.SysRoleDML) {
	sysRole.RoleId = snowflake.GenID()
	tx, err := roleService.data.Beginx()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else {
			err = tx.Commit()
		}
	}()
	roleService.roleDao.InsertRole(c, tx, sysRole)
	PermissionIds := sysRole.PermissionIds
	l := len(PermissionIds)
	if l != 0 {
		list := make([]*models.SysRoleMenu, 0, l)
		for _, PermissionId := range PermissionIds {
			intPermissionId, _ := strconv.ParseInt(PermissionId, 10, 64)
			list = append(list, &models.SysRoleMenu{RoleId: sysRole.RoleId, MenuId: intPermissionId})
		}
		roleService.insertRolePermission(c, tx, list)
	}

	return
}

func (roleService *RoleService) UpdateRole(c *gin.Context, sysRole *models.SysRoleDML) {
	tx, err := roleService.data.Beginx()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else {
			err = tx.Commit()
		}
	}()
	roleService.roleDao.UpdateRole(c, tx, sysRole)
	//todo
	//roleService.rolePermissionDao.DeleteRolePermissionByRoleId(tx, sysRole.RoleId)
	//PermissionIds := sysRole.PermissionIds
	//l := len(PermissionIds)
	//if l != 0 {
	//	list := make([]*models.SysRolePermission, 0, l)
	//	for _, PermissionId := range PermissionIds {
	//		intPermissionId, _ := strconv.ParseInt(PermissionId, 10, 64)
	//		list = append(list, &models.SysRolePermission{RoleId: sysRole.RoleId, PermissionId: intPermissionId})
	//	}
	//	roleService.insertRolePermission(c,tx, list)
	//}
	return
}

func (roleService *RoleService) UpdateRoleStatus(c *gin.Context, sysRole *models.SysRoleDML) {
	roleService.roleDao.UpdateRole(c, roleService.data, sysRole)
	return
}
func (roleService *RoleService) AuthDataScope(c *gin.Context, sysRole *models.SysRoleDML) {
	tx, err := roleService.data.Beginx()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else {
			err = tx.Commit()
		}
	}()
	roleService.roleDao.UpdateRole(c, tx, sysRole)
	roleService.roleDeptDao.DeleteRoleDeptByRoleId(c, tx, sysRole.RoleId)
	roleService.insertRoleDept(c, tx, sysRole)
	return
}

func (roleService *RoleService) DeleteRoleByIds(c *gin.Context, ids []int64) {
	tx, err := roleService.data.Beginx()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else {
			err = tx.Commit()
		}
	}()
	//todo
	//roleService.rolePermissionDao.DeleteRolePermission(tx, ids)
	roleService.roleDeptDao.DeleteRoleDept(c, tx, ids)
	roleService.roleDao.DeleteRoleByIds(c, tx, ids)
}
func (roleService *RoleService) CountUserRoleByRoleId(c *gin.Context, ids []int64) bool {
	return roleService.userRoleDao.CountUserRoleByRoleId(c, roleService.data, ids) > 0
}

func (roleService *RoleService) SelectBasicRolesByUserId(c *gin.Context, userId int64) (roles []*models.SysRole) {
	return roleService.roleDao.SelectBasicRolesByUserId(c, roleService.data, userId)

}

func (roleService *RoleService) RolePermissionByRoles(c *gin.Context, roles []*models.SysRole) (rolePerms []string, loginRoles []*baize.Role) {
	loginRoles = make([]*baize.Role, 0, len(roles))
	rolePerms = make([]string, 0, len(roles))
	for _, role := range roles {
		rolePerms = append(rolePerms, role.RoleKey)
		//todo
		//loginRoles = append(loginRoles, &baizeEntity.Role{RoleId: role.RoleId, DataScope: role.DataScope})
	}
	return
}

func (roleService *RoleService) insertRolePermission(c *gin.Context, db sqly.SqlyContext, list []*models.SysRoleMenu) {
	if len(list) != 0 {
		roleService.roleMenuDao.BatchRoleMenu(c, db, list)
	}
	return
}

func (roleService *RoleService) CheckRoleNameUnique(c *gin.Context, id int64, roleName string) bool {
	RoleId := roleService.roleDao.CheckRoleNameUnique(c, roleService.data, roleName)
	if RoleId == id || RoleId == 0 {
		return false
	}
	return true
}

func (roleService *RoleService) CheckRoleKeyUnique(c *gin.Context, id int64, roleKey string) bool {
	RoleId := roleService.roleDao.CheckRoleKeyUnique(c, roleService.data, roleKey)
	if RoleId == id || RoleId == 0 {
		return false
	}
	return true
}
func (roleService *RoleService) SelectUserRoleGroupByUserId(c *gin.Context, userId int64) string {
	roles := roleService.roleDao.SelectBasicRolesByUserId(c, roleService.data, userId)
	roleNames := make([]string, 0, len(roles))
	for _, role := range roles {
		roleNames = append(roleNames, role.RoleName)
	}
	return strings.Join(roleNames, ",")

}
func (roleService *RoleService) insertRoleDept(c *gin.Context, db sqly.SqlyContext, sysRole *models.SysRoleDML) {
	deptIds := sysRole.DeptIds
	if len(deptIds) != 0 {
		list := make([]*models.SysRoleDept, 0, len(deptIds))
		for _, deptId := range deptIds {
			intDeptId, _ := strconv.ParseInt(deptId, 10, 64)
			list = append(list, &models.SysRoleDept{RoleId: sysRole.RoleId, DeptId: intDeptId})
		}
		roleService.roleDeptDao.BatchRoleDept(c, db, list)
	}

}
func (roleService *RoleService) SelectAllocatedList(c *gin.Context, user *models.SysRoleAndUserDQL) (list []*models.SysUserVo, total *int64) {
	return roleService.roleDao.SelectAllocatedList(c, roleService.data, user)
}

func (roleService *RoleService) SelectUnallocatedList(c *gin.Context, user *models.SysRoleAndUserDQL) (list []*models.SysUserVo, total *int64) {
	return roleService.roleDao.SelectUnallocatedList(c, roleService.data, user)

}

func (roleService *RoleService) InsertAuthUsers(c *gin.Context, roleId int64, userIds []int64) {
	if len(userIds) != 0 {
		list := make([]*models.SysUserRole, 0, len(userIds))
		for _, userId := range userIds {
			role := models.NewSysUserRole(userId, roleId)
			list = append(list, role)
		}
		roleService.userRoleDao.BatchUserRole(c, roleService.data, list)
	}
}
func (roleService *RoleService) DeleteAuthUsers(c *gin.Context, roleId int64, userIds []int64) {
	roleService.userRoleDao.DeleteUserRoleInfos(c, roleService.data, roleId, userIds)
}
func (roleService *RoleService) DeleteAuthUserRole(c *gin.Context, userRole *models.SysUserRole) {
	roleService.userRoleDao.DeleteUserRoleInfo(c, roleService.data, userRole)
}
