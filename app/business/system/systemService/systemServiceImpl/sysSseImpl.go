package systemServiceImpl

import (
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/constant/sessionStatus"
	"baize/app/datasource/cache"
	"baize/app/middlewares/session"
	"baize/app/setting"
	"baize/app/utils/arrayUtils"
	"baize/app/utils/baizeContext"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"sync"
)

type SseService struct {
	ChannelsMap map[string]chan *systemModels.SseType
	userMap     map[int64][]string
	mutex       sync.RWMutex
	cache       cache.Cache
}

func NewSseService(cache cache.Cache) systemService.ISseService {
	return &SseService{
		ChannelsMap: make(map[string]chan *systemModels.SseType),
		userMap:     make(map[int64][]string),
		cache:       cache,
	}
}

func (s *SseService) BuildNotificationChannel(c *gin.Context) {
	manager := session.NewManger(s.cache)
	sess, err := manager.Get(c, c.Param("token"))
	if err != nil {
		baizeContext.InvalidToken(c)
	}
	c.Set(sessionStatus.SessionKey, sess)
	closeNotify := c.Request.Context().Done()
	id := baizeContext.GetSession(c).Id()
	userId := baizeContext.GetUserId(c)
	s.mutex.Lock()
	var newChannel = make(chan *systemModels.SseType)
	s.ChannelsMap[id] = newChannel
	ids := s.userMap[userId]
	if ids == nil {
		ids = []string{id}
	} else if !arrayUtils.IsInArray(id, ids) {
		ids = append(ids, id)
	}
	s.userMap[userId] = ids
	s.mutex.Unlock()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				zap.L().Error("SSE断开链接错误", zap.Any("error", err))
			}
		}()
		<-closeNotify
		s.mutex.Lock()
		userIds := s.userMap[userId]
		for i, v := range userIds {
			if v == id {
				userIds = append(userIds[:i], userIds[i+1:]...)
				break
			}
		}
		if len(userIds) == 0 {
			delete(s.userMap, userId)
		} else {
			s.userMap[userId] = userIds
		}
		delete(s.ChannelsMap, id)
		s.mutex.Unlock()
		return
	}()

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-s.ChannelsMap[id]; ok {
			c.SSEvent(msg.Key, msg.Value)
			return true
		} else {
			return false
		}
	})

}
func (s *SseService) SendNotification(c context.Context, ss *systemModels.Sse) {
	if setting.Conf.Cluster && !ss.RedisPublish {
		ss.RedisPublish = true
		marshal, err := json.Marshal(ss)
		if err != nil {
			panic(err)
		}
		s.cache.Publish(c, "notification", marshal)
		return
	}

	for _, userId := range ss.UserIds {
		s.mutex.RLock()
		tokens := s.userMap[userId]
		s.mutex.RUnlock()
		if tokens == nil {
			return
		}
		for _, token := range tokens {
			channel, ok := s.ChannelsMap[token]
			if ok {
				channel <- ss.Sse
			}
		}
	}

}
