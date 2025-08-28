package systemController

import (
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/middlewares"
	"baize/app/utils/baizeContext"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Permission struct {
	ps systemService.ISysPermissionService
}

func NewPermission(ps systemService.ISysPermissionService) *Permission {
	return &Permission{ps: ps}
}

func (pc *Permission) PrivateRoutes(router *gin.RouterGroup) {
	systemPost := router.Group("/system/permission")
	systemPost.GET("/list", middlewares.HasPermission("system:permission"), pc.PermissionList)
	systemPost.GET("/:permissionId", middlewares.HasPermission("system:permission:query"), pc.PermissionGetInfo)
	systemPost.POST("", middlewares.SetLog("权限管理", middlewares.Insert), middlewares.HasPermission("system:permission:add"), pc.PermissionAdd)
	systemPost.PUT("", middlewares.SetLog("权限管理", middlewares.Update), middlewares.HasPermission("system:permission:edit"), pc.PermissionEdit)
	systemPost.DELETE("/:permissionId", middlewares.SetLog("权限管理", middlewares.Delete), middlewares.HasPermission("system:permission:remove"), pc.PermissionRemove)

}

// PermissionList 系统权限列表
// @Summary 系统权限列表
// @Description 系统权限列表
// @Tags 系统权限
// @Param object query systemModels.SysPermissionDQL false "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=[]systemModels.SysPermissionVo} "成功"
// @Router /system/permission/list [get]
func (pc *Permission) PermissionList(c *gin.Context) {
	permission := new(systemModels.SysPermissionDQL)
	_ = c.ShouldBind(permission)
	list := pc.ps.SelectPermissionList(c, permission)
	baizeContext.SuccessData(c, list)
}

// PermissionGetInfo 根据id查询系统权限
// @Summary 根据id查询系统权限
// @Description 根据id查询系统权限
// @Tags 系统权限
// @Param permissionId path string true "permissionId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=systemModels.SysPermissionVo} "成功"
// @Router /system/permission/{permissionId}  [get]
func (pc *Permission) PermissionGetInfo(c *gin.Context) {
	permissionId := c.Param("permissionId")
	permission := pc.ps.SelectPermissionById(c, permissionId)
	baizeContext.SuccessData(c, permission)
}

// PermissionListByRoleIds 根据RoleIds查询系统权限
// @Summary 根据id查询系统权限
// @Description 根据id查询系统权限
// @Tags 系统权限
// @Param roleIds path string true "roleIds"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=[]systemModels.SysPermissionVo} "成功"
// @Router /system/permission/byRoleIds/{roleIds} [get]
func (pc *Permission) PermissionListByRoleIds(c *gin.Context) {
	roleIds := strings.Split(c.Param("roleIds"), ",")
	permission := pc.ps.SelectPermissionListByRoleIds(c, roleIds)
	baizeContext.SuccessData(c, permission)
}

// PermissionAdd 新增系统权限
// @Summary 新增系统权限
// @Description 新增系统权限
// @Tags 系统权限
// @Param  object body systemModels.SysPermissionAdd true "系统权限"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object} response.ResponseData
// @Router /system/permission [post]
func (pc *Permission) PermissionAdd(c *gin.Context) {
	permission := new(systemModels.SysPermissionAdd)
	err := c.ShouldBind(permission)
	if err != nil {
		zap.L().Debug("参数错误", zap.Any("err", err))
		baizeContext.ParameterError(c)
		return
	}
	permission.SetCreateBy(baizeContext.GetUserId(c))
	pc.ps.InsertPermission(c, permission)
	baizeContext.Success(c)
}

// PermissionEdit 修改系统权限
// @Summary 修改系统权限
// @Description 修改系统权限
// @Tags 系统权限
// @Param  object body systemModels.SysPermissionEdit true "系统权限"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object} response.ResponseData
// @Router /system/permission [put]
func (pc *Permission) PermissionEdit(c *gin.Context) {
	permission := new(systemModels.SysPermissionEdit)
	_ = c.ShouldBind(permission)
	permission.SetUpdateBy(baizeContext.GetUserId(c))
	pc.ps.UpdatePermission(c, permission)
	baizeContext.Success(c)
}

// PermissionRemove 删除系统权限
// @Summary 删除系统权限
// @Description 删除系统权限
// @Tags 系统权限
// @Param permissionId path string true "permissionId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object} response.ResponseData
// @Router /system/permission/{permissionId} [delete]
func (pc *Permission) PermissionRemove(c *gin.Context) {

	permissionId := c.Param("permissionId")
	if pc.ps.HasChildByPermissionId(c, permissionId) {
		baizeContext.Waring(c, "有子权限不可以删除")
		return
	}
	pc.ps.DeletePermissionById(c, permissionId)
	baizeContext.Success(c)
}
