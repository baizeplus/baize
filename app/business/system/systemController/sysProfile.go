package systemController

import (
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/business/system/systemService/systemServiceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type Profile struct {
	rs systemService.IRoleService
	ps systemService.IPostService
	us systemService.IUserService
}

func NewProfile(rs *systemServiceImpl.RoleService, ps *systemServiceImpl.PostService, us *systemServiceImpl.UserService) *Profile {
	return &Profile{rs: rs, ps: ps, us: us}
}

func (pc *Profile) Profile(c *gin.Context) {

	m := make(map[string]interface{})
	m["user"] = pc.us.SelectUserById(c, baizeContext.GetUserId(c))
	m["roleGroup"] = pc.rs.SelectUserRoleGroupByUserId(c, baizeContext.GetUserId(c))
	m["postGroup"] = pc.ps.SelectUserPostGroupByUserId(c, baizeContext.GetUserId(c))
	baizeContext.SuccessData(c, m)
}

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

func (pc *Profile) ProfileAvatar(c *gin.Context) {
	file, err := c.FormFile("avatarfile")
	if err != nil {
		baizeContext.ParameterError(c)
		return
	}
	baizeContext.SuccessData(c, pc.us.UpdateUserAvatar(c, file))
}
