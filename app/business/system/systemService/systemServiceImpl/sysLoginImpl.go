package systemServiceImpl

import (
	"baize/app/business/monitor/monitorDao"
	"baize/app/business/monitor/monitorDao/monitorDaoImpl"
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemDao/systemDaoImpl"
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/constant/dataScopeAspect"
	"baize/app/utils/bCryptPasswordEncoder"
	"baize/app/utils/ipUtils"
	"encoding/json"

	"baize/app/constant/sessionStatus"
	"baize/app/utils/baizeContext"
	"baize/app/utils/session"
	"context"

	"baize/app/utils/snowflake"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"image/color"
	"time"
)

type LoginService struct {
	data        *sqly.DB
	userDao     systemDao.IUserDao
	menuDao     systemDao.IMenuDao
	roleDao     systemDao.IRoleDao
	loginforDao monitorDao.ILogininforDao
	driver      *base64Captcha.DriverMath
	store       base64Captcha.Store
	cs          systemService.IConfigService
}

func NewLoginService(data *sqly.DB, ud *systemDaoImpl.SysUserDao, md *systemDaoImpl.SysMenuDao, rd *systemDaoImpl.SysRoleDao, ld *monitorDaoImpl.LogininforDao, cs *ConfigService) *LoginService {
	return &LoginService{data: data, userDao: ud, menuDao: md, roleDao: rd, loginforDao: ld, cs: cs,
		driver: base64Captcha.NewDriverMath(38, 106, 0, 0, &color.RGBA{0, 0, 0, 0}, nil, []string{"wqy-microhei.ttc"}),
		store:  base64Captcha.DefaultMemStore,
	}
}

func (loginService *LoginService) Login(c *gin.Context, user *systemModels.User, l *monitorModels.Logininfor) string {
	l.Status = 0
	l.Msg = "登录成功"
	manager := session.NewManger()
	session, _ := manager.InitSession(c, user.UserId)
	roles := loginService.roleDao.SelectBasicRolesByUserId(c, loginService.data, user.UserId)
	byRoles, loginRoles := loginService.RolePermissionByRoles(roles)
	rb, _ := json.Marshal(byRoles)
	session.Set(c, sessionStatus.Role, loginRoles)
	session.Set(c, sessionStatus.RolePerms, rb)
	permission := loginService.getPermission(c, user.UserId)
	session.Set(c, sessionStatus.Permission, permission)
	session.Set(c, sessionStatus.IpAddr, c.ClientIP())
	session.Set(c, sessionStatus.LoginTime, time.Now().Unix())
	session.Set(c, sessionStatus.Os, l.Os)
	session.Set(c, sessionStatus.Browser, l.Browser)
	session.Set(c, sessionStatus.UserName, user.UserName)
	session.Set(c, sessionStatus.Avatar, user.Avatar)
	session.Set(c, sessionStatus.DeptId, user.DeptId)
	session.Set(c, sessionStatus.DataScopeAspect, user.DataScope)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				zap.L().Error("登录日志记录错误", zap.Any("error", err))
			}
		}()
		l.LoginLocation = ipUtils.GetRealAddressByIP(l.IpAddr)
		loginService.loginforDao.InserLogininfor(context.Background(), loginService.data, l)
	}()
	return session.Id()
}

func (loginService *LoginService) Register(c *gin.Context, user *systemModels.LoginBody) {
	u := new(systemModels.SysUserDML)
	u.Password = bCryptPasswordEncoder.HashPassword(user.Password)
	u.DataScope = dataScopeAspect.NoDataScope
	u.UserId = snowflake.GenID()
	u.NickName = user.Username
	u.UserName = user.Username
	u.Status = "0"
	u.DeptId = 100
	u.SetCreateBy(u.UserId)
	loginService.userDao.InsertUser(c, loginService.data, u)
}

func (loginService *LoginService) RecordLoginInfo(c *gin.Context, loginUser *monitorModels.Logininfor) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				zap.L().Error("登录日志记录错误", zap.Any("error", err))
			}
		}()
		loginUser.InfoId = snowflake.GenID()
		loginService.loginforDao.InserLogininfor(c, loginService.data, loginUser)
	}()

}

func (loginService *LoginService) getPermission(c *gin.Context, userId int64) []string {
	perms := make([]string, 0)
	if baizeContext.IsAdmin(c) {
		perms = loginService.menuDao.SelectMenuPermsAll(c, loginService.data)
	} else {
		perms = loginService.menuDao.SelectMenuPermsByUserId(c, loginService.data, userId)
	}
	return perms
}

func (loginService *LoginService) GenerateCode(c *gin.Context) (m *systemModels.CaptchaVo) {
	m = new(systemModels.CaptchaVo)
	key := loginService.cs.SelectConfigValueByKey(c, "sys.account.captchaEnabled")
	if key != "false" {
		captcha := base64Captcha.NewCaptcha(loginService.driver, loginService.store)
		id, b64s, _, err := captcha.Generate()
		if err != nil {
			panic(err)
		}
		m.Id = id
		m.Img = b64s
		m.CaptchaEnabled = true
	}
	key = loginService.cs.SelectConfigValueByKey(c, "sys.account.registerUser")
	if key == "true" {
		m.RegisterEnabled = true
	}
	return m
}

func (loginService *LoginService) VerityCaptcha(c *gin.Context, id, base64 string) bool {
	return loginService.store.Verify(id, base64, true)
}

func (loginService *LoginService) ForceLogout(c *gin.Context, token string) {
	panic("等待补充")
}

func (loginService *LoginService) RolePermissionByRoles(roles []*systemModels.SysRole) (rolePerms []string, loginRoles []int64) {
	loginRoles = make([]int64, 0, len(roles))
	rolePerms = make([]string, 0, len(roles))
	for _, role := range roles {
		rolePerms = append(rolePerms, role.RoleKey)
		loginRoles = append(loginRoles, role.RoleId)
	}
	return
}
func (loginService *LoginService) GetInfo(c *gin.Context) *systemModels.GetInfo {
	getInfo := new(systemModels.GetInfo)
	u := new(systemModels.User)
	u.UserId = baizeContext.GetUserId(c)
	u.UserName = baizeContext.GetUserName(c)
	u.Avatar = baizeContext.GetAvatar(c)
	getInfo.User = u
	getInfo.Roles = baizeContext.GetRolesPerms(c)
	getInfo.Permissions = baizeContext.GetPermission(c)
	return getInfo
}
