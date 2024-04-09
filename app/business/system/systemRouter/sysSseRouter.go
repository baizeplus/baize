package systemRoutes

import (
	"baize/app/business/system/systemController"
	"github.com/gin-gonic/gin"
)

func InitSseRouter(router *gin.RouterGroup, sse *systemController.Sse) {
	systemUser := router.Group("/system")
	systemUser.GET("/sse/:token", sse.BuildSse)
}
