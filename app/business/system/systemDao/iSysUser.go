package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IUserDao interface {
	SelectUserNameByUserName(ctx context.Context, userName []string) []string
	CheckUserNameUnique(ctx context.Context, userName string) int
	CheckPhoneUnique(ctx context.Context, phonenumber string) string
	CheckEmailUnique(ctx context.Context, email string) string
	InsertUser(ctx context.Context, sysUser *systemModels.SysUserDML)
	BatchInsertUser(ctx context.Context, sysUser []*systemModels.SysUserDML)
	UpdateUser(ctx context.Context, sysUser *systemModels.SysUserDML)
	SelectUserByUserName(ctx context.Context, userName string) (loginUser *systemModels.User)
	SelectUserById(ctx context.Context, userId string) (sysUser *systemModels.SysUserVo)
	SelectUserList(ctx context.Context, user *systemModels.SysUserDQL) (sysUserList []*systemModels.SysUserVo, total int64)
	SelectUserListAll(ctx context.Context, user *systemModels.SysUserDQL) (list []*systemModels.SysUserVo)
	DeleteUserByIds(ctx context.Context, ids []string)
	UpdateUserAvatar(ctx context.Context, userId string, avatar string)
	ResetUserPwd(ctx context.Context, userId string, password string)
	SelectPasswordByUserId(ctx context.Context, userId string) string
	SelectUserIdsByDeptIds(ctx context.Context, deptIds []string) []string
}
