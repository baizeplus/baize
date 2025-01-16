package monitorServiceImpl

import (
	"baize/app/business/monitor/monitorDao"
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/monitor/monitorService"
	"baize/app/business/monitor/task"
	"baize/app/datasource/cache"
	"baize/app/setting"
	"baize/app/utils/baizeContext"
	"baize/app/utils/snowflake"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"time"
)

type JobService struct {
	cache     cache.Cache
	jd        monitorDao.IJobDao
	funMap    map[string]func(...string)
	cronMap   map[int64]*cron.Cron
	normal    string
	pause     string
	quartzKey string
}

func NewJobService(cache cache.Cache, jd monitorDao.IJobDao) monitorService.IJobService {
	funMap := make(map[string]func(...string))
	funMap["NoParams"] = task.NoParams
	funMap["Params"] = task.Params
	return &JobService{cache: cache, jd: jd, funMap: funMap, cronMap: make(map[int64]*cron.Cron),
		normal: "0", pause: "1", quartzKey: "quartz:"}
}

func (js *JobService) SelectJobList(c *gin.Context, job *monitorModels.JobDQL) (list []*monitorModels.JobVo, total int64) {
	list, total = js.jd.SelectJobList(c, job)
	return
}
func (js *JobService) SelectJobById(c *gin.Context, id int64) (job *monitorModels.JobVo) {
	job = js.jd.SelectJobById(c, id)
	if job != nil {
		schedule, err := cron.ParseStandard(job.CronExpression)
		if err != nil {
			panic(err)
		}
		now := time.Now()
		next := schedule.Next(now)
		job.NextValidTime = &next
	}
	return
}

func (js *JobService) DeleteJobByIds(c *gin.Context, jobIds []int64) {
	for _, id := range jobIds {
		m := new(monitorModels.JobRedis)
		m.Id = id
		js.DeleteRunCron(c, m)
	}
	js.jd.DeleteJobByIds(c, jobIds)
}

func (js *JobService) ChangeStatus(c *gin.Context, job *monitorModels.JobStatus) {
	m := new(monitorModels.JobRedis)
	m.Id = job.JobId
	js.DeleteRunCron(c, m)
	if job.Status == js.normal {
		js.StartRunCron(c, m)
	}
	d := new(monitorModels.JobDML)
	d.JobId = job.JobId
	d.Status = job.Status
	d.SetUpdateBy(baizeContext.GetUserId(c))
	js.jd.UpdateJob(c, d)

}
func (js *JobService) Run(c *gin.Context, job *monitorModels.JobStatus) {
	vo := js.jd.SelectJobById(c, job.JobId)
	jr := new(monitorModels.JobRun)
	jr.JobId = vo.JobId
	jr.JobName = vo.JobName
	jr.JobParams = vo.JobParams
	jr.InvokeTarget = vo.InvokeTarget
	jr.CronExpression = vo.CronExpression
	go js.runFunction(jr)()
}
func (js *JobService) InsertJob(c *gin.Context, job *monitorModels.JobDML) {
	job.JobId = snowflake.GenID()
	js.jd.InsertJob(c, job)
	if job.Status == js.normal {
		m := new(monitorModels.JobRedis)
		m.Id = job.JobId
		js.StartRunCron(c, m)
	}
}
func (js *JobService) FunIsExist(invokeTarget string) bool {
	_, ok := js.funMap[invokeTarget]
	return ok
}
func (js *JobService) GetFunList() []string {
	keys := make([]string, 0, len(js.funMap))
	for k := range js.funMap {
		keys = append(keys, k)
	}
	return keys
}
func (js *JobService) UpdateJob(c *gin.Context, job *monitorModels.JobDML) {
	m := new(monitorModels.JobRedis)
	m.Id = job.JobId
	js.DeleteRunCron(c, m)
	if job.Status == js.normal {
		js.StartRunCron(c, m)
	}
	js.jd.UpdateJob(c, job)
}

func (js *JobService) InitJobRun() {
	ctx := context.Background()
	list := js.jd.SelectJobAll(ctx)
	for _, vo := range list {
		_, ok := js.funMap[vo.InvokeTarget]
		if !ok {
			panic("目标方法未找到")
		}
		if vo.Status == js.normal {
			jr := new(monitorModels.JobRun)
			jr.JobId = vo.JobId
			jr.JobName = vo.JobName
			jr.JobParams = vo.JobParams
			jr.InvokeTarget = vo.InvokeTarget
			jr.CronExpression = vo.CronExpression
			cr := cron.New()
			_, err := cr.AddFunc(vo.CronExpression, js.runFunction(jr))
			if err != nil {
				panic(err)
			}
			cr.Start()
			js.cronMap[vo.JobId] = cr
		}
	}
}

func (js *JobService) runFunction(job *monitorModels.JobRun) func() {
	return func() {
		ctx := context.Background()

		if setting.Conf.Cluster {
			schedule, err := cron.ParseStandard(job.CronExpression)
			if err != nil {
				panic(err)
			}
			now := time.Now()
			nextTime := schedule.Next(now).Unix()
			success := js.cache.SetNX(ctx, fmt.Sprintf("scheduled_task_lock:%s:%d", job.JobName, nextTime), "locked", time.Minute)
			if !success {
				return
			}
		}
		m := new(monitorModels.JobLog)
		m.JobLogId = snowflake.GenID()
		m.JobId = job.JobId
		m.InvokeTarget = job.InvokeTarget
		m.JobName = job.JobName
		m.InvokeTarget = job.InvokeTarget
		m.CreateTime = time.Now()
		m.JobParams = job.JobParams
		defer func() {
			m.CostTime = int64(time.Since(m.CreateTime))
			if e := recover(); e != nil {
				m.Status = "1"
				m.ExceptionInfo = e.(error).Error()
			} else {
				m.Status = "0"
			}
			js.jd.InsertJobLog(ctx, m)
		}()
		f := js.funMap[job.InvokeTarget]
		f(job.JobParams.Data...)
	}
}
func (js *JobService) StartRunCron(c context.Context, jo *monitorModels.JobRedis) {
	if setting.Conf.Cluster && !jo.RedisPublish {
		jo.RedisPublish = true
		jo.Type = 0
		marshal, err := json.Marshal(jo)
		if err != nil {
			panic(err)
		}
		js.cache.Publish(c, "job", marshal)
		return
	}
	cr, ok := js.cronMap[jo.Id]
	if ok {
		cr.Stop()
		delete(js.cronMap, jo.Id)
	}
	vo := js.jd.SelectJobById(c, jo.Id)
	cr = cron.New()
	jr := new(monitorModels.JobRun)
	jr.JobId = vo.JobId
	jr.JobName = vo.JobName
	jr.JobParams = vo.JobParams
	jr.InvokeTarget = vo.InvokeTarget
	jr.CronExpression = vo.CronExpression
	_, err := cr.AddFunc(jr.CronExpression, js.runFunction(jr))
	if err != nil {
		panic(err)
	}
	cr.Start()
	js.cronMap[jo.Id] = cr
}
func (js *JobService) DeleteRunCron(c context.Context, jo *monitorModels.JobRedis) {
	if setting.Conf.Cluster && !jo.RedisPublish {
		jo.RedisPublish = true
		jo.Type = 1
		marshal, err := json.Marshal(jo)
		if err != nil {
			panic(err)
		}
		js.cache.Publish(c, "job", marshal)
		return
	}
	cr, ok := js.cronMap[jo.Id]
	if ok {
		cr.Stop()
		delete(js.cronMap, jo.Id)
	}
}

func (js *JobService) SelectJobLogList(c *gin.Context, job *monitorModels.JobLogDql) (list []*monitorModels.JobLog, total int64) {
	return js.jd.SelectJobLogList(c, job)
}

func (js *JobService) SelectJobLogById(c *gin.Context, id int64) (vo *monitorModels.JobLog) {
	return js.jd.SelectJobLogById(c, id)
}

func (js *JobService) SelectJobIdAndNameAll(c *gin.Context) (list []*monitorModels.JobIdAndName) {
	return js.jd.SelectJobIdAndNameAll(c)
}
