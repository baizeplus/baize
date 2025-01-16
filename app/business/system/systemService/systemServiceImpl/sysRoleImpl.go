package systemServiceImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemDao/systemDaoImpl"
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/utils/excel"
	"github.com/baizeplus/sqly"

	"baize/app/utils/snowflake"
	"github.com/gin-gonic/gin"
	"strconv"
)

type RoleService struct {
	ms          sqly.SqlyContext
	roleDao     systemDao.IRoleDao
	roleMenuDao systemDao.IRoleMenuDao
	userRoleDao systemDao.IUserRoleDao
}

func NewRoleService(ms sqly.SqlyContext, rd systemDao.IRoleDao, rmd systemDao.IRoleMenuDao, urd systemDao.IUserRoleDao) systemService.IRoleService {
	return &RoleService{
		ms:          ms,
		roleDao:     rd,
		roleMenuDao: rmd,
		userRoleDao: urd,
	}
}

func (roleService *RoleService) SelectRoleList(c *gin.Context, role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo, total int64) {
	return roleService.roleDao.SelectRoleList(c, role)

}
func (roleService *RoleService) RoleExport(c *gin.Context, role *systemModels.SysRoleDQL) (data []byte) {

	list := roleService.roleDao.SelectRoleAll(c, role)
	toExcel, err := excel.SliceToExcel(list)
	if err != nil {
		panic(err)
	}
	buffer, err := toExcel.WriteToBuffer()
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

func (roleService *RoleService) SelectRoleById(c *gin.Context, roseId int64) (role *systemModels.SysRoleVo) {
	return roleService.roleDao.SelectRoleById(c, roseId)

}

func (roleService *RoleService) InsertRole(c *gin.Context, sysRole *systemModels.SysRoleDML) {
	sysRole.RoleId = snowflake.GenID()
	tx := roleService.ms.MustBeginTx(c, nil)

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else {
			tx.Commit()
		}
	}()
	rd := systemDaoImpl.NewSysRoleDao(tx)
	rd.InsertRole(c, sysRole)
	menuIds := sysRole.MenuIds
	l := len(menuIds)
	if l != 0 {
		list := make([]*systemModels.SysRoleMenu, 0, l)
		for _, PermissionId := range menuIds {
			intPermissionId, _ := strconv.ParseInt(PermissionId, 10, 64)
			list = append(list, &systemModels.SysRoleMenu{RoleId: sysRole.RoleId, MenuId: intPermissionId})
		}
		if len(list) != 0 {
			md := systemDaoImpl.NewSysRoleMenuDao(tx)
			md.BatchRoleMenu(c, list)
		}
	}

	return
}

func (roleService *RoleService) UpdateRole(c *gin.Context, sysRole *systemModels.SysRoleDML) {
	tx := roleService.ms.MustBeginTx(c, nil)

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else {
			tx.Commit()
		}
	}()
	rd := systemDaoImpl.NewSysRoleDao(tx)
	rd.UpdateRole(c, sysRole)
	md := systemDaoImpl.NewSysRoleMenuDao(tx)
	md.DeleteRoleMenuByRoleId(c, sysRole.RoleId)
	menuIds := sysRole.MenuIds
	l := len(menuIds)
	if l != 0 {
		list := make([]*systemModels.SysRoleMenu, 0, l)
		for _, PermissionId := range menuIds {
			intPermissionId, _ := strconv.ParseInt(PermissionId, 10, 64)
			list = append(list, &systemModels.SysRoleMenu{RoleId: sysRole.RoleId, MenuId: intPermissionId})
		}
		if len(list) != 0 {
			md.BatchRoleMenu(c, list)
		}
	}
	return
}

func (roleService *RoleService) UpdateRoleStatus(c *gin.Context, sysRole *systemModels.SysRoleDML) {
	roleService.roleDao.UpdateRole(c, sysRole)
	return
}

func (roleService *RoleService) DeleteRoleByIds(c *gin.Context, ids []int64) {
	tx := roleService.ms.MustBeginTx(c, nil)
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else {
			tx.Commit()
		}
	}()
	rd := systemDaoImpl.NewSysRoleDao(tx)
	md := systemDaoImpl.NewSysRoleMenuDao(tx)
	md.DeleteRoleMenu(c, ids)
	rd.DeleteRoleByIds(c, ids)
}
func (roleService *RoleService) CountUserRoleByRoleId(c *gin.Context, ids []int64) bool {
	return roleService.userRoleDao.CountUserRoleByRoleId(c, ids) > 0
}

func (roleService *RoleService) SelectBasicRolesByUserId(c *gin.Context, userId int64) (roles []*systemModels.SysRole) {
	return roleService.roleDao.SelectBasicRolesByUserId(c, userId)

}

func (roleService *RoleService) RolePermissionByRoles(c *gin.Context, roles []*systemModels.SysRole) (rolePerms []string, loginRoles []int64) {
	loginRoles = make([]int64, 0, len(roles))
	rolePerms = make([]string, 0, len(roles))
	for _, role := range roles {
		rolePerms = append(rolePerms, role.RoleKey)
		loginRoles = append(loginRoles, role.RoleId)
	}
	return
}

func (roleService *RoleService) CheckRoleNameUnique(c *gin.Context, id int64, roleName string) bool {
	RoleId := roleService.roleDao.CheckRoleNameUnique(c, roleName)
	if RoleId == id || RoleId == 0 {
		return false
	}
	return true
}

func (roleService *RoleService) CheckRoleKeyUnique(c *gin.Context, id int64, roleKey string) bool {
	RoleId := roleService.roleDao.CheckRoleKeyUnique(c, roleKey)
	if RoleId == id || RoleId == 0 {
		return false
	}
	return true
}

func (roleService *RoleService) SelectAllocatedList(c *gin.Context, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total int64) {
	return roleService.roleDao.SelectAllocatedList(c, user)
}

func (roleService *RoleService) SelectUnallocatedList(c *gin.Context, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total int64) {
	return roleService.roleDao.SelectUnallocatedList(c, user)

}

func (roleService *RoleService) InsertAuthUsers(c *gin.Context, roleId int64, userIds []int64) {
	if len(userIds) != 0 {
		list := make([]*systemModels.SysUserRole, 0, len(userIds))
		for _, userId := range userIds {
			role := systemModels.NewSysUserRole(userId, roleId)
			list = append(list, role)
		}
		roleService.userRoleDao.BatchUserRole(c, list)
	}
}
func (roleService *RoleService) DeleteAuthUsers(c *gin.Context, roleId int64, userIds []int64) {
	roleService.userRoleDao.DeleteUserRoleInfos(c, roleId, userIds)
}
func (roleService *RoleService) DeleteAuthUserRole(c *gin.Context, userRole *systemModels.SysUserRole) {
	roleService.userRoleDao.DeleteUserRoleInfo(c, userRole)
}
