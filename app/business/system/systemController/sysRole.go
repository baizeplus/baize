package systemController

import (
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/business/system/systemService/systemServiceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type Role struct {
	rs systemService.IRoleService
}

func NewRole(rs *systemServiceImpl.RoleService) *Role {
	return &Role{rs: rs}
}

// RoleList 查询角色列表查询
// @Summary 查询角色列表查询
// @Description 查询角色列表查询
// @Tags 角色相关
// @Param  object query systemModels.SysRoleDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ListData{Rows=[]systemModels.SysRoleVo}}  "成功"
// @Router /system/role/list  [get]
func (rc *Role) RoleList(c *gin.Context) {
	role := new(systemModels.SysRoleDQL)
	_ = c.ShouldBind(role)
	list, count := rc.rs.SelectRoleList(c, role)
	baizeContext.SuccessListData(c, list, count)
}

func (rc *Role) RoleExport(c *gin.Context) {
	role := new(systemModels.SysRoleDQL)
	_ = c.ShouldBind(role)
	baizeContext.DataPackageExcel(c, rc.rs.RoleExport(c, role))
}

// RoleGetInfo 根据角色ID获取角色信息
// @Summary 根据角色ID获取角色信息
// @Description 根据角色ID获取角色信息
// @Tags 角色相关
// @Param id path string true "roleId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=systemModels.SysRoleVo}  "成功"
// @Router /system/role/{roleId}  [get]
func (rc *Role) RoleGetInfo(c *gin.Context) {
	roleId := baizeContext.ParamInt64(c, "roleId")
	if roleId == 0 {
		baizeContext.ParameterError(c)
		return
	}
	sysUser := rc.rs.SelectRoleById(c, roleId)
	baizeContext.SuccessData(c, sysUser)
}

// RoleAdd 添加角色
// @Summary 添加角色
// @Description 添加角色
// @Tags 角色相关
// @Param  object body systemModels.SysRoleDML true "公司信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/role  [post]
func (rc *Role) RoleAdd(c *gin.Context) {
	sysRole := new(systemModels.SysRoleDML)
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

// RoleEdit 修改角色
// @Summary 修改角色
// @Description 修改角色
// @Tags 角色相关
// @Param  object body systemModels.SysRoleDML true "公司信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/role  [put]
func (rc *Role) RoleEdit(c *gin.Context) {
	sysRole := new(systemModels.SysRoleDML)
	_ = c.ShouldBindJSON(sysRole)
	if sysRole.RoleId == 1 {
		baizeContext.Waring(c, "admin角色不能修改")
		return
	}
	if rc.rs.CheckRoleNameUnique(c, sysRole.RoleId, sysRole.RoleName) {
		baizeContext.Waring(c, "修改角色'"+sysRole.RoleName+"'失败，角色名称已存在")
		return
	}
	if rc.rs.CheckRoleKeyUnique(c, sysRole.RoleId, sysRole.RoleKey) {
		baizeContext.Waring(c, "修改角色'"+sysRole.RoleKey+"'失败，角色权限已存在")
		return
	}
	sysRole.SetUpdateBy(baizeContext.GetUserId(c))
	rc.rs.UpdateRole(c, sysRole)
	baizeContext.Success(c)
}

// RoleChangeStatus 修改角色状态
// @Summary 修改角色状态
// @Description 修改角色状态
// @Tags 角色相关
// @Param  object body systemModels.SysRoleDML true "公司信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/role/changeStatus  [put]
func (rc *Role) RoleChangeStatus(c *gin.Context) {
	sysRole := new(systemModels.SysRoleDML)
	_ = c.ShouldBindJSON(sysRole)
	sysRole.SetUpdateBy(baizeContext.GetUserId(c))
	rc.rs.UpdateRoleStatus(c, sysRole)
	baizeContext.Success(c)
}

// RoleRemove 删除角色
// @Summary 删除角色
// @Description 删除角色
// @Tags 角色相关
// @Param rolesIds path []string true "rolesIds"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/role/{rolesIds}  [delete]
func (rc *Role) RoleRemove(c *gin.Context) {
	ids := baizeContext.ParamInt64Array(c, "rolesIds")
	if rc.rs.CountUserRoleByRoleId(c, ids) {
		baizeContext.Waring(c, "角色已分配，不能删除")
		return
	}
	rc.rs.DeleteRoleByIds(c, ids)
	baizeContext.Success(c)
}

// AllocatedList 查询角色授权用户列表查询
// @Summary 查询角色授权用户列表查询
// @Description 查询角色授权用户列表查询
// @Tags 角色相关
// @Param  object query systemModels.SysRoleAndUserDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ListData{Rows=[]systemModels.SysUserVo}}  "成功"
// @Router /system/role/changeStatus  [get]
func (rc *Role) AllocatedList(c *gin.Context) {
	user := new(systemModels.SysRoleAndUserDQL)
	if err := c.ShouldBind(user); err != nil {
		baizeContext.ParameterError(c)
		return
	}
	user.DataScope = baizeContext.GetDataScope(c, "d")
	list, count := rc.rs.SelectAllocatedList(c, user)
	baizeContext.SuccessListData(c, list, count)

}

// UnallocatedList 查询角色未授权用户列表查询
// @Summary 查询角色未授权用户列表查询
// @Description 查询角色未授权用户列表查询
// @Tags 角色相关
// @Param  object query systemModels.SysRoleAndUserDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ListData{Rows=[]systemModels.SysUserVo}}  "成功"
// @Router /system/role/authUser/unallocatedList  [put]
func (rc *Role) UnallocatedList(c *gin.Context) {
	user := new(systemModels.SysRoleAndUserDQL)
	if err := c.ShouldBind(user); err != nil {
		baizeContext.ParameterError(c)
		return
	}
	user.DataScope = baizeContext.GetDataScope(c, "d")
	list, count := rc.rs.SelectUnallocatedList(c, user)
	baizeContext.SuccessListData(c, list, count)
}

// InsertAuthUser 角色授权用户
// @Summary 角色授权用户
// @Description 角色授权用户
// @Tags 角色相关
// @Param  roleId query string true "角色Id"
// @Param  userIds query []string true "用户Ids"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/role/authUser/selectAll  [put]
func (rc *Role) InsertAuthUser(c *gin.Context) {
	rc.rs.InsertAuthUsers(c, baizeContext.QueryInt64(c, "roleId"), baizeContext.QueryInt64Array(c, "userIds"))
	baizeContext.Success(c)
}

// CancelAuthUser 取消用户角色
// @Summary 取消用户角色
// @Description 取消用户角色
// @Tags 角色相关
// @Param  object body systemModels.SysUserRole  true  "用户id角色id"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/role/authUser/cancel  [put]
func (rc *Role) CancelAuthUser(c *gin.Context) {
	userRole := new(systemModels.SysUserRole)
	if err := c.ShouldBindJSON(userRole); err != nil {
		baizeContext.ParameterError(c)
		return
	}
	rc.rs.DeleteAuthUserRole(c, userRole)
	baizeContext.Success(c)
}

// CancelAuthUserAll 取消角色授权用户
// @Summary 取消角色授权用户
// @Description 取消角色授权用户
// @Tags 角色相关
// @Param  roleId query string true "角色Id"
// @Param  userIds query []string true "用户Ids"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/role/authUser/cancelAll  [put]
func (rc *Role) CancelAuthUserAll(c *gin.Context) {
	rc.rs.DeleteAuthUsers(c, baizeContext.QueryInt64(c, "roleId"), baizeContext.QueryInt64Array(c, "userIds"))
	baizeContext.Success(c)
}
