package IOFile

import (
	"baize/app/utils/pathUtils"
	"bytes"
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

func (l *localHostIOFile) PublicUploadFile(file *fileParams) (string, error) {
	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(file.data)
	if err != nil {
		return "", err
	}
	b := buf.Bytes()
	pathAndName := l.publicPath + file.keyName
	err = pathUtils.CreateMutiDir(filepath.Dir(pathAndName))
	if err != nil {
		return "", err
	}
	err = os.WriteFile(pathAndName, b, 0664)
	if err != nil {
		return "", err
	}
	return l.domainName + ResourcePrefix + "/" + file.keyName, nil
}

func (l *localHostIOFile) privateUploadFile(file *fileParams) (string, error) {
	buf := &bytes.Buffer{}
	_, err := buf.ReadFrom(file.data)
	if err != nil {
		return "", err
	}
	pathAndName := l.privatePath + file.keyName
	err = pathUtils.CreateMutiDir(filepath.Dir(pathAndName))
	if err != nil {
		return "", err
	}
	b := buf.Bytes()
	err = os.WriteFile(pathAndName, b, 0664)
	if err != nil {
		return "", err
	}
	return file.keyName, nil
}
