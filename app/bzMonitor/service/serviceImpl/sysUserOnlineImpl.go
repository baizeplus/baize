package serviceImpl

import (
	"baize/app/bzMonitor/models"
	"baize/app/constant/sessionStatus"
	"baize/app/datasource"
	"baize/app/utils/session/redis"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type UserOnlineService struct {
}

func NewUserOnlineService() *UserOnlineService {
	return new(UserOnlineService)
}

func (userOnlineService *UserOnlineService) SelectUserOnlineList(c *gin.Context) (list []*models.SysUserOnline, total *int64) {

	var cursor uint64 = 0
	keyAll := make([]string, 0, 16)
	for {
		keys, newCursor, err := datasource.RedisDb.Scan(c, cursor, redis.SessionKey+":*", 10).Result()
		if err != nil {
			panic(err)
		}
		// 处理从Scan中返回的键值对集合
		for _, key := range keys {
			keyAll = append(keyAll, key)
		}
		// 如果新游标为0，则意味着所有键都已经扫描完成
		if newCursor == 0 {
			break
		}
		// 更新游标，继续下一轮扫描
		cursor = newCursor
	}

	list = make([]*models.SysUserOnline, 0, len(keyAll))
	for _, key := range keyAll {
		sk := strings.TrimPrefix(key, redis.SessionKey+":")
		newSession := redis.NewSession(sk)
		oui := new(models.SysUserOnline)
		oui.TokenId = sk
		oui.UserName = newSession.Get(c, sessionStatus.UserName)
		oui.Browser = newSession.Get(c, sessionStatus.Browser)
		oui.Ipaddr = newSession.Get(c, sessionStatus.IpAddr)
		oui.Os = newSession.Get(c, sessionStatus.Os)
		oui.LoginTime, _ = strconv.ParseInt(newSession.Get(c, sessionStatus.LoginTime), 10, 64)
		list = append(list, oui)
	}

	i := int64(len(list))
	total = &i
	return
}

func (userOnlineService *UserOnlineService) ForceLogout(c *gin.Context, tokenId string) {
	datasource.RedisDb.Del(c, redis.SessionKey+":"+tokenId)
}
