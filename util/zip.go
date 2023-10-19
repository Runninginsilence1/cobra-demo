package util

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// UnzipWithTimeout unzip a zip file to specific path with a timeout canceled.
func UnzipWithTimeout(source, target string, timeout time.Duration) error {
	r, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer r.Close()

	// 检查是否超时
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	for _, file := range r.File {
		select {
		case <-timer.C:
			return fmt.Errorf("解压超时")
		default:
			// 创建对应的目录或文件
			path := filepath.Join(target, file.Name)

			if file.FileInfo().IsDir() {
				// 创建目录
				err := os.MkdirAll(path, file.Mode())
				if err != nil {
					return err
				}
			} else {
				// 创建文件
				writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
				if err != nil {
					return err
				}

				// 读取zip文件中的内容，写入到目标文件中
				rc, err := file.Open()
				if err != nil {
					writer.Close()
					return err
				}

				_, err = io.Copy(writer, rc)
				if err != nil {
					writer.Close()
					rc.Close()
					return err
				}

				writer.Close()
				rc.Close()
			}
		}
	}

	return nil
}
