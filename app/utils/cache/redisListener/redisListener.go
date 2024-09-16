package redisListener

import (
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService/systemServiceImpl"
	"baize/app/setting"
	"baize/app/utils/cache"
	"context"
	"encoding/json"
	"go.uber.org/zap"
)

func StartRedisListener() {
	if setting.Conf.Cluster {
		go SubscribeNotification()
	}

}
func SubscribeNotification() {
	background := context.Background()
	subscribe := cache.RedisClient.Subscribe(background, "notification")
	defer subscribe.Close()
	ch := subscribe.Channel()
	for msg := range ch {
		var sse systemModels.Sse
		err := json.Unmarshal([]byte(msg.Payload), &sse)
		if err != nil {
			zap.L().Error("sse unmarshal error", zap.Error(err))
			continue
		}
		systemServiceImpl.SseServiceInstance.SendNotification(background, &sse)
	}

}
