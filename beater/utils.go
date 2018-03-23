package beater

import (
	"strings"
)

const (
	HSperfPrefix           = "hsperfdata_"
	ErrNotImplementedError = "not implemented yet"
)

func GetUserFromHSperfDir(dirName string) (user string, ok bool) {
	if strings.HasPrefix(dirName, HSperfPrefix) {
		return strings.TrimPrefix(dirName, HSperfPrefix), true
	}
	return
}

func IsOSError(err error) bool {
	// Skip ErrNotImplementedError
	if err != nil && err.Error() != ErrNotImplementedError {
		return true
	}
	return false
}
