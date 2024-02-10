package service

import (
	"baize/app/baize"
	"baize/app/bzSystem/models"
	"github.com/gin-gonic/gin"
)

type IRoleService interface {
	SelectRoleList(c *gin.Context, role *models.SysRoleDQL) (list []*models.SysRoleVo, count *int64)
	RoleExport(c *gin.Context, role *models.SysRoleDQL) (data []byte)
	SelectRoleById(c *gin.Context, roseId int64) (role *models.SysRoleVo)
	InsertRole(c *gin.Context, sysRole *models.SysRoleDML)
	UpdateRole(c *gin.Context, sysRole *models.SysRoleDML)
	UpdateRoleStatus(c *gin.Context, sysRole *models.SysRoleDML)
	AuthDataScope(c *gin.Context, sysRole *models.SysRoleDML)
	DeleteRoleByIds(c *gin.Context, ids []int64)
	CountUserRoleByRoleId(c *gin.Context, ids []int64) bool
	SelectBasicRolesByUserId(c *gin.Context, userId int64) (roles []*models.SysRole)
	RolePermissionByRoles(c *gin.Context, roles []*models.SysRole) (rolePerms []string, loginRoles []*baize.Role)

	CheckRoleNameUnique(c *gin.Context, id int64, roleName string) bool
	CheckRoleKeyUnique(c *gin.Context, id int64, roleKey string) bool
	SelectUserRoleGroupByUserId(c *gin.Context, userId int64) string
	SelectAllocatedList(c *gin.Context, user *models.SysRoleAndUserDQL) (list []*models.SysUserVo, total *int64)
	SelectUnallocatedList(c *gin.Context, user *models.SysRoleAndUserDQL) (list []*models.SysUserVo, total *int64)
	InsertAuthUsers(c *gin.Context, roleId int64, userIds []int64)
	DeleteAuthUsers(c *gin.Context, roleId int64, userIds []int64)
	DeleteAuthUserRole(c *gin.Context, user *models.SysUserRole)
}
