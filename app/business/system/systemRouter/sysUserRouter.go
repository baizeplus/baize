package systemRoutes

import (
	"baize/app/business/system/systemController"
	"baize/app/middlewares"

	"github.com/gin-gonic/gin"
)

func InitSysUserRouter(router *gin.RouterGroup, user *systemController.User) {
	systemUser := router.Group("/system/user")
	systemUser.GET("/list", middlewares.HasPermission("system:user:list"), user.UserList)
	systemUser.GET("/", middlewares.HasPermission("system:user:query"), user.UserGetInfo)
	systemUser.GET("/authRole/:userId", middlewares.HasPermission("system:user:edit"), user.UserAuthRole)
	systemUser.GET("/:userId", middlewares.HasPermission("system:user:query"), user.UserGetInfoById)
	systemUser.POST("", middlewares.HasPermission("system:user:add"), user.UserAdd)
	systemUser.PUT("", middlewares.HasPermission("system:user:edit"), user.UserEdit)
	systemUser.PUT("/resetPwd", middlewares.HasPermission("system:user:edit"), user.ResetPwd)
	systemUser.PUT("/changeStatus", middlewares.HasPermission("system:user:edit"), user.ChangeStatus)
	systemUser.DELETE("", middlewares.HasPermission("system:user:remove"), user.UserRemove)
	systemUser.POST("/importData", middlewares.HasPermission("system:user:import"), user.UserImportData)
	systemUser.POST("/importTemplate", middlewares.HasPermission("system:user:add"), user.ImportTemplate)
	systemUser.POST("/export", middlewares.HasPermission("system:user:export"), user.UserExport)
	systemUser.PUT("/authRole", middlewares.HasPermission("system:user:edit"), user.InsertAuthRole)
}
