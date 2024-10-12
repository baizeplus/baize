package redisListener

import (
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/monitor/monitorService/monitorServiceImpl"
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
		go SubscribeJob()
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
		systemServiceImpl.GetSeeService().SendNotification(background, &sse)
	}

}

func SubscribeJob() {
	background := context.Background()
	subscribe := cache.RedisClient.Subscribe(background, "job")
	defer subscribe.Close()
	ch := subscribe.Channel()
	for msg := range ch {
		var jb monitorModels.JobRedis
		err := json.Unmarshal([]byte(msg.Payload), &jb)
		if err != nil {
			zap.L().Error("sse unmarshal error", zap.Error(err))
			continue
		}
		service := monitorServiceImpl.GetJobService()
		if jb.Type == 0 {
			service.StartRunCron(background, &jb)
		} else {
			service.DeleteRunCron(background, &jb)
		}
	}

}
