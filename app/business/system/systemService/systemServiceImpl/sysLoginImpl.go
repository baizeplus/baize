package systemServiceImpl

import (
	"baize/app/baize"
	"baize/app/business/monitor/monitorDao"
	"baize/app/business/monitor/monitorDao/monitorDaoImpl"
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemDao/systemDaoImpl"
	"baize/app/business/system/systemModels"
	"baize/app/utils/ipUtils"

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
}

func NewLoginService(data *sqly.DB, ud *systemDaoImpl.SysUserDao, md *systemDaoImpl.SysMenuDao, rd *systemDaoImpl.SysRoleDao, ld *monitorDaoImpl.LogininforDao) *LoginService {
	return &LoginService{data: data, userDao: ud, menuDao: md, roleDao: rd, loginforDao: ld,
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
	session.Set(c, sessionStatus.Role, loginRoles)
	session.Set(c, sessionStatus.RolePerms, byRoles)
	permission := loginService.getPermissionPermission(c, user.UserId)
	session.Set(c, sessionStatus.Permission, permission)
	session.Set(c, sessionStatus.IpAddr, c.ClientIP())
	session.Set(c, sessionStatus.LoginTime, time.Now().Unix())
	session.Set(c, sessionStatus.Os, l.Os)
	session.Set(c, sessionStatus.Browser, l.Browser)
	session.Set(c, sessionStatus.UserName, user.UserName)
	session.Set(c, sessionStatus.Avatar, user.Avatar)
	session.Set(c, sessionStatus.DeptId, user.DeptId)
	go func() {
		l.LoginLocation = ipUtils.GetRealAddressByIP(l.IpAddr)
		loginService.loginforDao.InserLogininfor(context.Background(), loginService.data, l)
	}()
	return session.Id()
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

func (loginService *LoginService) getPermissionPermission(c *gin.Context, userId int64) []string {
	perms := make([]string, 0)
	if baizeContext.IsAdmin(c) {
		//perms = append(perms, "*:*:*")
		perms = loginService.menuDao.SelectMenuPermsAll(c, loginService.data)
	} else {
		perms = loginService.menuDao.SelectMenuPermsByUserId(c, loginService.data, userId)

		//for _, perm := range mysqlPerms {
		//	if len(perm) != 0 {
		//		perms = append(perms, perm)
		//	}
		//}
	}
	return perms
}

func (loginService *LoginService) GenerateCode(c *gin.Context) (m *systemModels.CaptchaVo) {
	captcha := base64Captcha.NewCaptcha(loginService.driver, loginService.store)
	id, b64s, _, err := captcha.Generate()
	if err != nil {
		panic(err)
	}
	m = new(systemModels.CaptchaVo)
	m.Id = id
	m.Img = b64s
	return m
}

func (loginService *LoginService) VerityCaptcha(c *gin.Context, id, base64 string) bool {
	return loginService.store.Verify(id, base64, true)
}

func (loginService *LoginService) ForceLogout(c *gin.Context, token string) {
	panic("等待补充")
}

func (loginService *LoginService) RolePermissionByRoles(roles []*systemModels.SysRole) (rolePerms []string, loginRoles []*baize.Role) {
	loginRoles = make([]*baize.Role, 0, len(roles))
	rolePerms = make([]string, 0, len(roles))
	for _, role := range roles {
		rolePerms = append(rolePerms, role.RoleKey)
		loginRoles = append(loginRoles, &baize.Role{RoleId: role.RoleId, DataScope: role.DataScope})
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
