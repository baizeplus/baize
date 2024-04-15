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
	systemUser.GET("/authRole/:userId", middlewares.SetLog(userManagement, middlewares.Update), middlewares.HasPermission(SystemUserEdit), user.UserAuthRole)
	systemUser.GET("/:userId", middlewares.HasPermission(SystemUserQuery), user.UserGetInfoById)
	systemUser.POST("", middlewares.SetLog(userManagement, middlewares.Insert), middlewares.HasPermission("system:user:add"), user.UserAdd)
	systemUser.GET("/dataScope/:userId", middlewares.HasPermission(SystemUserQuery), user.SelectUserDataScope)
	systemUser.PUT("/dataScope", middlewares.SetLog(userManagement, middlewares.Other), middlewares.HasPermission(SystemUserEdit), user.UpdateUserDataScope)
	systemUser.PUT("", middlewares.SetLog(userManagement, middlewares.Update), middlewares.HasPermission(SystemUserEdit), user.UserEdit)
	systemUser.PUT("/resetPwd", middlewares.SetLog(userManagement, middlewares.Update), middlewares.HasPermission(SystemUserEdit), user.ResetPwd)
	systemUser.PUT("/changeStatus", middlewares.SetLog(userManagement, middlewares.Update), middlewares.HasPermission(SystemUserEdit), user.ChangeStatus)
	systemUser.DELETE("/:userIds", middlewares.SetLog(userManagement, middlewares.Delete), middlewares.HasPermission("system:user:remove"), user.UserRemove)
	systemUser.POST("/importData", middlewares.HasPermission("system:user:import"), user.UserImportData)
	systemUser.POST("/importTemplate", middlewares.HasPermission("system:user:add"), user.ImportTemplate)
	systemUser.POST("/export", middlewares.HasPermission("system:user:export"), user.UserExport)
	systemUser.PUT("/authRole", middlewares.SetLog(userManagement, middlewares.Update), middlewares.HasPermission(SystemUserEdit), user.InsertAuthRole)
}
