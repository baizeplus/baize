package systemServiceImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemDao/systemDaoImpl"
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/utils/baizeId"
	"baize/app/utils/excel"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
)

type RoleService struct {
	ms          sqly.SqlyContext
	roleDao     systemDao.IRoleDao
	rpDao       systemDao.IRolePermissionDao
	userRoleDao systemDao.IUserRoleDao
}

func NewRoleService(ms sqly.SqlyContext, rd systemDao.IRoleDao, rmd systemDao.IRolePermissionDao, urd systemDao.IUserRoleDao) systemService.IRoleService {
	return &RoleService{
		ms:          ms,
		roleDao:     rd,
		rpDao:       rmd,
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

func (roleService *RoleService) SelectRoleById(c *gin.Context, roleId string) (role *systemModels.SysRoleVo) {
	role = roleService.roleDao.SelectRoleById(c, roleId)
	if role.RoleId != "" {
		role.PermissionIds = roleService.rpDao.SelectPermissionIdsByRoleId(c, roleId)
	}
	return role

}

func (roleService *RoleService) InsertRole(c *gin.Context, sysRole *systemModels.SysRoleDML) {
	sysRole.RoleId = baizeId.GetId()
	sysRole.DelFlag = "0"
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
	pIds := sysRole.PermissionIds
	l := len(pIds)
	if l != 0 {
		list := make([]*systemModels.SysRolePermission, 0, l)
		for _, permissionId := range pIds {
			list = append(list, &systemModels.SysRolePermission{RoleId: sysRole.RoleId, PermissionId: permissionId})
		}
		if len(list) != 0 {
			md := systemDaoImpl.NewSysRolePermissionDao(tx)
			md.BatchRolePermission(c, list)
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
	md := systemDaoImpl.NewSysRolePermissionDao(tx)
	md.DeleteRolePermissionByRoleId(c, sysRole.RoleId)
	menuIds := sysRole.PermissionIds
	l := len(menuIds)
	if l != 0 {
		list := make([]*systemModels.SysRolePermission, 0, l)
		for _, permissionId := range menuIds {
			list = append(list, &systemModels.SysRolePermission{RoleId: sysRole.RoleId, PermissionId: permissionId})
		}
		if len(list) != 0 {
			md.BatchRolePermission(c, list)
		}
	}
	return
}

func (roleService *RoleService) UpdateRoleStatus(c *gin.Context, sysRole *systemModels.SysRoleDML) {
	roleService.roleDao.UpdateRole(c, sysRole)
	return
}

func (roleService *RoleService) DeleteRoleByIds(c *gin.Context, ids []string) {
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
	md := systemDaoImpl.NewSysRolePermissionDao(tx)
	md.DeleteRolePermission(c, ids)
	rd.DeleteRoleByIds(c, ids)
}
func (roleService *RoleService) CountUserRoleByRoleId(c *gin.Context, ids []string) bool {
	return roleService.userRoleDao.CountUserRoleByRoleId(c, ids) > 0
}

func (roleService *RoleService) SelectBasicRolesByUserId(c *gin.Context, userId string) (roles []*systemModels.SysRole) {
	return roleService.roleDao.SelectBasicRolesByUserId(c, userId)

}

func (roleService *RoleService) RolePermissionByRoles(c *gin.Context, roles []*systemModels.SysRole) (loginRoles []string) {
	loginRoles = make([]string, 0, len(roles))
	for _, role := range roles {
		loginRoles = append(loginRoles, role.RoleId)
	}
	return
}

func (roleService *RoleService) CheckRoleNameUnique(c *gin.Context, id string, roleName string) bool {
	RoleId := roleService.roleDao.CheckRoleNameUnique(c, roleName)
	if RoleId == id || RoleId == "" {
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

func (roleService *RoleService) InsertAuthUsers(c *gin.Context, roleId string, userIds []string) {
	if len(userIds) != 0 {
		list := make([]*systemModels.SysUserRole, 0, len(userIds))
		for _, userId := range userIds {
			role := systemModels.NewSysUserRole(userId, roleId)
			list = append(list, role)
		}
		roleService.userRoleDao.BatchUserRole(c, list)
	}
}
func (roleService *RoleService) DeleteAuthUsers(c *gin.Context, roleId string, userIds []string) {
	roleService.userRoleDao.DeleteUserRoleInfos(c, roleId, userIds)
}
func (roleService *RoleService) DeleteAuthUserRole(c *gin.Context, userRole *systemModels.SysUserRole) {
	roleService.userRoleDao.DeleteUserRoleInfo(c, userRole)
}
