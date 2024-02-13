package service

import (
	"baize/app/bzSystem/models"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type IUserService interface {
	SelectUserByUserName(c *gin.Context, userName string) *models.User
	SelectUserList(c *gin.Context, user *models.SysUserDQL) (sysUserList []*models.SysUserVo, count *int64)
	//UserExport(user *models.SysUserDQL) (data []byte)
	SelectUserById(c *gin.Context, userId int64) (sysUser *models.SysUserVo)
	InsertUser(c *gin.Context, sysUser *models.SysUserDML)
	UpdateUser(c *gin.Context, sysUser *models.SysUserDML)
	UpdateUserStatus(c *gin.Context, sysUser *models.EditUserStatus)
	ResetPwd(c *gin.Context, userId int64, password string)
	CheckUserNameUnique(c *gin.Context, userName string) bool
	CheckPhoneUnique(c *gin.Context, id int64, phonenumber string) bool
	CheckEmailUnique(c *gin.Context, id int64, email string) bool
	DeleteUserByIds(c *gin.Context, ids []int64)
	//UserImportData(rows [][]string, userId int64, deptId int64) (msg string, failureNum int)
	UpdateLoginInformation(c *gin.Context, userId int64, ip string)
	UpdateUserAvatar(c *gin.Context, file *multipart.FileHeader) string
	ResetUserPwd(c *gin.Context, userId int64, password string)
	UpdateUserProfile(c *gin.Context, sysUser *models.SysUserDML)
	MatchesPassword(c *gin.Context, rawPassword string, userId int64) bool
	InsertUserAuth(c *gin.Context, userId int64, roleIds []int64)
	GetUserAuthRole(c *gin.Context, userId int64) *models.UserAndRoles
	SelectUserAndAccreditById(c *gin.Context, userId int64) (sysUser *models.UserAndAccredit)
	SelectAccredit(c *gin.Context) (sysUser *models.Accredit)
	//ImportTemplate() (data []byte)
}
