package systemRoutes

import (
	"baize/app/business/system/systemController"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysRoleRouter(router *gin.RouterGroup, role *systemController.Role) {
	rr := router.Group("/system/role")
	rr.GET("/list", middlewares.HasPermission(SystemRoleList), role.RoleList)
	rr.POST("/export", middlewares.HasPermission("system:role:export"), role.RoleExport)
	rr.GET("/:roleId", middlewares.HasPermission("system:role:query"), role.RoleGetInfo)
	rr.POST("", middlewares.SetLog(roleManagement, middlewares.Insert), middlewares.HasPermission("system:role:add"), role.RoleAdd)
	rr.PUT("", middlewares.SetLog(roleManagement, middlewares.Update), middlewares.HasPermission(SystemRoleEdit), role.RoleEdit)
	rr.PUT("/changeStatus", middlewares.HasPermission(SystemRoleEdit), role.RoleChangeStatus)
	rr.DELETE("/:rolesIds", middlewares.SetLog(roleManagement, middlewares.Delete), middlewares.HasPermission("system:role:remove"), role.RoleRemove)
	rr.GET("/authUser/allocatedList", middlewares.HasPermission(SystemRoleList), role.AllocatedList)
	rr.GET("/authUser/unallocatedList", middlewares.HasPermission(SystemRoleList), role.UnallocatedList)
	rr.PUT("/authUser/selectAll", middlewares.SetLog(roleManagement, middlewares.Update), middlewares.HasPermission(SystemRoleEdit), role.InsertAuthUser)
	rr.PUT("/authUser/cancelAll", middlewares.SetLog(roleManagement, middlewares.Update), middlewares.HasPermission(SystemRoleEdit), role.CancelAuthUserAll)
	rr.PUT("/authUser/cancel", middlewares.SetLog(roleManagement, middlewares.Update), middlewares.HasPermission(SystemRoleEdit), role.CancelAuthUser)

}
