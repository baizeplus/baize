package fileUtils

import (
	"mime/multipart"
	"net/http"
)

func GetFileContentType(file multipart.File) string {

	data := make([]byte, 512)
	_, err := file.Read(data)
	if err != nil {
		panic(err)
	}
	return http.DetectContentType(data)
}
