package systemRoutes

import (
	"baize/app/bzSystem/controller"
	"github.com/gin-gonic/gin"
)

func InitLoginRouter(router *gin.RouterGroup, login *controller.Login) {
	router.GET("/captchaImage", login.GetCode) //获取验证码
	router.POST("/login", login.Login)         //登录
	router.POST("/logout", login.Logout)
}
func InitGetUser(router *gin.RouterGroup, login *controller.Login) {
	router.GET("/getInfo", login.GetInfo)
	router.GET("/getRouters", login.GetRouters)

}
