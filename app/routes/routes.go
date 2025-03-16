package routes

import (
	"baize/app/business/monitor/monitorController"
	"baize/app/business/system/systemController"
	"baize/app/business/tool/toolController"
	"baize/app/datasource/cache"
	"baize/app/datasource/objectFile/localhostObject"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"time"

	"baize/app/middlewares"

	"baize/app/docs"
	"baize/app/setting"
	"baize/app/utils/logger"
	"github.com/gin-contrib/cors"
	"github.com/google/wire"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var ProviderSet = wire.NewSet(NewGinEngine)

func NewGinEngine(
	cache cache.Cache,
	sc *systemController.System,
	mc *monitorController.Monitor,
	gc *toolController.Tool,
) *gin.Engine {

	if setting.Conf.Mode != "dev" {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.NewLoggerMiddlewareBuilder().
		IgnorePaths("/ping").Build())
	r.Use(newCors())
	group := r.Group("")
	if setting.Conf.Mode != "dev" {
		host := setting.Conf.Host
		docs.SwaggerInfo.Host = host[strings.Index(host, "//")+2:]
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		group.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	}

	//不做鉴权的
	{

		if viper.GetString("upload_file.type") == "local" {
			group.Static(localhostObject.ResourcePrefix, viper.GetString("upload_file.localhost.public_path"))
		}
		sc.Login.PublicRoutes(group) //登录
		sc.Sse.PublicRoutes(group)   //SSE链接

	}
	//做鉴权的
	group.Use(middlewares.NewSessionAuthMiddlewareBuilder(cache).Build())
	{

		sc.Profile.PrivateRoutes(group)    //个人资料
		sc.Login.PrivateRoutes(group)      //登录
		sc.User.PrivateRoutes(group)       //用户
		sc.Dept.PrivateRoutes(group)       //部门
		sc.DictType.PrivateRoutes(group)   //字典类型
		sc.DictData.PrivateRoutes(group)   //地点数据
		sc.Role.PrivateRoutes(group)       //角色
		sc.Post.PrivateRoutes(group)       //岗位
		sc.Permission.PrivateRoutes(group) //岗位
		sc.Config.PrivateRoutes(group)     //配置
		sc.File.PrivateRoutes(group)       //文件
		sc.Notice.PrivateRoutes(group)     //消息
		mc.Server.PrivateRoutes(group)     //服务器详情
		mc.Oper.PrivateRoutes(group)       //操作日志
		mc.UserOnline.PrivateRoutes(group) //在线用户
		mc.Logfor.PrivateRoutes(group)     //登录日志
		mc.Job.PrivateRoutes(group)        //定时任务
		gc.GenTable.PrivateRoutes(group)   //代码生成
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
