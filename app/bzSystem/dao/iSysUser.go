package dao

import (
	"baize/app/bzSystem/models"
	"context"
	"github.com/baizeplus/sqly"
)

type IUserDao interface {
	CheckUserNameUnique(ctx context.Context, db sqly.SqlyContext, userName string) int
	CheckPhoneUnique(ctx context.Context, db sqly.SqlyContext, phonenumber string) int64
	CheckEmailUnique(ctx context.Context, db sqly.SqlyContext, email string) int64
	InsertUser(ctx context.Context, db sqly.SqlyContext, sysUser *models.SysUserDML)
	UpdateUser(ctx context.Context, db sqly.SqlyContext, sysUser *models.SysUserDML)
	SelectUserByUserName(ctx context.Context, db sqly.SqlyContext, userName string) (loginUser *models.User)
	SelectUserById(ctx context.Context, db sqly.SqlyContext, userId int64) (sysUser *models.SysUserVo)
	SelectUserList(ctx context.Context, db sqly.SqlyContext, user *models.SysUserDQL) (sysUserList []*models.SysUserVo, total *int64)
	SelectUserListAll(ctx context.Context, db sqly.SqlyContext, user *models.SysUserDQL) (list []*models.SysUserVo)
	DeleteUserByIds(ctx context.Context, db sqly.SqlyContext, ids []int64)
	UpdateLoginInformation(ctx context.Context, db sqly.SqlyContext, userId int64, ip string)
	UpdateUserAvatar(ctx context.Context, db sqly.SqlyContext, userId int64, avatar string)
	ResetUserPwd(ctx context.Context, db sqly.SqlyContext, userId int64, password string)
	SelectPasswordByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) string
}
