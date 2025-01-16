package cache

import (
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/monitor/monitorService"
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/setting"
	"context"
	"encoding/json"
	"go.uber.org/zap"
)

func NewRedisSubscribe(cache Cache, ss systemService.ISseService, js monitorService.IJobService) *RedisSubscribe {
	return &RedisSubscribe{cache: cache, ss: ss, js: js}
}

type RedisSubscribe struct {
	cache Cache
	ss    systemService.ISseService
	js    monitorService.IJobService
}

func (r *RedisSubscribe) Run() {
	if setting.Conf.Cluster {
		go r.SubscribeNotification()
		go r.SubscribeJob()
	}

}

func (r *RedisSubscribe) SubscribeNotification() {
	background := context.Background()
	subscribe := r.cache.Subscribe(background, "notification")
	defer subscribe.Close()
	ch := subscribe.Channel()
	for msg := range ch {
		var sse systemModels.Sse
		err := json.Unmarshal([]byte(msg.Payload), &sse)
		if err != nil {
			zap.L().Error("sse unmarshal error", zap.Error(err))
			continue
		}
		r.ss.SendNotification(background, &sse)
	}

}

func (r *RedisSubscribe) SubscribeJob() {
	background := context.Background()
	subscribe := r.cache.Subscribe(background, "job")
	defer subscribe.Close()
	ch := subscribe.Channel()
	for msg := range ch {
		var jb monitorModels.JobRedis
		err := json.Unmarshal([]byte(msg.Payload), &jb)
		if err != nil {
			zap.L().Error("sse unmarshal error", zap.Error(err))
			continue
		}
		if jb.Type == 0 {
			r.js.StartRunCron(background, &jb)
		} else {
			r.js.DeleteRunCron(background, &jb)
		}
	}

}
