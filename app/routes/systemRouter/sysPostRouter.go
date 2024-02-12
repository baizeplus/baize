package systemRoutes

import (
	"baize/app/bzSystem/controller"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysPostRouter(router *gin.RouterGroup, post *controller.Post) {
	systemPost := router.Group("/system/post")
	systemPost.GET("/list", middlewares.HasPermission("system:post:list"), post.PostList)
	systemPost.GET("/export", middlewares.HasPermission("system:post:export"), post.PostExport)
	systemPost.GET("/:postId", middlewares.HasPermission("system:post:query"), post.PostGetInfo)
	systemPost.POST("", middlewares.HasPermission("system:post:add"), post.PostAdd)
	systemPost.PUT("", middlewares.HasPermission("system:post:edit"), post.PostEdit)
	systemPost.DELETE("/:postIds", middlewares.HasPermission("system:post:remove"), post.PostRemove)

}
