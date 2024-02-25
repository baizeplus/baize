package systemRoutes

import (
	"baize/app/business/system/systemController"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysDeptRouter(router *gin.RouterGroup, dept *systemController.Dept) {
	systemDept := router.Group("/system/dept")
	systemDept.GET("/list", middlewares.HasPermissions([]string{"system:dept:list", "system:user:list"}), dept.DeptList)
	systemDept.GET("/:deptId", middlewares.HasPermission("system:dept:query"), dept.DeptGetInfo)
	systemDept.GET("/roleDeptTreeSelect/:roleId", middlewares.HasPermission("system:dept:query"), dept.RoleDeptTreeSelect)
	systemDept.POST("", middlewares.HasPermission("system:dept:add"), dept.DeptAdd)
	systemDept.PUT("", middlewares.HasPermission("system:dept:edit"), dept.DeptEdit)
	systemDept.DELETE("/:deptId", middlewares.HasPermission("system:dept:remove"), dept.DeptRemove)

}
