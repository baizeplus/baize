package systemRoutes

import (
	"baize/app/business/system/systemController"
	"baize/app/middlewares"

	"github.com/gin-gonic/gin"
)

func InitSysUserRouter(router *gin.RouterGroup, user *systemController.User) {
	systemUser := router.Group("/system/user")
	systemUser.GET("/list", middlewares.HasPermission("system:user:list"), user.UserList)
	systemUser.GET("/", middlewares.HasPermission(SystemUserQuery), user.UserGetInfo)
	systemUser.GET("/authRole/:userId", middlewares.HasPermission(SystemUserEdit), user.UserAuthRole)
	systemUser.GET("/:userId", middlewares.HasPermission(SystemUserQuery), user.UserGetInfoById)
	systemUser.POST("", middlewares.SetLog("用户", middlewares.INSERT), middlewares.HasPermission("system:user:add"), user.UserAdd)
	systemUser.GET("/dataScope/:userId", middlewares.HasPermission(SystemUserQuery), user.SelectUserDataScope)
	systemUser.PUT("/dataScope", middlewares.HasPermission(SystemUserEdit), user.UpdateUserDataScope)
	systemUser.PUT("", middlewares.HasPermission(SystemUserEdit), user.UserEdit)
	systemUser.PUT("/resetPwd", middlewares.HasPermission(SystemUserEdit), user.ResetPwd)
	systemUser.PUT("/changeStatus", middlewares.HasPermission(SystemUserEdit), user.ChangeStatus)
	systemUser.DELETE("/:userIds", middlewares.HasPermission("system:user:remove"), user.UserRemove)
	systemUser.POST("/importData", middlewares.HasPermission("system:user:import"), user.UserImportData)
	systemUser.POST("/importTemplate", middlewares.HasPermission("system:user:add"), user.ImportTemplate)
	systemUser.POST("/export", middlewares.HasPermission("system:user:export"), user.UserExport)
	systemUser.PUT("/authRole", middlewares.HasPermission(SystemUserEdit), user.InsertAuthRole)
}
