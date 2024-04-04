package systemServiceImpl

import (
	"baize/app/business/system/systemModels"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
	"io"
	"sync"
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
	ids := s.userMap[userId]
	if ids == nil {
		ids = []string{id}
	} else {
		ids = append(ids, id)
	}
	s.userMap[userId] = ids
	s.mutex.Unlock()
	go func() {
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
			c.SSEvent(msg.Key, msg.Value)
			return true
		} else {
			return false
		}
	})
}
func (s *SseService) SendNotification(userId int64, ss *systemModels.Sse) {
	s.mutex.RLock()
	userIds := s.userMap[userId]
	s.mutex.RUnlock()
	if userIds == nil {
		return
	}

	for _, id := range userIds {
		channel := s.channelsMap[id]
		channel <- ss
	}
}
