package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IUserRoleDao interface {
	DeleteUserRole(ctx context.Context, ids []int64)
	BatchUserRole(ctx context.Context, users []*systemModels.SysUserRole)
	DeleteUserRoleByUserId(ctx context.Context, userId int64)
	CountUserRoleByRoleId(ctx context.Context, ids []int64) int
	DeleteUserRoleInfo(ctx context.Context, userRole *systemModels.SysUserRole)
	DeleteUserRoleInfos(ctx context.Context, roleId int64, userIds []int64)
}
