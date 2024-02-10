package systemRoutes

import (
	"baize/app/bzSystem/controller"
	"github.com/gin-gonic/gin"
)

func InitLoginRouter(router *gin.RouterGroup, loginController *controller.LoginController) {
	router.GET("/captchaImage", loginController.GetCode) //获取验证码
	router.POST("/login", loginController.Login)         //登录
	router.POST("/logout", loginController.Logout)
}
func InitGetUser(router *gin.RouterGroup, loginController *controller.LoginController) {
	router.GET("/getInfo", loginController.GetInfo)
	router.GET("/getRouters", loginController.GetRouters)

}
