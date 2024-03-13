package routes

import (
	"baize/app/baize"
	"baize/app/business/monitor/monitorController"
	"baize/app/business/monitor/monitorRouter"
	"baize/app/business/system/systemController"
	systemRoutes "baize/app/business/system/systemRouter"
	"time"

	"baize/app/middlewares"

	"baize/app/docs"
	"baize/app/setting"
	"baize/app/utils/IOFile"
	"baize/app/utils/logger"
	"github.com/gin-contrib/cors"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var ProviderSet = wire.NewSet(NewGinEngine)
var noRefresh = baize.NewSet([]string{})

func NewGinEngine(
	login *systemController.Login,
	user *systemController.User,
	dept *systemController.Dept,
	dictType *systemController.DictType,
	dictData *systemController.DictData,
	menu *systemController.Menu,
	role *systemController.Role,
	post *systemController.Post,
	profile *systemController.Profile,
	config *systemController.Config,
	file *systemController.File,
	server *monitorController.InfoServer,
	userOnline *monitorController.UserOnline,
	logfor *monitorController.Logininfor,
	oper *monitorController.OperLog,
) *gin.Engine {

	if setting.Conf.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery())
	r.Use(newCors())
	group := r.Group("")

	host := setting.Conf.Host
	docs.SwaggerInfo.Host = host[strings.Index(host, "//")+2:]
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	group.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//不做鉴权的
	{
		if setting.Conf.UploadFile.Type == "localhost" {
			group.Static(IOFile.ResourcePrefix, setting.Conf.UploadFile.Localhost.PublicPath)
		}
		systemRoutes.InitLoginRouter(group, login) //获取登录信息

	}
	//做鉴权的
	group.Use(middlewares.SessionAuthMiddleware(noRefresh))
	{
		systemRoutes.InitSysProfileRouter(group, profile)   //个人信息
		systemRoutes.InitGetUser(group, login)              //获取登录信息
		systemRoutes.InitSysUserRouter(group, user)         //用户相关
		systemRoutes.InitSysDeptRouter(group, dept)         //部门相关
		systemRoutes.InitSysDictTypeRouter(group, dictType) //数据字典属性
		systemRoutes.InitSysDictDataRouter(group, dictData) //数据字典信息
		systemRoutes.InitSysMenuRouter(group, menu)         //菜单相关
		systemRoutes.InitSysRoleRouter(group, role)         //角色相关
		systemRoutes.InitSysPostRouter(group, post)         //岗位属性
		systemRoutes.InitSysConfigRouter(group, config)     //配置文件
		systemRoutes.InitFileRouter(group, file)            //配置文件
		monitorRouter.InitServerRouter(group, server)
		monitorRouter.InitSysOperLogRouter(group, oper)
		monitorRouter.InitSysUserOnlineRouter(group, userOnline) //在线用户监控
		monitorRouter.InitSysLogininforRouter(group, logfor)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r

}

func newCors() gin.HandlerFunc {
	ss := []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Cache-Control", "Content-Language", "Content-Type", "Expires", "Last-Modified", "Pragma", "FooBar"}
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"} //允许访问的域名
	config.AllowMethods = []string{"PUT", "PATCH", "GET", "POST", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"*"}
	config.ExposeHeaders = ss
	config.MaxAge = time.Hour
	config.AllowCredentials = false
	return cors.New(config)

}
