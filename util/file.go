package util

import (
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"os"
	"path/filepath"
)

const (
	MarkerContent = "If you see this message, it means the file has been copied successfully."
)

// GetFileInfo returns file or dir info. Returns an error if file or dir is not exist.
func GetFileInfo(path string) (os.FileInfo, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("not exist file or path: %v", path)
	}
	return info, nil
}

func MkdirAll(path string) {
	os.MkdirAll(path, 0755)
}

// AddMarker adds marker file to validate whether backup dir is successful. If you use CopyFile to back up, it won't work.
func AddMarker(dst string) error {
	filePath := filepath.Join(dst, "marker.txt")

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("create file failed: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(MarkerContent)
	if err != nil {
		return fmt.Errorf("writing marker content failed: %v", err)
	}

	return nil
}

// ValidateMarker validates marker file before you start recovering backup dir. If you back up file instead of file, it won't work.
func ValidateMarker(dst string) error {
	filePath := filepath.Join(dst, "marker.txt")

	content, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("open file failed: %v", err)
	}
	if string(content) != MarkerContent {
		return fmt.Errorf("marker is incomplete")
	}
	return nil
}

// CopyFile copy src file to dest file. If os is linux, it will use rsync command instead of io.Copy.
func CopyFile(srcPath, desPath string) error {
	return fileutil.CopyFile(srcPath, desPath)
}

func IsDir(path string) bool {
	return fileutil.IsDir(path)
}

func IsExist(path string) bool {
	return fileutil.IsExist(path)
}
