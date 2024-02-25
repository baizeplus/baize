package systemRoutes

import (
	"baize/app/business/system/systemController"
	"github.com/gin-gonic/gin"
)

func InitSysProfileRouter(router *gin.RouterGroup, profile *systemController.Profile) {
	systemProfile := router.Group("/system/user/profile")
	systemProfile.GET("", profile.Profile)
	systemProfile.PUT("", profile.ProfileUpdateProfile)
	systemProfile.PUT("/updatePwd", profile.ProfileUpdatePwd)
	systemProfile.POST("/avatar", profile.ProfileAvatar)
}
