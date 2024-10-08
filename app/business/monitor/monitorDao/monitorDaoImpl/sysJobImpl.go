package monitorDaoImpl

import (
	"baize/app/business/monitor/monitorModels"
	"context"
	"database/sql"
	"errors"
	"github.com/baizeplus/sqly"
)

func NewJobDao() *JobDao {
	return &JobDao{
		selectSql: "select job_id, job_name,job_params, job_group, invoke_target, cron_expression,status, create_by, create_time, remark  from sys_job",
	}
}

type JobDao struct {
	selectSql string
}

func (jd *JobDao) SelectJobList(ctx context.Context, db sqly.SqlyContext, job *monitorModels.JobDQL) (list []*monitorModels.JobVo, total int64) {
	whereSql := ``
	if job.JobName != "" {
		whereSql += " AND job_name like concat('%', :job_name, '%')"
	}
	if job.JobGroup != "" {
		whereSql += " AND job_group = :job_group"
	}
	if job.Status != nil {
		whereSql += " AND status = :status"
	}
	if job.InvokeTarget != "" {
		whereSql += " AND invoke_target like concat('%', :invoke_target, '%')"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}

	err := db.NamedSelectPageContext(ctx, &list, &total, jd.selectSql+whereSql, job)
	if err != nil {
		panic(err)
	}
	return
}
func (jd *JobDao) SelectJobAll(ctx context.Context, db sqly.SqlyContext) (list []*monitorModels.JobVo) {
	list = make([]*monitorModels.JobVo, 0)
	err := db.SelectContext(ctx, &list, jd.selectSql)
	if err != nil {
		panic(err)
	}
	return
}
func (jd *JobDao) SelectJobById(ctx context.Context, db sqly.SqlyContext, id int64) (job *monitorModels.JobVo) {
	job = new(monitorModels.JobVo)
	err := db.GetContext(ctx, job, jd.selectSql+" where job_id = ?", id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}
func (jd *JobDao) SelectJobByInvokeTarget(ctx context.Context, db sqly.SqlyContext, invokeTarget string) (job *monitorModels.JobVo) {
	job = new(monitorModels.JobVo)
	err := db.GetContext(ctx, job, jd.selectSql+" where invoke_target = ?", invokeTarget)
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	} else if err != nil {
		panic(err)
	}
	return
}
func (jd *JobDao) DeleteJobById(ctx context.Context, db sqly.SqlyContext, id int64) {
	_, err := db.ExecContext(ctx, "delete from sys_job where job_id = ", id)
	if err != nil {
		panic(err)
	}
	return
}
func (jd *JobDao) UpdateJob(ctx context.Context, db sqly.SqlyContext, job *monitorModels.JobDML) {
	updateSQL := `update sys_job set invoke_target = :invoke_target,update_time = now() , update_by = :update_by`

	if job.CronExpression != "" {
		updateSQL += ",cron_expression = :cron_expression"
	}

	if job.JobName != nil {
		updateSQL += ",job_name = :job_name"
	}

	if job.JobGroup != nil {
		updateSQL += ",job_group = :job_group"
	}

	if job.JobParams != "" {
		updateSQL += ",job_params = :job_params"
	}

	if job.Status != "" {
		updateSQL += ",status = :status"
	}
	if job.Remark != nil {
		updateSQL += ",remark = :remark"
	}

	updateSQL += " where job_id = :job_id"

	_, err := db.NamedExecContext(ctx, updateSQL, job)
	if err != nil {
		panic(err)
	}
	return
}
func (jd *JobDao) InsertJob(ctx context.Context, db sqly.SqlyContext, job *monitorModels.JobDML) {
	insertSQL := `insert into sys_job(job_id,job_name,job_group,job_params,invoke_target,cron_expression,status,create_by,create_time,update_by,update_time ,remark)
					values(:job_id,:job_name,:job_group,:job_params,:invoke_target,:cron_expression,:status,:create_by,now(),:update_by,now() ,:remark)`
	_, err := db.NamedExecContext(ctx, insertSQL, job)
	if err != nil {
		panic(err)
	}
	return
}
func (jd *JobDao) DeleteJobByIds(ctx context.Context, db sqly.SqlyContext, ids []int64) {
	query, i, err := sqly.In("delete from sys_job where job_id in (?)", ids)
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}
