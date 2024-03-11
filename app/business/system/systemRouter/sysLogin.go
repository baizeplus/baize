package systemRoutes

import (
	"baize/app/business/system/systemController"
	"github.com/gin-gonic/gin"
)

func InitLoginRouter(router *gin.RouterGroup, login *systemController.Login) {
	router.GET("/captchaImage", login.GetCode) //获取验证码
	router.POST("/login", login.Login)         //登录
	router.POST("/register", login.Register)   //登录
	router.POST("/logout", login.Logout)
}
func InitGetUser(router *gin.RouterGroup, login *systemController.Login) {
	router.GET("/getInfo", login.GetInfo)
	router.GET("/getRouters", login.GetRouters)

}
