package systemRoutes

import (
	"baize/app/business/system/systemController"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitSysPostRouter(router *gin.RouterGroup, post *systemController.Post) {
	systemPost := router.Group("/system/post")
	systemPost.GET("/list", middlewares.HasPermission("system:post:list"), post.PostList)
	systemPost.POST("/export", middlewares.HasPermission("system:post:export"), post.PostExport)
	systemPost.GET("/:postId", middlewares.HasPermission("system:post:query"), post.PostGetInfo)
	systemPost.POST("", middlewares.SetLog(postManagement, middlewares.Insert), middlewares.HasPermission("system:post:add"), post.PostAdd)
	systemPost.PUT("", middlewares.SetLog(postManagement, middlewares.Update), middlewares.HasPermission("system:post:edit"), post.PostEdit)
	systemPost.DELETE("/:postIds", middlewares.SetLog(postManagement, middlewares.Delete), middlewares.HasPermission("system:post:remove"), post.PostRemove)

}
