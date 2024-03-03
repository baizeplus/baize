package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IDeptDao interface {
	SelectDeptList(ctx context.Context, db sqly.SqlyContext, dept *systemModels.SysDeptDQL) (sysDeptList []*systemModels.SysDeptVo)
	SelectDeptById(ctx context.Context, db sqly.SqlyContext, deptId int64) (dept *systemModels.SysDeptVo)
	InsertDept(ctx context.Context, db sqly.SqlyContext, dept *systemModels.SysDeptVo)
	UpdateDept(ctx context.Context, db sqly.SqlyContext, dept *systemModels.SysDeptVo)
	DeleteDeptById(ctx context.Context, db sqly.SqlyContext, deptId int64)
	CheckDeptNameUnique(ctx context.Context, db sqly.SqlyContext, deptName string, parentId int64) int64
	HasChildByDeptId(ctx context.Context, db sqly.SqlyContext, deptId int64) int
	CheckDeptExistUser(ctx context.Context, db sqly.SqlyContext, deptId int64) int
}
