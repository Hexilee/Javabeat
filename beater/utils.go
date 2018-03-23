package beater

import "strings"

const (
	HSperfPrefix = "hsperfdata_"
)

func GetUserFromHSperfDir(dirName string) (user string, ok bool) {
	if strings.HasPrefix(dirName, HSperfPrefix) {
		return strings.TrimPrefix(dirName, HSperfPrefix), true
	}
	return
}
