package monitorServiceImpl

import (
	"baize/app/business/monitor/monitorDao"
	"baize/app/business/monitor/monitorDao/monitorDaoImpl"
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/monitor/task"
	"baize/app/setting"
	"baize/app/utils/cache"
	"baize/app/utils/snowflake"
	"context"
	"fmt"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"time"
)

type JobService struct {
	data      *sqly.DB
	jd        monitorDao.IJobDao
	funMap    map[string]func(...string)
	cronMap   map[int64]*cron.Cron
	normal    string
	pause     string
	quartzKey string
}

func NewJobService(data *sqly.DB, jd *monitorDaoImpl.JobDao) *JobService {
	funMap := make(map[string]func(...string))
	funMap["NoParams"] = task.NoParams
	funMap["Params"] = task.Params
	return &JobService{data: data, jd: jd, funMap: funMap, cronMap: make(map[int64]*cron.Cron),
		normal: "0", pause: "1", quartzKey: "quartz:"}
}

func (js *JobService) SelectJobList(c *gin.Context, job *monitorModels.JobDQL) (list []*monitorModels.JobVo, total int64) {
	list, total = js.jd.SelectJobList(c, js.data, job)
	return
}
func (js *JobService) SelectJobById(c *gin.Context, id int64) (job *monitorModels.JobVo) {
	job = js.jd.SelectJobById(c, js.data, id)
	return
}

func (js *JobService) DeleteJobByIds(c *gin.Context, jobIds []int64) {
	for _, id := range jobIds {
		cr, ok := js.cronMap[id]
		if ok {
			cr.Stop()
			delete(js.cronMap, id)
		}
	}
	js.jd.DeleteJobByIds(c, js.data, jobIds)
}

func (js *JobService) ChangeStatus(c *gin.Context, job *monitorModels.JobDML) bool {

	if job.Status == js.normal {
		jobVo := js.SelectJobById(c, job.JobId)
		_, ok := js.funMap[jobVo.InvokeTarget]
		if !ok {
			return false
		}
		_, ok = js.cronMap[job.JobId]
		if ok {
			cr := cron.New()
			j := js.SelectJobById(c, job.JobId)
			_, err := cr.AddFunc(jobVo.CronExpression, js.runFunction(j))
			if err != nil {
				panic(err)
			}
			cr.Start()
			js.cronMap[job.JobId] = cr
		}
	} else if job.Status == js.pause {
		cr, ok := js.cronMap[job.JobId]
		if ok {
			cr.Stop()
			delete(js.cronMap, job.JobId)
		}

	}
	js.jd.UpdateJob(c, js.data, job)
	return true
}
func (js *JobService) Run(c *gin.Context, job *monitorModels.JobVo) {
	j := js.SelectJobById(c, job.JobId)
	go js.runFunction(j)()
}
func (js *JobService) InsertJob(c *gin.Context, job *monitorModels.JobDML) {
	job.JobId = snowflake.GenID()
	js.jd.InsertJob(c, js.data, job)
	if job.Status == js.normal {
		cr := cron.New()
		j := js.SelectJobById(c, job.JobId)
		_, err := cr.AddFunc(job.CronExpression, js.runFunction(j))
		if err != nil {
			panic(err)
		}
		cr.Start()
		js.cronMap[job.JobId] = cr

	}
}
func (js *JobService) UpdateJob(c *gin.Context, job *monitorModels.JobDML) bool {
	_, ok := js.funMap[job.InvokeTarget]
	if !ok {
		return false
	}
	if job.Status == js.normal {
		cr, ok := js.cronMap[job.JobId]
		if ok {
			cr.Stop()
			delete(js.cronMap, job.JobId)
		}
		cr = cron.New()
		j := js.SelectJobById(c, job.JobId)
		_, err := cr.AddFunc(job.CronExpression, js.runFunction(j))
		if err != nil {
			panic(err)
		}
		cr.Start()
		js.cronMap[job.JobId] = cr

	} else if job.Status == js.pause {
		cr, ok := js.cronMap[job.JobId]
		if ok {
			cr.Stop()
			delete(js.cronMap, job.JobId)
		}

	}

	js.jd.UpdateJob(c, js.data, job)
	return true
}

func (js *JobService) runFunction(job *monitorModels.JobVo) func() {
	return func() {
		ctx := context.Background()
		if setting.Conf.Cluster {
			success, err := cache.RedisClient.SetNX(ctx, "scheduled_task_lock", "locked", time.Minute).Result()
			if err != nil {
				fmt.Println("Error connecting to Redis:", err)
				return
			}
			if success {
				defer func() {
					cache.RedisClient.Del(ctx, "scheduled_task_lock")
				}()
			} else {
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
			js.jd.InsertJobLog(ctx, js.data, m)
		}()
		f := js.funMap[job.InvokeTarget]
		f(job.JobParams.Data...)
	}
}
