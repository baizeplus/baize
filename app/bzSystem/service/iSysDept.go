package service

import (
	"baize/app/bzSystem/models"
	"github.com/gin-gonic/gin"
)

type IDeptService interface {
	SelectDeptList(c *gin.Context, dept *models.SysDeptDQL) (list []*models.SysDeptVo)
	SelectDeptById(c *gin.Context, deptId int64) (dept *models.SysDeptVo)
	InsertDept(c *gin.Context, dept *models.SysDeptVo)
	UpdateDept(c *gin.Context, dept *models.SysDeptVo)
	DeleteDeptById(c *gin.Context, dept int64)
	CheckDeptNameUnique(c *gin.Context, id, parentId int64, deptName string) bool
	HasChildByDeptId(c *gin.Context, deptId int64) bool
	CheckDeptExistUser(c *gin.Context, deptId int64) bool
	SelectDeptListByRoleId(c *gin.Context, roleId int64) []string
}
