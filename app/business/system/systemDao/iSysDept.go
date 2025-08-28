package systemDao

import (
	"baize/app/baize"
	"baize/app/business/system/systemModels"
	"context"
)

type IDeptDao interface {
	SelectDeptList(ctx context.Context, dept *systemModels.SysDeptDQL) (sysDeptList []*systemModels.SysDeptVo)
	SelectDeptListSelectBox(ctx context.Context, dept *baize.BaseEntityDQL) (list []*systemModels.SelectDept)
	SelectDeptById(ctx context.Context, deptId string) (dept *systemModels.SysDeptVo)
	InsertDept(ctx context.Context, dept *systemModels.SysDeptVo)
	UpdateDept(ctx context.Context, dept *systemModels.SysDeptVo)
	DeleteDeptById(ctx context.Context, deptId string)
	CheckDeptNameUnique(ctx context.Context, deptName string, parentId string) string
	HasChildByDeptId(ctx context.Context, deptId string) int
	CheckDeptExistUser(ctx context.Context, deptId string) int
}
