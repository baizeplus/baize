package monitorDaoImpl

import (
	"baize/app/business/monitor/monitorDao"
	"baize/app/business/monitor/monitorModels"
	"context"
	"database/sql"
	"errors"
	"github.com/baizeplus/sqly"
)

func NewJobDao(ms sqly.SqlyContext) monitorDao.IJobDao {
	return &JobDao{
		ms:           ms,
		selectSql:    "select job_id, job_name,job_params, invoke_target, cron_expression,status, create_by, create_time  from sys_job",
		selectLogSql: "select job_log_id,job_id, job_name,job_params, invoke_target,status, create_time ,cost_time  from sys_job_log",
	}
}

type JobDao struct {
	ms           sqly.SqlyContext
	selectSql    string
	selectLogSql string
}

func (jd *JobDao) SelectJobList(ctx context.Context, job *monitorModels.JobDQL) (list []*monitorModels.JobVo, total int64) {
	whereSql := ``
	if job.JobName != "" {
		whereSql += " AND job_name like concat('%', :job_name, '%')"
	}
	if job.Status != "" {
		whereSql += " AND status = :status"
	}
	if job.InvokeTarget != "" {
		whereSql += " AND invoke_target like concat('%', :invoke_target, '%')"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}

	err := jd.ms.NamedSelectPageContext(ctx, &list, &total, jd.selectSql+whereSql, job)
	if err != nil {
		panic(err)
	}
	return
}
func (jd *JobDao) SelectJobAll(ctx context.Context) (list []*monitorModels.JobVo) {
	list = make([]*monitorModels.JobVo, 0)
	err := jd.ms.SelectContext(ctx, &list, jd.selectSql)
	if err != nil {
		panic(err)
	}
	return
}
func (jd *JobDao) SelectJobIdAndNameAll(ctx context.Context) (list []*monitorModels.JobIdAndName) {
	list = make([]*monitorModels.JobIdAndName, 0)
	err := jd.ms.SelectContext(ctx, &list, "select job_id, job_name from sys_job")
	if err != nil {
		panic(err)
	}
	return
}
func (jd *JobDao) SelectJobById(ctx context.Context, id int64) (job *monitorModels.JobVo) {
	job = new(monitorModels.JobVo)
	err := jd.ms.GetContext(ctx, job, jd.selectSql+" where job_id = ?", id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}
func (jd *JobDao) SelectJobByInvokeTarget(ctx context.Context, invokeTarget string) (job *monitorModels.JobVo) {
	job = new(monitorModels.JobVo)
	err := jd.ms.GetContext(ctx, job, jd.selectSql+" where invoke_target = ?", invokeTarget)
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}
func (jd *JobDao) DeleteJobById(ctx context.Context, id int64) {
	_, err := jd.ms.ExecContext(ctx, "delete from sys_job where job_id = ", id)
	if err != nil {
		panic(err)
	}
	return
}
func (jd *JobDao) UpdateJob(ctx context.Context, job *monitorModels.JobDML) {
	updateSQL := `update sys_job set invoke_target = :invoke_target,update_time = :update_time , update_by = :update_by`

	if job.CronExpression != "" {
		updateSQL += ",cron_expression = :cron_expression"
	}

	if job.JobName != nil {
		updateSQL += ",job_name = :job_name"
	}

	if job.JobParams != nil {
		updateSQL += ",job_params = :job_params"
	}

	if job.Status != "" {
		updateSQL += ",status = :status"
	}

	updateSQL += " where job_id = :job_id"

	_, err := jd.ms.NamedExecContext(ctx, updateSQL, job)
	if err != nil {
		panic(err)
	}
	return
}
func (jd *JobDao) InsertJob(ctx context.Context, job *monitorModels.JobDML) {
	insertSQL := `insert into sys_job(job_id,job_name,job_params,invoke_target,cron_expression,status,create_by,create_time,update_by,update_time)
					values(:job_id,:job_name,:job_params,:invoke_target,:cron_expression,:status,:create_by,:create_time,:update_by,:update_time)`
	_, err := jd.ms.NamedExecContext(ctx, insertSQL, job)
	if err != nil {
		panic(err)
	}
	return
}
func (jd *JobDao) DeleteJobByIds(ctx context.Context, ids []int64) {
	query, i, err := sqly.In("delete from sys_job where job_id in (?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = jd.ms.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}

func (jd *JobDao) InsertJobLog(ctx context.Context, job *monitorModels.JobLog) {
	insertSQL := `insert into sys_job_log(job_log_id,job_id, job_name,job_params, invoke_target,status, create_time ,cost_time)
					values(:job_log_id,:job_id, :job_name,:job_params, :invoke_target,:status, :create_time ,:cost_time)`
	_, err := jd.ms.NamedExecContext(ctx, insertSQL, job)
	if err != nil {
		panic(err)
	}
	return
}
func (jd *JobDao) SelectJobLogList(ctx context.Context, job *monitorModels.JobLogDql) (list []*monitorModels.JobLog, total int64) {
	whereSql := ``
	if job.JobId != 0 {
		whereSql += " AND job_id = :job_id"
	}
	if job.Status != "" {
		whereSql += " AND status = :status"
	}
	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}

	err := jd.ms.NamedSelectPageContext(ctx, &list, &total, jd.selectLogSql+whereSql, job)
	if err != nil {
		panic(err)
	}
	return
}

func (jd *JobDao) SelectJobLogById(ctx context.Context, id int64) (vo *monitorModels.JobLog) {
	vo = new(monitorModels.JobLog)
	err := jd.ms.GetContext(ctx, vo, jd.selectLogSql+" where job_log_id = ?", id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	} else if err != nil {
		panic(err)
	}
	return vo
}
