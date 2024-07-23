package monitorServiceImpl

import (
	"baize/app/business/monitor/monitorModels"
	"baize/app/constant/sessionStatus"
	"baize/app/utils/cache"
	"baize/app/utils/session/sessionCache"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type UserOnlineService struct {
}

func NewUserOnlineService() *UserOnlineService {
	return new(UserOnlineService)
}

func (userOnlineService *UserOnlineService) SelectUserOnlineList(c *gin.Context, ol *monitorModels.SysUserOnlineDQL) (list []*monitorModels.SysUserOnline, total int64) {

	var cursor uint64 = 0
	keyAll := make([]string, 0, 16)
	for {
		keys, newCursor, err := cache.GetCache().Scan(c, cursor, sessionCache.SessionKey+":*", 10)
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

	list = make([]*monitorModels.SysUserOnline, 0, len(keyAll))
	for _, key := range keyAll {
		sk := strings.TrimPrefix(key, sessionCache.SessionKey+":")
		newSession := sessionCache.NewSession(sk)
		oui := new(monitorModels.SysUserOnline)
		oui.TokenId = sk
		oui.UserName = newSession.Get(c, sessionStatus.UserName)
		oui.Browser = newSession.Get(c, sessionStatus.Browser)
		oui.Ipaddr = newSession.Get(c, sessionStatus.IpAddr)
		oui.Os = newSession.Get(c, sessionStatus.Os)
		oui.LoginTime, _ = strconv.ParseInt(newSession.Get(c, sessionStatus.LoginTime), 10, 64)
		if ol.UserName != "" && !strings.Contains(oui.UserName, ol.UserName) {
			continue
		}
		if ol.Ipaddr != "" && !strings.Contains(oui.Ipaddr, ol.Ipaddr) {
			continue
		}
		list = append(list, oui)
	}

	total = int64(len(list))
	return
}

func (userOnlineService *UserOnlineService) ForceLogout(c *gin.Context, tokenId string) {
	cache.GetCache().Del(c, sessionCache.SessionKey+":"+tokenId)
}
