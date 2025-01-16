package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IUserDeptScopeDao interface {
	DeleteUserDeptScope(ctx context.Context, ids []int64)
	SelectUserDeptScopeDeptIdByUserId(ctx context.Context, id int64) []string
	DeleteUserDeptScopeByUserId(ctx context.Context, id int64)
	BatchUserDeptScope(ctx context.Context, list []*systemModels.SysUserDeptScope)
}
