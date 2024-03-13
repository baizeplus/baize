package IOFile

import (
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
	"time"
)

func GetRandomName(userId int64, extensionName string) string {
	uuid, _ := uuid.NewRandom()
	nameKey := gconv.String(userId) + "/" + time.Now().Format(time.DateOnly) + "/" + uuid.String() + extensionName
	return nameKey
}
func GetRandomPath(userId int64, name string) string {
	uuid, _ := uuid.NewRandom()
	nameKey := gconv.String(userId) + "/" + time.Now().Format(time.DateOnly) + "/" + uuid.String() + "/" + name
	return nameKey
}
