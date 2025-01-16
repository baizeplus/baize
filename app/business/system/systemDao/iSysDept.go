package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IDeptDao interface {
	SelectDeptList(ctx context.Context, dept *systemModels.SysDeptDQL) (sysDeptList []*systemModels.SysDeptVo)
	SelectDeptById(ctx context.Context, deptId int64) (dept *systemModels.SysDeptVo)
	InsertDept(ctx context.Context, dept *systemModels.SysDeptVo)
	UpdateDept(ctx context.Context, dept *systemModels.SysDeptVo)
	DeleteDeptById(ctx context.Context, deptId int64)
	CheckDeptNameUnique(ctx context.Context, deptName string, parentId int64) int64
	HasChildByDeptId(ctx context.Context, deptId int64) int
	CheckDeptExistUser(ctx context.Context, deptId int64) int
}
