package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type IUserDao interface {
	SelectUserNameByUserName(ctx context.Context, db sqly.SqlyContext, userName []string) []string
	CheckUserNameUnique(ctx context.Context, db sqly.SqlyContext, userName string) int
	CheckPhoneUnique(ctx context.Context, db sqly.SqlyContext, phonenumber string) int64
	CheckEmailUnique(ctx context.Context, db sqly.SqlyContext, email string) int64
	InsertUser(ctx context.Context, db sqly.SqlyContext, sysUser *systemModels.SysUserDML)
	BatchInsertUser(ctx context.Context, db sqly.SqlyContext, sysUser []*systemModels.SysUserDML)
	UpdateUser(ctx context.Context, db sqly.SqlyContext, sysUser *systemModels.SysUserDML)
	SelectUserByUserName(ctx context.Context, db sqly.SqlyContext, userName string) (loginUser *systemModels.User)
	SelectUserById(ctx context.Context, db sqly.SqlyContext, userId int64) (sysUser *systemModels.SysUserVo)
	SelectUserList(ctx context.Context, db sqly.SqlyContext, user *systemModels.SysUserDQL) (sysUserList []*systemModels.SysUserVo, total int64)
	SelectUserListAll(ctx context.Context, db sqly.SqlyContext, user *systemModels.SysUserDQL) (list []*systemModels.SysUserVo)
	DeleteUserByIds(ctx context.Context, db sqly.SqlyContext, ids []int64)
	UpdateLoginInformation(ctx context.Context, db sqly.SqlyContext, userId int64, ip string)
	UpdateUserAvatar(ctx context.Context, db sqly.SqlyContext, userId int64, avatar string)
	ResetUserPwd(ctx context.Context, db sqly.SqlyContext, userId int64, password string)
	SelectPasswordByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) string
	SelectUserIdsByDeptIds(ctx context.Context, db sqly.SqlyContext, deptIds []int64) []int64
}
