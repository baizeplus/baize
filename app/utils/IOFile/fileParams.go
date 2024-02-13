package IOFile

import (
	"mime/multipart"
	"net/http"
)

func NewFileParamsRandomName(keyName string, file multipart.File) *fileParams {
	f := new(fileParams)
	f.keyName = keyName
	f.data = file
	data := make([]byte, 512)
	_, err := file.Read(data)
	if err != nil {
		panic(err)
	}
	f.contentType = http.DetectContentType(data)
	return f
}
