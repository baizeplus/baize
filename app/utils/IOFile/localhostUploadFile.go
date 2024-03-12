package IOFile

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

const (
	ResourcePrefix = "/profile" // 资源映射路径 前缀
)

type localHostIOFile struct {
	publicPath  string
	privatePath string
	domainName  string
}

func (l *localHostIOFile) PublicUploadFile(ctx context.Context, file multipart.File, keyName string) (string, error) {
	if err := os.MkdirAll(filepath.Dir(l.publicPath+keyName), 0750); err != nil {
		return "", err
	}
	out, err := os.Create(l.publicPath + keyName)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	return l.domainName + ResourcePrefix + "/" + keyName, nil
}

func (l *localHostIOFile) PrivateUploadFile(ctx context.Context, file multipart.File, keyName string) (string, error) {
	if err := os.MkdirAll(filepath.Dir(l.privatePath+keyName), 0750); err != nil {
		return "", err
	}
	out, err := os.Create(l.privatePath + keyName)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	return keyName, nil
}
