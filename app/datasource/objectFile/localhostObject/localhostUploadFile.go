package localhostObject

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

type LocalHostFile struct {
	PublicPath  string
	PrivatePath string
	DomainName  string
}

func (l *LocalHostFile) PublicUploadFile(ctx context.Context, file multipart.File, keyName string) (string, error) {
	if err := os.MkdirAll(filepath.Dir(l.PublicPath+keyName), 0750); err != nil {
		return "", err
	}
	out, err := os.Create(l.PublicPath + keyName)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	return l.DomainName + ResourcePrefix + "/" + keyName, nil
}

func (l *LocalHostFile) PrivateUploadFile(ctx context.Context, file multipart.File, keyName string) (string, error) {
	if err := os.MkdirAll(filepath.Dir(l.PrivatePath+keyName), 0750); err != nil {
		return "", err
	}
	out, err := os.Create(l.PrivatePath + keyName)
	if err != nil {
		return "", err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	return keyName, nil
}
