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
	systemDept.POST("", middlewares.SetLog(deptManagement, middlewares.Insert), middlewares.HasPermission("system:dept:add"), dept.DeptAdd)
	systemDept.PUT("", middlewares.SetLog(deptManagement, middlewares.Update), middlewares.HasPermission("system:dept:edit"), dept.DeptEdit)
	systemDept.DELETE("/:deptId", middlewares.SetLog(deptManagement, middlewares.Delete), middlewares.HasPermission("system:dept:remove"), dept.DeptRemove)

}
