package systemServiceImpl

import (
	"baize/app/business/system/systemService"
	"baize/app/datasource/objectFile"
	"baize/app/utils/baizeContext"
	"baize/app/utils/fileUtils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mime/multipart"
	"path/filepath"
	"sync"
)

type FileService struct {
	of objectFile.ObjectFile
}

func NewFileService(of objectFile.ObjectFile) systemService.IFileService {
	return &FileService{
		of: of,
	}
}

func (fs *FileService) UploadFileRandomName(c *gin.Context, file *multipart.FileHeader) string {
	open, _ := file.Open()
	defer open.Close()
	name := fileUtils.GetRandomName(baizeContext.GetUserId(c), filepath.Ext(file.Filename))
	url, err := fs.of.PublicUploadFile(c, open, name)
	if err != nil {
		panic(err)
	}
	return url
}

func (fs *FileService) UploadFilesRandomName(c *gin.Context, files []*multipart.FileHeader) []string {

	strings := make([]string, len(files), len(files))
	var wg sync.WaitGroup
	for i, file := range files {
		wg.Add(1)
		go func(i1 int, f *multipart.FileHeader) {
			defer func() {
				if e := recover(); e != nil {
					zap.L().Error("fileServiceImpl.UploadFilesRandomName")
				}
			}()
			open, err := f.Open()
			if err != nil {
				panic(err)
			}
			name := fileUtils.GetRandomName(baizeContext.GetUserId(c), filepath.Ext(f.Filename))
			url, err := fs.of.PublicUploadFile(c, open, name)
			if err != nil {
				panic(err)
			}
			strings[i1] = url
			wg.Done()
		}(i, file)

	}
	wg.Wait()
	return strings
}

func (fs *FileService) UploadFileOriginalName(c *gin.Context, file *multipart.FileHeader) string {
	open, _ := file.Open()
	defer open.Close()
	name := fileUtils.GetRandomPath(baizeContext.GetUserId(c), filepath.Ext(file.Filename))
	url, err := fs.of.PublicUploadFile(c, open, name)
	if err != nil {
		panic(err)
	}
	return url
}

func (fs *FileService) UploadPrivateFileOriginalName(c *gin.Context, file *multipart.FileHeader) string {
	open, _ := file.Open()
	defer open.Close()
	name := fileUtils.GetRandomName(baizeContext.GetUserId(c), filepath.Ext(file.Filename))
	url, err := fs.of.PrivateUploadFile(c, open, name)
	if err != nil {
		panic(err)
	}
	return url
}

func (fs *FileService) DownloadPrivateFileRandomName(c *gin.Context, fileKey string) ([]byte, string) {
	//b, c, err := fs.s3.DownloadPrivateS3(ctx, &fileKey)
	//if err != nil {
	//	panic(err)
	//}
	//return b, c
	panic("")
}
