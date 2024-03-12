package systemRoutes

import (
	"baize/app/business/system/systemController"
	"github.com/gin-gonic/gin"
)

func InitFileRouter(router *gin.RouterGroup, fc *systemController.File) {
	systemDictType := router.Group("/file")
	systemDictType.POST("/uploadFileRandomName", fc.UploadFileRandomName)
	systemDictType.POST("/uploadFileOriginalName", fc.UploadFileOriginalName)
	systemDictType.POST("/uploadFiles", fc.UploadFiles)
	systemDictType.POST("/uploadPrivateFileOriginalName", fc.UploadPrivateFileOriginalName)
	systemDictType.GET("/downloadPrivateFile", fc.DownloadPrivateFile)
}
