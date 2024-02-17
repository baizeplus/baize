package routes

import (
	"baize/app/baize"
	controllerM "baize/app/bzMonitor/controller"
	"baize/app/bzSystem/controller"
	"baize/app/middlewares"
	"baize/app/routes/monitorRouter"
	systemRoutes "baize/app/routes/systemRouter"
	"baize/app/setting"
	"baize/app/utils/IOFile"
	"baize/app/utils/logger"
	"github.com/google/wire"
	swaggerFiles "github.com/swaggo/files"
	"net/http"
	"strings"

	"baize/app/docs"
	gs "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

var ProviderSet = wire.NewSet(NewGinEngine)
var noRefresh = baize.NewSet([]string{})

func NewGinEngine(
	login *controller.Login,
	user *controller.User,
	dept *controller.Dept,
	dictType *controller.DictType,
	dictData *controller.DictData,
	menu *controller.Menu,
	role *controller.Role,
	post *controller.Post,
	profile *controller.Profile,
	server *controllerM.InfoServer,
	userOnline *controllerM.UserOnline,
	logfor *controllerM.Logininfor,
) *gin.Engine {

	if setting.Conf.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery())
	r.Use(Cors())
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
		monitorRouter.InitServerRouter(group, server)
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

// Cors
// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "*")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
