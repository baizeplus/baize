package IOFile

import (
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/google/uuid"
	"time"
)

func GetTenantRandomName(userId int64, extensionName string) string {
	uuid, _ := uuid.NewRandom()
	nameKey := gconv.String(userId) + "/" + time.Now().Format(time.DateOnly) + "/" + uuid.String() + extensionName
	return nameKey
}
