package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IUserRoleDao interface {
	DeleteUserRole(ctx context.Context, db sqly.SqlyContext, ids []int64)
	BatchUserRole(ctx context.Context, db sqly.SqlyContext, users []*systemModels.SysUserRole)
	DeleteUserRoleByUserId(ctx context.Context, db sqly.SqlyContext, userId int64)
	CountUserRoleByRoleId(ctx context.Context, db sqly.SqlyContext, ids []int64) int
	DeleteUserRoleInfo(ctx context.Context, db sqly.SqlyContext, userRole *systemModels.SysUserRole)
	DeleteUserRoleInfos(ctx context.Context, db sqly.SqlyContext, roleId int64, userIds []int64)
}
