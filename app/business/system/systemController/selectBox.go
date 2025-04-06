package systemController

import (
	"baize/app/baize"
	"baize/app/business/system/systemService"
	"baize/app/middlewares"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type SelectBox struct {
	sbs systemService.ISelectBoxService
}

func NewSelectBox(sbs systemService.ISelectBoxService) *SelectBox {
	return &SelectBox{sbs: sbs}
}
func (s *SelectBox) PrivateRoutes(router *gin.RouterGroup) {
	sb := router.Group("/system/selectBox")
	sb.GET("/permission", middlewares.HasPermissions([]string{"system:permission:add", "system:permission:edit", "system:role:add", "system:role:edit"}), s.SelectPermission)
	sb.GET("/dept", middlewares.HasPermissions([]string{"system:user"}), s.SelectDept)
}

// SelectPermission 权限选择框
// @Summary 权限选择框
// @Description 权限选择框
// @Tags 下拉框选项
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=systemModels.SelectPermission}  "成功"
// @Router /system/selectBox/permission  [get]
func (s *SelectBox) SelectPermission(c *gin.Context) {
	baizeContext.SuccessData(c, s.sbs.SelectPermissionBox(c))
}

// SelectDept 部门选择框
// @Summary 部门选择框
// @Description 部门选择框
// @Tags 下拉框选项
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=systemModels.SelectPermission}  "成功"
// @Router /system/selectBox/dept  [get]
func (s *SelectBox) SelectDept(c *gin.Context) {
	be := new(baize.BaseEntityDQL)
	be.DataScope = baizeContext.GetDataScope(c, "d")
	baizeContext.SuccessData(c, s.sbs.SelectDeptBox(c, be))
}
