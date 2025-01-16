package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type IUserDao interface {
	SelectUserNameByUserName(ctx context.Context, userName []string) []string
	CheckUserNameUnique(ctx context.Context, userName string) int
	CheckPhoneUnique(ctx context.Context, phonenumber string) int64
	CheckEmailUnique(ctx context.Context, email string) int64
	InsertUser(ctx context.Context, sysUser *systemModels.SysUserDML)
	BatchInsertUser(ctx context.Context, sysUser []*systemModels.SysUserDML)
	UpdateUser(ctx context.Context, sysUser *systemModels.SysUserDML)
	SelectUserByUserName(ctx context.Context, userName string) (loginUser *systemModels.User)
	SelectUserById(ctx context.Context, userId int64) (sysUser *systemModels.SysUserVo)
	SelectUserList(ctx context.Context, user *systemModels.SysUserDQL) (sysUserList []*systemModels.SysUserVo, total int64)
	SelectUserListAll(ctx context.Context, user *systemModels.SysUserDQL) (list []*systemModels.SysUserVo)
	DeleteUserByIds(ctx context.Context, ids []int64)
	UpdateLoginInformation(ctx context.Context, userId int64, ip string)
	UpdateUserAvatar(ctx context.Context, userId int64, avatar string)
	ResetUserPwd(ctx context.Context, userId int64, password string)
	SelectPasswordByUserId(ctx context.Context, userId int64) string
	SelectUserIdsByDeptIds(ctx context.Context, deptIds []int64) []int64
}
