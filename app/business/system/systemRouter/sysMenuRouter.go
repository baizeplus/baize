package systemRoutes

import (
	"baize/app/business/system/systemController"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysMenuRouter(router *gin.RouterGroup, menu *systemController.Menu) {
	rm := router.Group("/system/menu")
	rm.GET("/list", middlewares.HasPermission("system:menu:list"), menu.MenuList)
	rm.GET("/:menuId", middlewares.HasPermission("system:menu:query"), menu.MenuGetInfo)
	rm.GET("/treeSelect", menu.MenuTreeSelect)
	rm.POST("", middlewares.SetLog(menuManagement, middlewares.Insert), middlewares.HasPermission("system:menu:add"), menu.MenuAdd)
	rm.PUT("", middlewares.SetLog(menuManagement, middlewares.Update), middlewares.HasPermission("system:menu:edit"), menu.MenuEdit)
	rm.DELETE("/:menuId", middlewares.SetLog(menuManagement, middlewares.Delete), middlewares.HasPermission("system:menu:remove"), menu.MenuRemove)
	rm.GET("/roleMenuTreeSelect/:roleId", menu.RoleMenuTreeSelect)
}
