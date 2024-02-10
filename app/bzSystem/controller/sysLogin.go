package controller

import (
	monitorModels "baize/app/bzMonitor/models"
	"baize/app/bzSystem/models"
	"baize/app/bzSystem/service"
	"baize/app/bzSystem/service/serviceImpl"
	"baize/app/constant/userStatus"
	"baize/app/utils/bCryptPasswordEncoder"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LoginController struct {
	ls service.ILoginService
	us service.IUserService
	ms service.IMenuService
}

func NewLoginController(ls *serviceImpl.LoginService, us *serviceImpl.UserService, ms *serviceImpl.MenuService) *LoginController {
	return &LoginController{ls: ls, us: us, ms: ms}
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录
// @Tags 登录
// @Param  object body models.LoginBody true "登录信息"
// @Produce application/json
// @Success 200 {object}  response.ResponseData "登录成功"
// @Failure 412 {object}  response.ResponseData "参数错误"
// @Failure 500 {object}  response.ResponseData "服务器错误"
// @Failure 600 {object}  response.ResponseData "用户名密码错误"
// @Router /login [post]
func (lc *LoginController) Login(c *gin.Context) {
	var login models.LoginBody
	if err := c.ShouldBindJSON(&login); err != nil {
		zap.L().Debug("参数错误", zap.Error(err))
		baizeContext.ParameterError(c)
		return
	}
	logininfor := new(monitorModels.Logininfor)
	logininfor.UserName = login.Username

	baizeContext.SetUserAgent(c, logininfor)
	captcha := lc.ls.VerityCaptcha(c, login.Uuid, login.Code)
	if !captcha {
		logininfor.Status = 1
		logininfor.Msg = "验证码错误"

		return
	}
	user := lc.us.SelectUserByUserName(c, login.Username)
	if user == nil {
		logininfor.Status = 1
		logininfor.Msg = login.Username + " 用户不存在"
		baizeContext.Waring(c, "用户不存在/密码错误")
		return
	} else if userStatus.Deleted == user.DelFlag {
		logininfor.Status = 1
		logininfor.Msg = login.Username + " 已被删除"
		baizeContext.Waring(c, "对不起，您的账号："+login.Username+" 已被删除")
		return
	} else if userStatus.Disable == user.Status {
		logininfor.Status = 1
		logininfor.Msg = login.Username + " 已停用"
		baizeContext.Waring(c, "对不起，您的账号："+login.Username+" 已停用")
		return
	} else if !bCryptPasswordEncoder.CheckPasswordHash(login.Password, user.Password) {
		logininfor.Status = 1
		logininfor.Msg = login.Username + "密码错误"
		baizeContext.Waring(c, "用户不存在/密码错误")
		return
	}

	baizeContext.SuccessData(c, lc.ls.Login(c, user, logininfor))
}

// GetInfo 获取用户个人信息
// @Summary 获取用户个人信息
// @Description 获取用户个人信息
// @Tags 登录
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=models.GetInfo}  "获取成功"
// @Router /getInfo [get]
func (lc *LoginController) GetInfo(c *gin.Context) {

	baizeContext.SuccessData(c, lc.ls.GetInfo(c))

}

// Logout 退出
// @Summary 退出
// @Description 退出
// @Tags 登录
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "退出成功"
// @Router /logout [post]
func (lc *LoginController) Logout(c *gin.Context) {
	//bzc := baizeContext.NewBaiZeContext(c)
	//loginUser := bzc.GetCurrentUser()
	//if loginUser != nil {
	//	lc.ls.ForceLogout(loginUser.Token)
	//}
	//bzc.Success()
}

// GetCode 获取验证码
// @Summary 获取验证码
// @Description 获取验证码
// @Tags 登录
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "获取成功"
// @Router /captchaImage [get]
func (lc *LoginController) GetCode(c *gin.Context) {
	baizeContext.SuccessData(c, lc.ls.GenerateCode(c))
}
func (lc *LoginController) GetRouters(c *gin.Context) {
	userId := baizeContext.GetUserId(c)
	menus := lc.ms.SelectMenuTreeByUserId(c, userId)
	buildMenus := lc.ms.BuildMenus(c, menus)
	baizeContext.SuccessData(c, buildMenus)

}
