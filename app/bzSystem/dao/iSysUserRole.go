package systemDao

import (
	"baize/app/bzSystem/models"
	"context"
	"github.com/baizeplus/sqly"
)

type IUserRoleDao interface {
	DeleteUserRole(ctx context.Context, db sqly.SqlyContext, ids []int64)
	BatchUserRole(ctx context.Context, db sqly.SqlyContext, users []*models.SysUserRole)
	DeleteUserRoleByUserId(ctx context.Context, db sqly.SqlyContext, userId int64)
	CountUserRoleByRoleId(ctx context.Context, db sqly.SqlyContext, ids []int64) int
	DeleteUserRoleInfo(ctx context.Context, db sqly.SqlyContext, userRole *models.SysUserRole)
	DeleteUserRoleInfos(ctx context.Context, db sqly.SqlyContext, roleId int64, userIds []int64)
}
