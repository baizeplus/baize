package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IUserDeptScopeDao interface {
	DeleteUserDeptScope(ctx context.Context, db sqly.SqlyContext, ids []int64)
	SelectUserDeptScopeDeptIdByUserId(ctx context.Context, db sqly.SqlyContext, id int64) []string
	DeleteUserDeptScopeByUserId(ctx context.Context, db sqly.SqlyContext, id int64)
	BatchUserDeptScope(ctx context.Context, db sqly.SqlyContext, list []*systemModels.SysUserDeptScope)
}
