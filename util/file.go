package util

import (
	"fmt"
	"os"
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
