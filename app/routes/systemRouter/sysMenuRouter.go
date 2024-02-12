package systemRoutes

import (
	"baize/app/bzSystem/controller"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysMenuRouter(router *gin.RouterGroup, menu *controller.Menu) {
	rm := router.Group("/system/menu")
	rm.GET("/list", middlewares.HasPermission("system:menu:list"), menu.MenuList)
	rm.GET("/:menuId", middlewares.HasPermission("system:menu:query"), menu.MenuGetInfo)
	rm.GET("/treeSelect", menu.MenuTreeSelect)
	rm.POST("", middlewares.HasPermission("system:menu:add"), menu.MenuAdd)
	rm.PUT("", middlewares.HasPermission("system:menu:edit"), menu.MenuEdit)
	rm.DELETE("/:menuId", middlewares.HasPermission("system:menu:remove"), menu.MenuRemove)
	rm.GET("/roleMenuTreeSelect/:roleId", menu.RoleMenuTreeSelect)
}
