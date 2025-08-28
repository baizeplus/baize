package systemService

import (
	"baize/app/business/system/systemModels"

	"github.com/gin-gonic/gin"
)

type IRoleService interface {
	SelectRoleList(c *gin.Context, role *systemModels.SysRoleDQL) (list []*systemModels.SysRoleVo, total int64)
	RoleExport(c *gin.Context, role *systemModels.SysRoleDQL) (data []byte)
	SelectRoleById(c *gin.Context, roseId string) (role *systemModels.SysRoleVo)
	InsertRole(c *gin.Context, sysRole *systemModels.SysRoleDML)
	UpdateRole(c *gin.Context, sysRole *systemModels.SysRoleDML)
	UpdateRoleStatus(c *gin.Context, sysRole *systemModels.SysRoleDML)
	DeleteRoleByIds(c *gin.Context, ids []string)
	CountUserRoleByRoleId(c *gin.Context, ids []string) bool
	SelectBasicRolesByUserId(c *gin.Context, userId string) (roles []*systemModels.SysRole)
	RolePermissionByRoles(c *gin.Context, roles []*systemModels.SysRole) (loginRoles []int64)

	CheckRoleNameUnique(c *gin.Context, id string, roleName string) bool
	SelectAllocatedList(c *gin.Context, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total int64)
	SelectUnallocatedList(c *gin.Context, user *systemModels.SysRoleAndUserDQL) (list []*systemModels.SysUserVo, total int64)
	InsertAuthUsers(c *gin.Context, roleId string, userIds []string)
	DeleteAuthUsers(c *gin.Context, roleId string, userIds []string)
	DeleteAuthUserRole(c *gin.Context, user *systemModels.SysUserRole)
}
