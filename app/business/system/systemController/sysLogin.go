package systemController

import (
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/business/system/systemService/systemServiceImpl"
	"baize/app/constant/userStatus"
	"baize/app/utils/bCryptPasswordEncoder"
	"baize/app/utils/baizeContext"
	"baize/app/utils/session"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Login struct {
	ls systemService.ILoginService
	us systemService.IUserService
	ms systemService.IMenuService
	cs systemService.IConfigService
}

func NewLogin(ls *systemServiceImpl.LoginService, us *systemServiceImpl.UserService, ms *systemServiceImpl.MenuService, cs *systemServiceImpl.ConfigService) *Login {
	return &Login{ls: ls, us: us, ms: ms, cs: cs}
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录
// @Tags 登录
// @Param  object body systemModels.LoginBody true "登录信息"
// @Success 200 {object}  response.ResponseData "登录成功"
// @Failure 412 {object}  response.ResponseData "参数错误"
// @Failure 500 {object}  response.ResponseData "服务器错误"
// @Failure 600 {object}  response.ResponseData "用户名密码错误"
// @Router /login [post]
func (lc *Login) Login(c *gin.Context) {
	var login systemModels.LoginBody
	if err := c.ShouldBindJSON(&login); err != nil {
		zap.L().Debug("参数错误", zap.Error(err))
		baizeContext.ParameterError(c)
		return
	}
	logininfor := new(monitorModels.Logininfor)
	logininfor.UserName = login.Username
	baizeContext.SetUserAgent(c, logininfor)
	if lc.cs.SelectConfigValueByKey(c, "sys.account.captchaEnabled") != "false" {
		captcha := lc.ls.VerityCaptcha(c, login.Uuid, login.Code)
		if !captcha {
			logininfor.Status = 1
			logininfor.Msg = "验证码错误"
			baizeContext.Waring(c, "验证码错误")
			return
		}
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

// Register 用户登录
// @Summary 用户登录
// @Description 用户登录
// @Tags 登录
// @Param  object body systemModels.LoginBody true "登录信息"
// @Success 200 {object}  response.ResponseData "注册成功"
// @Router /register [post]
func (lc *Login) Register(c *gin.Context) {
	login := new(systemModels.LoginBody)
	if err := c.ShouldBindJSON(login); err != nil {
		zap.L().Debug("参数错误", zap.Error(err))
		baizeContext.ParameterError(c)
		return
	}
	if lc.cs.SelectConfigValueByKey(c, "sys.account.captchaEnabled") != "false" {
		captcha := lc.ls.VerityCaptcha(c, login.Uuid, login.Code)
		if !captcha {
			baizeContext.Waring(c, "验证码错误")
			return
		}
	}
	if lc.us.CheckUserNameUnique(c, login.Username) {
		baizeContext.Waring(c, "登录账号已存在")
		return
	}
	lc.ls.Register(c, login)
	baizeContext.Success(c)

}

// GetInfo 获取用户个人信息
// @Summary 获取用户个人信息
// @Description 获取用户个人信息
// @Tags 登录
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=systemModels.GetInfo}  "获取成功"
// @Router /getInfo [get]
func (lc *Login) GetInfo(c *gin.Context) {

	baizeContext.SuccessData(c, lc.ls.GetInfo(c))

}

// Logout 退出
// @Summary 退出
// @Description 退出
// @Tags 登录
// @Produce application/json
// @Success 200 {object}  response.ResponseData "退出成功"
// @Router /logout [post]
func (lc *Login) Logout(c *gin.Context) {
	manager := session.NewManger()
	manager.RemoveSession(c)
	baizeContext.Success(c)
}

// GetCode 获取验证码
// @Summary 获取验证码
// @Description 获取验证码
// @Tags 登录
// @Produce application/json
// @Success 200 {object}  response.ResponseData "获取成功"
// @Router /captchaImage [get]
func (lc *Login) GetCode(c *gin.Context) {
	baizeContext.SuccessData(c, lc.ls.GenerateCode(c))
}

// GetRouters 获取路由
// @Summary 获取路由
// @Description 获取路由
// @Tags 登录
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ListData{Rows=[]systemModels.RouterVo}} "获取成功"
// @Router /getRouters [get]
func (lc *Login) GetRouters(c *gin.Context) {
	userId := baizeContext.GetUserId(c)
	menus := lc.ms.SelectMenuTreeByUserId(c, userId)
	buildMenus := lc.ms.BuildMenus(c, menus)
	baizeContext.SuccessData(c, buildMenus)
}
