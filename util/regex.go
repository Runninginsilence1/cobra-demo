package util

import (
	"regexp"
)

const (
	FilePathPattern = "^[a-zA-Z0-9_-]+$"
)

// Match returns whether string match regex pattern. Returns error if pattern is not correct.
func Match(pattern, s string) (bool, error) {
	match, err := regexp.MatchString(pattern, s)
	if err != nil {
		return false, err
	}
	return match, nil
}

// IsValidFileName returns whether string match standard linux file name(not path name).
func IsValidFileName(fileName string) bool {
	match, _ := Match(FilePathPattern, fileName)
	return match
}
