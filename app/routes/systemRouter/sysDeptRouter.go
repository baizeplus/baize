package systemRoutes

import (
	"baize/app/bzSystem/controller"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysDeptRouter(router *gin.RouterGroup, deptController *controller.DeptController) {
	systemDept := router.Group("/system/dept")
	systemDept.GET("/list", middlewares.HasPermission("system:dept:list"), deptController.DeptList)
	systemDept.GET("/:deptId", middlewares.HasPermission("system:dept:query"), deptController.DeptGetInfo)
	systemDept.GET("/roleDeptTreeSelect/:roleId", middlewares.HasPermission("system:dept:query"), deptController.RoleDeptTreeSelect)
	systemDept.POST("", middlewares.HasPermission("system:dept:add"), deptController.DeptAdd)
	systemDept.PUT("", middlewares.HasPermission("system:dept:edit"), deptController.DeptEdit)
	systemDept.DELETE("/:deptId", middlewares.HasPermission("system:dept:remove"), deptController.DeptRemove)

}
