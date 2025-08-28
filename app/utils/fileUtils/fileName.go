package fileUtils

import (
	"baize/app/utils/stringUtils"
	"time"
)

func GetRandomName(userId string, extensionName string) string {
	return userId + "/" + time.Now().Format(time.DateOnly) + "/" + stringUtils.GetUUID() + extensionName
}
func GetRandomPath(userId string, name string) string {
	return userId + "/" + time.Now().Format(time.DateOnly) + "/" + stringUtils.GetUUID() + "/" + name
}
