package systemService

import (
	"baize/app/business/system/systemModels"
	"github.com/gin-gonic/gin"
)

type IRoleService interface {
	SelectRoleList(c *gin.Context, role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo, total int64)
	RoleExport(c *gin.Context, role *systemModels.SysRoleDQL) (data []byte)
	SelectRoleById(c *gin.Context, roseId int64) (role *systemModels.SysRoleVo)
	InsertRole(c *gin.Context, sysRole *systemModels.SysRoleDML)
	UpdateRole(c *gin.Context, sysRole *systemModels.SysRoleDML)
	UpdateRoleStatus(c *gin.Context, sysRole *systemModels.SysRoleDML)
	DeleteRoleByIds(c *gin.Context, ids []int64)
	CountUserRoleByRoleId(c *gin.Context, ids []int64) bool
	SelectBasicRolesByUserId(c *gin.Context, userId int64) (roles []*systemModels.SysRole)
	RolePermissionByRoles(c *gin.Context, roles []*systemModels.SysRole) (rolePerms []string, loginRoles []int64)

	CheckRoleNameUnique(c *gin.Context, id int64, roleName string) bool
	CheckRoleKeyUnique(c *gin.Context, id int64, roleKey string) bool
	SelectAllocatedList(c *gin.Context, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total int64)
	SelectUnallocatedList(c *gin.Context, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total int64)
	InsertAuthUsers(c *gin.Context, roleId int64, userIds []int64)
	DeleteAuthUsers(c *gin.Context, roleId int64, userIds []int64)
	DeleteAuthUserRole(c *gin.Context, user *systemModels.SysUserRole)
}
