package systemService

import (
	"baize/app/business/system/systemModels"
	"github.com/gin-gonic/gin"
)

type IDeptService interface {
	SelectDeptList(c *gin.Context, dept *systemModels.SysDeptDQL) (list []*systemModels.SysDeptVo)
	SelectDeptById(c *gin.Context, deptId int64) (dept *systemModels.SysDeptVo)
	InsertDept(c *gin.Context, dept *systemModels.SysDeptVo)
	UpdateDept(c *gin.Context, dept *systemModels.SysDeptVo)
	DeleteDeptById(c *gin.Context, dept int64)
	CheckDeptNameUnique(c *gin.Context, id, parentId int64, deptName string) bool
	HasChildByDeptId(c *gin.Context, deptId int64) bool
	CheckDeptExistUser(c *gin.Context, deptId int64) bool
}
