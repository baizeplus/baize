package systemService

import (
	"baize/app/business/system/systemModels"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type IUserService interface {
	SelectUserByUserName(c *gin.Context, userName string) *systemModels.User
	SelectUserList(c *gin.Context, user *systemModels.SysUserDQL) (sysUserList []*systemModels.SysUserVo, count *int64)
	UserExport(c *gin.Context, user *systemModels.SysUserDQL) (data []byte)
	SelectUserById(c *gin.Context, userId int64) (sysUser *systemModels.SysUserVo)
	InsertUser(c *gin.Context, sysUser *systemModels.SysUserDML)
	UpdateUser(c *gin.Context, sysUser *systemModels.SysUserDML)
	UpdateUserStatus(c *gin.Context, sysUser *systemModels.EditUserStatus)
	ResetPwd(c *gin.Context, userId int64, password string)
	CheckUserNameUnique(c *gin.Context, userName string) bool
	CheckPhoneUnique(c *gin.Context, id int64, phonenumber string) bool
	CheckEmailUnique(c *gin.Context, id int64, email string) bool
	DeleteUserByIds(c *gin.Context, ids []int64)
	//UserImportData(rows [][]string, userId int64, deptId int64) (msg string, failureNum int)
	UpdateLoginInformation(c *gin.Context, userId int64, ip string)
	UpdateUserAvatar(c *gin.Context, file *multipart.FileHeader) string
	ResetUserPwd(c *gin.Context, userId int64, password string)
	UpdateUserProfile(c *gin.Context, sysUser *systemModels.SysUserDML)
	MatchesPassword(c *gin.Context, rawPassword string, userId int64) bool
	InsertUserAuth(c *gin.Context, userId int64, roleIds []int64)
	GetUserAuthRole(c *gin.Context, userId int64) *systemModels.UserAndRoles
	SelectUserAndAccreditById(c *gin.Context, userId int64) (sysUser *systemModels.UserAndAccredit)
	SelectAccredit(c *gin.Context) (sysUser *systemModels.Accredit)
	//ImportTemplate() (data []byte)
}