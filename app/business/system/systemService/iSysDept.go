package systemService

import (
	"baize/app/business/system/systemModels"

	"github.com/gin-gonic/gin"
)

type IDeptService interface {
	SelectDeptList(c *gin.Context, dept *systemModels.SysDeptDQL) (list []*systemModels.SysDeptVo)
	SelectDeptById(c *gin.Context, deptId string) (dept *systemModels.SysDeptVo)
	InsertDept(c *gin.Context, dept *systemModels.SysDeptVo)
	UpdateDept(c *gin.Context, dept *systemModels.SysDeptVo)
	DeleteDeptById(c *gin.Context, dept string)
	CheckDeptNameUnique(c *gin.Context, id, parentId string, deptName string) bool
	HasChildByDeptId(c *gin.Context, deptId string) bool
	CheckDeptExistUser(c *gin.Context, deptId string) bool
}
