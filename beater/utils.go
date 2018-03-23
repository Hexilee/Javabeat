package beater

import (
	"strings"
	"strconv"
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

func AToI32(pidStr string) (int32, error) {
	pid, err := strconv.Atoi(pidStr)
	return int32(pid), err
}
