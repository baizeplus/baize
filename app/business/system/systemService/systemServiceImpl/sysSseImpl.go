package systemServiceImpl

import (
	"baize/app/business/system/systemModels"
	"baize/app/utils/arrayUtils"
	"baize/app/utils/baizeContext"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"sync"
	"time"
)

type SseService struct {
	channelsMap map[string]chan *systemModels.Sse
	userMap     map[int64][]string
	mutex       sync.RWMutex
}

func NewSseService() *SseService {

	return &SseService{
		channelsMap: make(map[string]chan *systemModels.Sse),
		userMap:     make(map[int64][]string),
	}
}

func (s *SseService) BuildNotificationChannel(c *gin.Context) {
	closeNotify := c.Request.Context().Done()
	id := baizeContext.GetSession(c).Id()
	userId := baizeContext.GetUserId(c)
	s.mutex.Lock()
	var newChannel = make(chan *systemModels.Sse)
	s.channelsMap[id] = newChannel
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
		delete(s.channelsMap, id)
		s.mutex.Unlock()
		return
	}()
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-s.channelsMap[id]; ok {
			fmt.Println("b", time.Now().UnixNano())
			c.SSEvent(msg.Key, msg.Value)
			return true
		} else {
			return false
		}
	})
}
func (s *SseService) SendNotification(userIds []int64, ss *systemModels.Sse) {
	for _, userId := range userIds {
		s.mutex.RLock()
		tokens := s.userMap[userId]
		s.mutex.RUnlock()
		if tokens == nil {
			return
		}
		for _, token := range tokens {
			channel, ok := s.channelsMap[token]
			if ok {
				channel <- ss
			}
		}
	}

}
