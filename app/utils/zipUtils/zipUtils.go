package zipUtils

import (
	"archive/zip"
	"fmt"
)

func AddFileToZip(zipWriter *zip.Writer, filename, content string) error {
	writer, err := zipWriter.Create(filename)
	if err != nil {
		return fmt.Errorf("无法创建 zip 归档: %v", err)
	}
	_, err = writer.Write([]byte(content))
	if err != nil {
		return fmt.Errorf("无法写入 zip 归档: %v", err)
	}
	return nil
}
