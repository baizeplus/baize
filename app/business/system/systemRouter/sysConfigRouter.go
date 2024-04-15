package systemRoutes

import (
	"baize/app/business/system/systemController"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysConfigRouter(router *gin.RouterGroup, config *systemController.Config) {
	systemConfig := router.Group("/system/config")
	systemConfig.GET("/list", middlewares.HasPermission("system:config:list"), config.ConfigList)
	systemConfig.POST("/export", middlewares.HasPermission("system:config:export"), config.ConfigExport)
	systemConfig.GET("/:configId", middlewares.HasPermission("system:config:query"), config.ConfigGetInfo)
	systemConfig.POST("", middlewares.SetLog(configManagement, middlewares.Insert), middlewares.HasPermission("system:config:add"), config.ConfigAdd)
	systemConfig.PUT("", middlewares.SetLog(configManagement, middlewares.Update), middlewares.HasPermission("system:config:edit"), config.ConfigEdit)
	systemConfig.DELETE("/:configId", middlewares.SetLog(configManagement, middlewares.Delete), middlewares.HasPermission("system:config:remove"), config.ConfigRemove)

}
