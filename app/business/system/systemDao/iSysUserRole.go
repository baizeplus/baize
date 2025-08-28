package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IUserRoleDao interface {
	DeleteUserRole(ctx context.Context, ids []string)
	BatchUserRole(ctx context.Context, users []*systemModels.SysUserRole)
	DeleteUserRoleByUserId(ctx context.Context, userId string)
	CountUserRoleByRoleId(ctx context.Context, ids []string) int
	DeleteUserRoleInfo(ctx context.Context, userRole *systemModels.SysUserRole)
	DeleteUserRoleInfos(ctx context.Context, roleId string, userIds []string)
}
