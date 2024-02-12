package systemRoutes

import (
	"baize/app/bzSystem/controller"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysMenuRouter(router *gin.RouterGroup, menuController *controller.MenuController) {
	menu := router.Group("/system/menu")
	menu.GET("/list", middlewares.HasPermission("system:menu:list"), menuController.MenuList)
	menu.GET("/:menuId", middlewares.HasPermission("system:menu:query"), menuController.MenuGetInfo)
	menu.GET("/treeSelect", menuController.MenuTreeSelect)
	menu.POST("", middlewares.HasPermission("system:menu:add"), menuController.MenuAdd)
	menu.PUT("", middlewares.HasPermission("system:menu:edit"), menuController.MenuEdit)
	menu.DELETE("/:menuId", middlewares.HasPermission("system:menu:remove"), menuController.MenuRemove)
	menu.GET("/roleMenuTreeSelect/:roleId", menuController.RoleMenuTreeSelect)
}
