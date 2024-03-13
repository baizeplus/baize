package systemController

import (
	"baize/app/business/system/systemService"
	"baize/app/business/system/systemService/systemServiceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type File struct {
	fs systemService.IFileService
}

func NewFile(fs *systemServiceImpl.FileService) *File {
	return &File{fs: fs}
}

// UploadFileRandomName 上传文件
// @Summary 上传文件随即文件名
// @Description 上传文件随即文件名
// @Tags 租户上传文件
// @Security BearerAuth
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Produce application/json
// @Success 200 {object} response.ResponseData
// @Router /file/uploadFileRandomName [post]
func (fc *File) UploadFileRandomName(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		baizeContext.ParameterError(c)
		return
	}
	_, err = file.Open()
	if err != nil {
		baizeContext.ParameterError(c)
		return
	}
	url := fc.fs.UploadFileRandomName(c, file)
	baizeContext.SuccessData(c, url)
}

// UploadFiles 上传多个文件
// @Summary 上传多个文件随即文件名
// @Description 上传多个文件随即文件名
// @Tags 租户上传文件
// @Security BearerAuth
// @Accept multipart/form-data
// @Param files formData file true "files"
// @Produce application/json
// @Success 200 {object} response.ResponseData
// @Router /file/uploadFiles [post]
func (fc *File) UploadFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		baizeContext.ParameterError(c)
		return
	}
	files := form.File["files"]

	url := fc.fs.UploadFilesRandomName(c, files)
	baizeContext.SuccessData(c, url)
}

// UploadFileOriginalName 上传文件
// @Summary 上传文件原始文件名
// @Description 上传文件原始文件名
// @Tags 租户上传文件
// @Security BearerAuth
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Produce application/json
// @Success 200 {object} response.ResponseData
// @Router /file/uploadFileOriginalName [post]
func (fc *File) UploadFileOriginalName(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		baizeContext.ParameterError(c)
		return
	}
	_, err = file.Open()
	if err != nil {
		baizeContext.ParameterError(c)
		return
	}
	url := fc.fs.UploadFileOriginalName(c, file)
	baizeContext.SuccessData(c, url)
}

// UploadPrivateFileOriginalName 上传私有文件原始文件名
// @Summary 上传私有文件原始文件名
// @Description 上传私有文件原始文件名
// @Tags 租户上传文件
// @Security BearerAuth
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Produce application/json
// @Success 200 {object} response.ResponseData
// @Router /file/uploadPrivateFileOriginalName [post]
func (fc *File) UploadPrivateFileOriginalName(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		baizeContext.ParameterError(c)
		return
	}
	_, err = file.Open()
	if err != nil {
		baizeContext.ParameterError(c)
		return
	}
	url := fc.fs.UploadPrivateFileOriginalName(c, file)
	baizeContext.SuccessData(c, url)
}

// DownloadPrivateFile 租户下载私有文件
// @Summary 租户下载私有文件
// @Description 租户下载私有文件
// @Tags 租户上传文件
// @Security BearerAuth
// @Param key query string true "key"
// @Produce application/octet-stream
// @Success 200
// @Router /file/downloadPrivateFile [get]
func (fc *File) DownloadPrivateFile(c *gin.Context) {
	//b, ct := fc.fs.DownloadPrivateFileRandomName(c, c.Query("key"))
	//knitContext.DataPackageFile(c, ct, b)
}
