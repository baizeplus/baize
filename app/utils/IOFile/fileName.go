package IOFile

import (
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"time"
)

var fileExtension = map[string]string{
	"image/png":  "png",
	"image/jpg":  "jpg",
	"image/jpeg": "jpeg",
	"image/bmp":  "bmp",
	"image/gif":  "gif",
}

func GetTenantRandomName(userId int64, extensionName string) string {
	uuid, _ := uuid.NewRandom()
	nameKey := gconv.String(userId) + "/" + time.Now().Format(time.DateOnly) + "/" + uuid.String() + extensionName
	return nameKey
}

func GetExtension(file *multipart.FileHeader) string {
	ext := filepath.Ext(file.Filename)
	if ext == "" {
		open, _ := file.Open()
		data := make([]byte, 512)
		_, err := open.Read(data)
		if err != nil {
			panic(err)
		}
		return fileExtension[http.DetectContentType(data)]
	}
	return ext
}
