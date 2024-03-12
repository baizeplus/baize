package systemService

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

type IFileService interface {
	UploadFileRandomName(c *gin.Context, file *multipart.FileHeader) string
	UploadFilesRandomName(c *gin.Context, file []*multipart.FileHeader) []string
	UploadFileOriginalName(c *gin.Context, file *multipart.FileHeader) string

	UploadPrivateFileOriginalName(c *gin.Context, file *multipart.FileHeader) string
	DownloadPrivateFileRandomName(c *gin.Context, fileKey string) ([]byte, string)
}
