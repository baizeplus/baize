package controller

import (
	"github.com/gin-gonic/gin"
)

type PermissionController struct {
}

func NewPermissionController() *PermissionController {
	return &PermissionController{}
}

// PermissionList 查询权限列表查询
// @Summary 查询权限列表查询
// @Description 查询权限列表查询
// @Tags 权限相关
// @Param  object query models.SysDeptDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ResponseData{Rows=[]models.SysPermissionVo}}  "成功"
// @Router /bzSystem/permission/list  [get]
func (mc *PermissionController) PermissionList(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//Permission := new(models.SysPermissionDQL)
	//_ = c.ShouldBind(Permission)
	//list := mc.ms.SelectPermissionList(Permission, bzc.GetUserId())
	//bzc.SuccessData(list)
}

// PermissionGetInfo 根据权限ID获取权限信息
// @Summary 根据权限ID获取权限信息
// @Description 根据权限ID获取权限信息
// @Tags 权限相关
// @Param id path string true "permissionId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=models.SysPermissionVo}  "成功"
// @Router /bzSystem/permission/{permissionId}  [get]
func (mc *PermissionController) PermissionGetInfo(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//PermissionId := bzc.ParamInt64("permissionId")
	//if PermissionId == 0 {
	//	bzc.ParameterError()
	//	return
	//}
	//Permission := mc.ms.SelectPermissionById(PermissionId)
	//bzc.SuccessData(Permission)
}
func (mc *PermissionController) PermissionTreeSelect(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//userId := bzc.GetUserId()
	//bzc.SuccessData(mc.ms.SelectPermissionList(new(models.SysPermissionDQL), userId))
}

// PermissionAdd 添加权限
// @Summary 添加权限
// @Description 添加权限
// @Tags 权限相关
// @Param  object body models.SysPermissionVo true "公司信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /bzSystem/permission  [post]
func (mc *PermissionController) PermissionAdd(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//sysPermission := new(models.SysPermissionDML)
	//c.ShouldBind(sysPermission)
	//if mc.ms.CheckPermissionNameUnique(sysPermission) {
	//	bzc.Waring("新增菜单'" + sysPermission.PermissionName + "'失败，菜单名称已存在")
	//	return
	//}
	//sysPermission.SetCreateBy(bzc.GetUserId())
	//mc.ms.InsertPermission(sysPermission)
	//bzc.Success()
}

// PermissionEdit 修改权限
// @Summary 修改权限
// @Description 修改权限
// @Tags 权限相关
// @Param  object body models.SysPermissionVo true "公司信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /bzSystem/permission  [put]
func (mc *PermissionController) PermissionEdit(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//sysPermission := new(models.SysPermissionDML)
	//if mc.ms.CheckPermissionNameUnique(sysPermission) {
	//	bzc.Waring("修改菜单'" + sysPermission.PermissionName + "'失败，菜单名称已存在")
	//	return
	//}
	//c.ShouldBind(sysPermission)
	//sysPermission.SetCreateBy(bzc.GetUserId())
	//mc.ms.UpdatePermission(sysPermission)
	//bzc.Success()
}

// PermissionRemove 删除权限
// @Summary 删除权限
// @Description 删除权限
// @Tags 权限相关
// @Param permissionId path string true "permissionId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /bzSystem/permission/{permissionId}  [delete]
func (mc *PermissionController) PermissionRemove(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//PermissionId := bzc.ParamInt64("permissionId")
	//if PermissionId == 0 {
	//	zap.L().Error("参数错误")
	//	bzc.ParameterError()
	//	return
	//}
	//if mc.ms.HasChildByPermissionId(PermissionId) {
	//	bzc.Waring("存在子菜单,不允许删除")
	//	return
	//}
	//if mc.ms.CheckPermissionExistRole(PermissionId) {
	//	bzc.Waring("菜单已分配,不允许删除")
	//	return
	//}
	//mc.ms.DeletePermissionById(PermissionId)
	//bzc.Success()
}
func (mc *PermissionController) RolePermissionTreeselect(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//roleId := bzc.ParamInt64("roleId")
	//if roleId == 0 {
	//	zap.L().Error("参数错误")
	//	bzc.ParameterError()
	//}
	//userId := bzc.GetUserId()
	//m := make(map[string]interface{})
	//m["checkedKeys"] = mc.ms.SelectPermissionListByRoleId(roleId)
	//m["permissions"] = mc.ms.SelectPermissionList(new(models.SysPermissionDQL), userId)
	//bzc.SuccessData(m)
}
