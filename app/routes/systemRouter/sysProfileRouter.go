package systemRoutes

import (
	"baize/app/bzSystem/controller"
	"github.com/gin-gonic/gin"
)

func InitSysProfileRouter(router *gin.RouterGroup, profile *controller.Profile) {
	systemProfile := router.Group("/system/user/profile")
	systemProfile.GET("", profile.Profile)
	systemProfile.PUT("", profile.ProfileUpdateProfile)
	systemProfile.PUT("/updatePwd", profile.ProfileUpdatePwd)
	systemProfile.POST("/avatar", profile.ProfileAvatar)
}
