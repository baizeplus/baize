package IOFile

import (
	"baize/app/utils/stringUtils"
	"strconv"
	"time"
)

func GetRandomName(userId int64, extensionName string) string {
	nameKey := strconv.FormatInt(userId, 10) + "/" + time.Now().Format(time.DateOnly) + "/" + stringUtils.GetUUID() + extensionName
	return nameKey
}
func GetRandomPath(userId int64, name string) string {
	nameKey := strconv.FormatInt(userId, 10) + "/" + time.Now().Format(time.DateOnly) + "/" + stringUtils.GetUUID() + "/" + name
	return nameKey
}
