package routes

import (
	"baize/app/baize"
	"baize/app/business/monitor/monitorController"
	"baize/app/business/monitor/monitorRouter"
	"baize/app/business/system/systemController"
	"baize/app/business/system/systemRouter"
	"baize/app/business/tool/toolController"
	"baize/app/business/tool/toolRouter"
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
	sc *systemController.System,
	mc *monitorController.Monitor,
	gc *toolController.Tool,
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
		if setting.Conf.UploadFile.Type == "local" {
			group.Static(IOFile.ResourcePrefix, setting.Conf.UploadFile.Localhost.PublicPath)
		}
		systemRoutes.InitLoginRouter(group, sc.Login) //获取登录信息
		systemRoutes.InitSseRouter(group, sc.Sse)     //SSE链接

	}
	//做鉴权的
	group.Use(middlewares.SessionAuthMiddleware(noRefresh))
	{
		systemRoutes.InitSysProfileRouter(group, sc.Profile)   //个人信息
		systemRoutes.InitGetUser(group, sc.Login)              //获取登录信息
		systemRoutes.InitSysUserRouter(group, sc.User)         //用户相关
		systemRoutes.InitSysDeptRouter(group, sc.Dept)         //部门相关
		systemRoutes.InitSysDictTypeRouter(group, sc.DictType) //数据字典属性
		systemRoutes.InitSysDictDataRouter(group, sc.DictData) //数据字典信息
		systemRoutes.InitSysMenuRouter(group, sc.Menu)         //菜单相关
		systemRoutes.InitSysRoleRouter(group, sc.Role)         //角色相关
		systemRoutes.InitSysPostRouter(group, sc.Post)         //岗位属性
		systemRoutes.InitSysConfigRouter(group, sc.Config)     //配置文件
		systemRoutes.InitFileRouter(group, sc.File)            //文件管理
		systemRoutes.InitSysRouterRouter(group, sc.Notice)     //文件管理
		monitorRouter.InitServerRouter(group, mc.Server)
		monitorRouter.InitSysOperLogRouter(group, mc.Oper)
		monitorRouter.InitSysUserOnlineRouter(group, mc.UserOnline) //在线用户监控
		monitorRouter.InitSysLogininforRouter(group, mc.Logfor)
		monitorRouter.InitJobRouter(group, mc.Job)
		toolRouter.InitGenTableRouter(group, gc.GenTable)
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
