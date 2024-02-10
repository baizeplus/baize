package systemDao

import (
	"baize/app/bzSystem/models"
	"context"
	"github.com/baizeplus/sqly"
)

type IDeptDao interface {
	SelectDeptList(ctx context.Context, db sqly.SqlyContext, dept *models.SysDeptDQL) (sysDeptList []*models.SysDeptVo)
	SelectDeptById(ctx context.Context, db sqly.SqlyContext, deptId int64) (dept *models.SysDeptVo)
	InsertDept(ctx context.Context, db sqly.SqlyContext, dept *models.SysDeptVo)
	UpdateDept(ctx context.Context, db sqly.SqlyContext, dept *models.SysDeptVo)
	DeleteDeptById(ctx context.Context, db sqly.SqlyContext, deptId int64)
	CheckDeptNameUnique(ctx context.Context, db sqly.SqlyContext, deptName string, parentId int64) int64
	HasChildByDeptId(ctx context.Context, db sqly.SqlyContext, deptId int64) int
	CheckDeptExistUser(ctx context.Context, db sqly.SqlyContext, deptId int64) int
	SelectDeptListByRoleId(ctx context.Context, db sqly.SqlyContext, roleId int64, deptCheckStrictly bool) (deptIds []string)
}
