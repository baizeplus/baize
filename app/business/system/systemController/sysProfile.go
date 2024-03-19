package systemController

import (
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/business/system/systemService/systemServiceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type Profile struct {
	us systemService.IUserService
}

func NewProfile(us *systemServiceImpl.UserService) *Profile {
	return &Profile{us: us}
}

// Profile 查看个人资料
// @Summary 查看个人资料
// @Description 查看个人资料
// @Tags 个人资料
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=systemModels.UserProfile} "成功"
// @Router /system/user/profile  [get]
func (pc *Profile) Profile(c *gin.Context) {
	baizeContext.SuccessData(c, pc.us.GetUserProfile(c))
}

// ProfileUpdateProfile 修改个人资料
// @Summary 修改个人资料
// @Description 修改个人资料
// @Tags 个人资料
// @Param  object body systemModels.SysUserDML true "公司信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/user/profile  [put]
func (pc *Profile) ProfileUpdateProfile(c *gin.Context) {
	sysUser := new(systemModels.SysUserDML)
	sysUser.UserId = baizeContext.GetUserId(c)
	_ = c.ShouldBindJSON(sysUser)
	if pc.us.CheckPhoneUnique(c, sysUser.UserId, sysUser.Phonenumber) {
		baizeContext.Waring(c, "修改失败'"+sysUser.Phonenumber+"'失败，手机号码已存在")
		return
	}
	if pc.us.CheckEmailUnique(c, sysUser.UserId, sysUser.Email) {
		baizeContext.Waring(c, "修改失败'"+sysUser.Email+"'失败，邮箱账号已存在")
		return
	}
	sysUser.SetUpdateBy(sysUser.UserId)
	pc.us.UpdateUserProfile(c, sysUser)
	baizeContext.Success(c)
}

// ProfileUpdatePwd 修改密码
// @Summary 修改密码
// @Description 修改密码
// @Tags 个人资料
// @Param  oldPassword query string true "旧密码"
// @Param  newPassword query string true "新密码"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/user/profile/updatePwd  [put]
func (pc *Profile) ProfileUpdatePwd(c *gin.Context) {
	oldPassword := c.Query("oldPassword")
	password := c.Query("newPassword")
	if oldPassword == password {
		baizeContext.Waring(c, "新密码不能与旧密码相同")
		return
	}
	userId := baizeContext.GetUserId(c)
	if !pc.us.MatchesPassword(c, oldPassword, userId) {
		baizeContext.Waring(c, "修改密码失败，旧密码错误")
		return
	}
	pc.us.ResetUserPwd(c, userId, password)
	baizeContext.Success(c)
}

// ProfileAvatar 修改头像
// @Summary 修改头像
// @Description 修改头像
// @Tags 个人资料
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /system/user/profile/avatar  [post]
func (pc *Profile) ProfileAvatar(c *gin.Context) {
	file, err := c.FormFile("avatarfile")
	if err != nil {
		baizeContext.ParameterError(c)
		return
	}
	baizeContext.SuccessData(c, pc.us.UpdateUserAvatar(c, file))
}
