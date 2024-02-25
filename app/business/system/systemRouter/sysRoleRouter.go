package systemRoutes

import (
	"baize/app/business/system/systemController"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysRoleRouter(router *gin.RouterGroup, role *systemController.Role) {
	rr := router.Group("/system/role")
	rr.GET("/list", middlewares.HasPermission("system:role:list"), role.RoleList)
	rr.POST("/export", middlewares.HasPermission("system:role:export"), role.RoleExport)
	rr.GET("/:roleId", middlewares.HasPermission("system:role:query"), role.RoleGetInfo)
	rr.POST("", middlewares.HasPermission("system:role:add"), role.RoleAdd)
	rr.PUT("", middlewares.HasPermission("system:role:edit"), role.RoleEdit)
	rr.PUT("/dataScope", middlewares.HasPermission("system:role:edit"), role.RoleDataScope)
	rr.PUT("/changeStatus", middlewares.HasPermission("system:role:edit"), role.RoleChangeStatus)
	rr.DELETE("/:rolesIds", middlewares.HasPermission("system:role:remove"), role.RoleRemove)
	rr.GET("/authUser/allocatedList", middlewares.HasPermission("system:role:list"), role.AllocatedList)
	rr.GET("/authUser/unallocatedList", middlewares.HasPermission("system:role:list"), role.UnallocatedList)
	rr.PUT("/authUser/selectAll", middlewares.HasPermission("system:role:edit"), role.InsertAuthUser)
	rr.PUT("/authUser/cancelAll", middlewares.HasPermission("system:role:edit"), role.CancelAuthUserAll)
	rr.PUT("/authUser/cancel", middlewares.HasPermission("system:role:edit"), role.CancelAuthUser)

}
