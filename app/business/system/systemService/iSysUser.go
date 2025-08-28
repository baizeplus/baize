package systemService

import (
	"baize/app/business/system/systemModels"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type IUserService interface {
	SelectUserByUserName(c *gin.Context, userName string) *systemModels.User
	SelectUserList(c *gin.Context, user *systemModels.SysUserDQL) (sysUserList []*systemModels.SysUserVo, total int64)
	UserExport(c *gin.Context, user *systemModels.SysUserDQL) (data []byte)
	InsertUser(c *gin.Context, sysUser *systemModels.SysUserDML)
	UpdateUser(c *gin.Context, sysUser *systemModels.SysUserDML)

	UpdateUserDataScope(c *gin.Context, uds *systemModels.SysUserDataScope)
	SelectUserDataScope(c *gin.Context, userId string) *systemModels.SysUserDataScope

	UpdateUserStatus(c *gin.Context, sysUser *systemModels.EditUserStatus)
	ResetPwd(c *gin.Context, userId string, password string)
	CheckUserNameUnique(c *gin.Context, userName string) bool
	CheckPhoneUnique(c *gin.Context, id string, phoneNumber string) bool
	CheckEmailUnique(c *gin.Context, id string, email string) bool
	DeleteUserByIds(c *gin.Context, ids []string)
	UserImportData(c *gin.Context, file *multipart.FileHeader) (msg string, failureNum int)
	UpdateUserAvatar(c *gin.Context, file *multipart.FileHeader) string
	ResetUserPwd(c *gin.Context, userId string, password string)
	UpdateUserProfile(c *gin.Context, sysUser *systemModels.SysUserDML)
	MatchesPassword(c *gin.Context, rawPassword string, userId string) bool
	InsertUserAuth(c *gin.Context, userId string, roleIds []string)
	GetUserAuthRole(c *gin.Context, userId string) *systemModels.UserAndRoles
	SelectUserAndAccreditById(c *gin.Context, userId string) (sysUser *systemModels.UserAndAccredit)
	SelectAccredit(c *gin.Context) (sysUser *systemModels.Accredit)
	ImportTemplate(c *gin.Context) (data []byte)
	GetUserProfile(c *gin.Context) *systemModels.UserProfile
}
