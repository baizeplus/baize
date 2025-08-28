package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IUserDeptScopeDao interface {
	DeleteUserDeptScope(ctx context.Context, ids []string)
	SelectUserDeptScopeDeptIdByUserId(ctx context.Context, id string) []string
	DeleteUserDeptScopeByUserId(ctx context.Context, id string)
	BatchUserDeptScope(ctx context.Context, list []*systemModels.SysUserDeptScope)
}
