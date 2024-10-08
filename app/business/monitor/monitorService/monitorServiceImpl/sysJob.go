package monitorServiceImpl

import (
	"baize/app/business/monitor/monitorDao"
	"baize/app/business/monitor/monitorDao/monitorDaoImpl"
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/monitor/task"
	"baize/app/utils/cache"
	"baize/app/utils/snowflake"
	"context"
	"encoding/json"
	"fmt"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
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
			c := cron.New()
			_, err := c.AddFunc(jobVo.CronExpression, js.runFunction(jobVo.InvokeTarget))
			if err != nil {
				panic(err)
			}
			c.Start()
			js.cronMap[job.JobId] = c
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
	jobVo := js.SelectJobById(c, job.JobId)
	go js.runFunction(jobVo.InvokeTarget)
}
func (js *JobService) InsertJob(c *gin.Context, job *monitorModels.JobDML) {
	job.JobId = snowflake.GenID()
	js.jd.InsertJob(c, js.data, job)
}
func (js *JobService) UpdateJob(c *gin.Context, job *monitorModels.JobDML) bool {
	if job.Status == js.normal {
		//jobVo := js.SelectJobById(c, job.JobId)
		_, ok := js.funMap[job.InvokeTarget]
		if !ok {
			return false
		}
		_, ok = js.cronMap[job.JobId]
		if ok {
			c := cron.New()
			_, err := c.AddFunc(job.CronExpression, js.runFunction(job.InvokeTarget))
			if err != nil {
				panic(err)
			}
			c.Start()
			js.cronMap[job.JobId] = c
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
func (js *JobService) getQuartzCache(c context.Context, invokeTarget string) string {
	s := js.quartzKey + invokeTarget
	getString, _ := cache.GetCache().Get(c, s)
	if getString == "" {
		job := js.jd.SelectJobByInvokeTarget(c, js.data, invokeTarget)
		getString = job.JobParams
		if getString == "" {
			getString = "noParams"
		}
		cache.GetCache().Set(c, s, getString, -1)
	}
	return getString
}
func (js *JobService) runFunction(invokeTarget string) func() {
	return func() {
		ctx := context.Background()
		f := js.funMap[invokeTarget]
		quartzCache := js.getQuartzCache(ctx, invokeTarget)
		if quartzCache == "noParams" {
			f()
		} else {
			var fruits []string
			// 使用 json.Unmarshal 将 JSON 字符串解析到变量中
			err := json.Unmarshal([]byte(quartzCache), &fruits)
			if err != nil {
				fmt.Println("Error parsing JSON:", err)
				return
			}
			f(fruits...)
		}
	}
}
