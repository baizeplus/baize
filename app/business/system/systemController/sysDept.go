package systemController

import (
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/business/system/systemService/systemServiceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Dept struct {
	ds systemService.IDeptService
}

func NewDept(ds *systemServiceImpl.DeptService) *Dept {
	return &Dept{ds: ds}
}

// DeptList 查询部门列表查询
// @Summary 查询部门列表查询
// @Description 查询部门列表查询
// @Tags 部门相关
// @Param  object query systemModels.SysDeptDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ListData{Rows=[]systemModels.SysDeptVo}}  "成功"
// @Router /system/dept  [get]
func (dc *Dept) DeptList(c *gin.Context) {
	dept := new(systemModels.SysDeptDQL)
	_ = c.ShouldBind(dept)
	dept.DataScope = baizeContext.GetDataScope(c, "d")
	baizeContext.SuccessData(c, dc.ds.SelectDeptList(c, dept))

}

// DeptGetInfo 根据部门ID获取部门信息
// @Summary 根据部门ID获取部门信息
// @Description 根据部门ID获取部门信息
// @Tags 部门相关
// @Param id path string true "deptId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=systemModels.SysDeptVo}  "成功"
// @Router /system/dept/{deptId}  [get]
func (dc *Dept) DeptGetInfo(c *gin.Context) {
	deptId := baizeContext.ParamInt64(c, "deptId")
	if deptId == 0 {
		zap.L().Debug("参数错误")
		baizeContext.ParameterError(c)
		return
	}
	baizeContext.SuccessData(c, dc.ds.SelectDeptById(c, deptId))
}

// RoleDeptTreeSelect 获取角色部门
// @Summary 获取角色部门
// @Description 获取角色部门
// @Tags 部门相关
// @Param id path string true "roleId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=systemModels.RoleDeptTree}  "成功"
// @Router /system/dept/roleDeptTreeSelect/{roleId}  [get]
func (dc *Dept) RoleDeptTreeSelect(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//roleId := bzc.ParamInt64("roleId")
	//if roleId == 0 {
	//	zap.L().Error("参数错误")
	//	bzc.ParameterError()
	//}
	//rdt := new(systemModels.RoleDeptTree)
	//rdt.CheckedKeys = dc.ds.SelectDeptListByRoleId(roleId)
	//rdt.Depts = dc.ds.SelectDeptList(new(systemModels.SysDeptDQL))
	//bzc.SuccessData(rdt)
}

// DeptAdd 添加部门
// @Summary 添加部门
// @Description 添加部门
// @Tags 部门相关
// @Param  object body systemModels.SysDeptVo true "公司信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/dept  [post]
func (dc *Dept) DeptAdd(c *gin.Context) {
	sysDept := new(systemModels.SysDeptVo)
	_ = c.ShouldBindJSON(sysDept)
	if dc.ds.CheckDeptNameUnique(c, 0, sysDept.ParentId, sysDept.DeptName) {
		baizeContext.Waring(c, "新增部门'"+sysDept.DeptName+"'失败，部门名称已存在")
		return
	}
	sysDept.SetCreateBy(baizeContext.GetUserId(c))
	dc.ds.InsertDept(c, sysDept)
	baizeContext.Success(c)
}

// DeptEdit 修改部门
// @Summary 修改部门
// @Description 修改部门
// @Tags 部门相关
// @Param  object body systemModels.SysDeptVo true "公司信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/dept  [put]
func (dc *Dept) DeptEdit(c *gin.Context) {
	sysDept := new(systemModels.SysDeptVo)
	_ = c.ShouldBindJSON(sysDept)
	if dc.ds.CheckDeptNameUnique(c, sysDept.DeptId, sysDept.ParentId, sysDept.DeptName) {
		baizeContext.Waring(c, "修改部门'"+sysDept.DeptName+"'失败，部门名称已存在")
		return
	}
	sysDept.SetUpdateBy(baizeContext.GetUserId(c))
	dc.ds.UpdateDept(c, sysDept)
	baizeContext.Success(c)
}

// DeptRemove 删除部门
// @Summary 删除部门
// @Description 删除部门
// @Tags 部门相关
// @Param id path string true "deptId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/dept/{deptId}  [delete]
func (dc *Dept) DeptRemove(c *gin.Context) {
	deptId := baizeContext.ParamInt64(c, "deptId")
	if deptId == 0 {
		zap.L().Debug("参数错误")
		baizeContext.ParameterError(c)
		return
	}
	if dc.ds.HasChildByDeptId(c, deptId) {
		baizeContext.Waring(c, "存在下级部门,不允许删除")
		return
	}
	if dc.ds.CheckDeptExistUser(c, deptId) {
		baizeContext.Waring(c, "部门存在用户,不允许删除")
		return
	}
	dc.ds.DeleteDeptById(c, deptId)
	baizeContext.Success(c)
}
