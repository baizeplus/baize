package controller

import (
	"baize/app/bzSystem/models"
	"baize/app/bzSystem/service"
	"baize/app/bzSystem/service/serviceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type RoleController struct {
	rs service.IRoleService
}

func NewRoleController(rs *serviceImpl.RoleService) *RoleController {
	return &RoleController{rs: rs}
}

func (rc *RoleController) RoleList(c *gin.Context) {
	role := new(models.SysRoleDQL)
	_ = c.ShouldBind(role)
	role.DataScope = baizeContext.GetDataScope(c, "d")
	list, count := rc.rs.SelectRoleList(c, role)
	baizeContext.SuccessListData(c, list, count)
}

func (rc *RoleController) RoleExport(c *gin.Context) {
	role := new(models.SysRoleDQL)
	_ = c.ShouldBind(role)
	role.DataScope = baizeContext.GetDataScope(c, "d")
	list, count := rc.rs.SelectRoleList(c, role)
	baizeContext.SuccessListData(c, list, count)
}
func (rc *RoleController) RoleGetInfo(c *gin.Context) {
	roleId := baizeContext.ParamInt64(c, "roleId")
	if roleId == 0 {
		baizeContext.ParameterError(c)
		return
	}
	sysUser := rc.rs.SelectRoleById(c, roleId)
	baizeContext.SuccessData(c, sysUser)
}
func (rc *RoleController) RoleAdd(c *gin.Context) {
	sysRole := new(models.SysRoleDML)
	_ = c.ShouldBindJSON(sysRole)
	if rc.rs.CheckRoleNameUnique(c, 0, sysRole.RoleName) {
		baizeContext.Waring(c, "新增角色'"+sysRole.RoleName+"'失败，角色名称已存在")
		return
	}
	if rc.rs.CheckRoleKeyUnique(c, 0, sysRole.RoleKey) {
		baizeContext.Waring(c, "新增角色'"+sysRole.RoleKey+"'失败，角色权限已存在")
		return
	}
	sysRole.SetCreateBy(baizeContext.GetUserId(c))
	rc.rs.InsertRole(c, sysRole)
	baizeContext.Success(c)

}
func (rc *RoleController) RoleEdit(c *gin.Context) {
	sysRole := new(models.SysRoleDML)
	_ = c.ShouldBindJSON(sysRole)
	if rc.rs.CheckRoleNameUnique(c, sysRole.RoleId, sysRole.RoleName) {
		baizeContext.Waring(c, "新增角色'"+sysRole.RoleName+"'失败，角色名称已存在")
		return
	}
	if rc.rs.CheckRoleKeyUnique(c, sysRole.RoleId, sysRole.RoleKey) {
		baizeContext.Waring(c, "新增角色'"+sysRole.RoleKey+"'失败，角色权限已存在")
		return
	}
	sysRole.SetUpdateBy(baizeContext.GetUserId(c))
	rc.rs.UpdateRole(c, sysRole)
	baizeContext.Success(c)
}
func (rc *RoleController) RoleDataScope(c *gin.Context) {
	sysRole := new(models.SysRoleDML)
	_ = c.ShouldBindJSON(sysRole)
	sysRole.SetUpdateBy(baizeContext.GetUserId(c))
	rc.rs.AuthDataScope(c, sysRole)
	baizeContext.Success(c)
}

func (rc *RoleController) RoleChangeStatus(c *gin.Context) {
	sysRole := new(models.SysRoleDML)
	_ = c.ShouldBindJSON(sysRole)
	sysRole.SetUpdateBy(baizeContext.GetUserId(c))
	rc.rs.UpdateRoleStatus(c, sysRole)
	baizeContext.Success(c)
}
func (rc *RoleController) RoleRemove(c *gin.Context) {
	ids := baizeContext.ParamInt64Array(c, "rolesIds")
	if rc.rs.CountUserRoleByRoleId(c, ids) {
		baizeContext.Waring(c, "角色已分配，不能删除")
		return
	}
	rc.rs.DeleteRoleByIds(c, ids)
	baizeContext.Success(c)
}
func (rc *RoleController) AllocatedList(c *gin.Context) {
	user := new(models.SysRoleAndUserDQL)
	if err := c.ShouldBind(user); err != nil {
		baizeContext.ParameterError(c)
		return
	}
	user.DataScope = baizeContext.GetDataScope(c, "d")
	list, count := rc.rs.SelectAllocatedList(c, user)
	baizeContext.SuccessListData(c, list, count)

}
func (rc *RoleController) UnallocatedList(c *gin.Context) {
	user := new(models.SysRoleAndUserDQL)
	if err := c.ShouldBind(user); err != nil {
		baizeContext.ParameterError(c)
		return
	}
	user.DataScope = baizeContext.GetDataScope(c, "d")
	list, count := rc.rs.SelectUnallocatedList(c, user)
	baizeContext.SuccessListData(c, list, count)
}
func (rc *RoleController) InsertAuthUser(c *gin.Context) {
	rc.rs.InsertAuthUsers(c, baizeContext.QueryInt64(c, "roleId"), baizeContext.QueryInt64Array(c, "userIds"))
	baizeContext.Success(c)
}
func (rc *RoleController) CancelAuthUser(c *gin.Context) {
	userRole := new(models.SysUserRole)
	if err := c.ShouldBindJSON(userRole); err != nil {
		baizeContext.ParameterError(c)
		return
	}
	rc.rs.DeleteAuthUserRole(c, userRole)
	baizeContext.Success(c)
}
func (rc *RoleController) CancelAuthUserAll(c *gin.Context) {
	rc.rs.DeleteAuthUsers(c, baizeContext.QueryInt64(c, "roleId"), baizeContext.QueryInt64Array(c, "userIds"))
	baizeContext.Success(c)
}
