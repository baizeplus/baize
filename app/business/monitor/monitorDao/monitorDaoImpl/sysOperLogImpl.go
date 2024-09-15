package monitorDaoImpl

import (
	"baize/app/business/monitor/monitorModels"
	"context"
	"database/sql"
	"errors"
	"github.com/baizeplus/sqly"
)

type OperLogDao struct {
	selectSql string
}

func NewOperLog() *OperLogDao {
	return &OperLogDao{
		selectSql: ` select oper_id, title, business_type, method, request_method, user_id, oper_name, oper_url, oper_ip, oper_param, json_result, status,  oper_time,cost_time from sys_oper_log`,
	}
}

func (operLogDao *OperLogDao) InsertOperLog(ctx context.Context, db sqly.SqlyContext, operLog *monitorModels.SysOperLog) {
	_, err := db.NamedExecContext(ctx, "insert into sys_oper_log(oper_id,title, business_type, method, request_method, user_id, oper_name, oper_url, oper_ip, oper_param, json_result, status,  oper_time,cost_time)  values (:oper_id,:title, :business_type, :method, :request_method, :user_id, :oper_name, :oper_url, :oper_ip,  :oper_param, :json_result, :status, now(),:cost_time)", operLog)
	if err != nil {
		panic(err)
	}
	return
}
func (operLogDao *OperLogDao) SelectOperLogList(ctx context.Context, db sqly.SqlyContext, openLog *monitorModels.SysOperLogDQL) (list []*monitorModels.SysOperLog, total int64) {
	whereSql := ``
	if openLog.Title != "" {
		whereSql += " AND title like concat('%', :title, '%')"
	}
	if openLog.BusinessType != nil {
		whereSql += " AND business_type = :business_type"
	}
	if openLog.Status != "" {
		whereSql += " AND status = :status"
	}
	if openLog.OperName != "" {
		whereSql += " AND oper_name like concat('%', :oper_name, '%')"
	}
	if openLog.BeginTime != "" {
		whereSql += " AND date_format(oper_time,'%y%m%d') >= :begin_time"
	}
	if openLog.EndTime != "" {
		whereSql += " AND date_format(oper_time,'%y%m%d') >= :end_time"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	err := db.NamedSelectPageContext(ctx, &list, &total, operLogDao.selectSql+whereSql, openLog)
	if err != nil {
		panic(err)
	}
	return
}
func (operLogDao *OperLogDao) SelectOperLogListAll(ctx context.Context, db sqly.SqlyContext, openLog *monitorModels.SysOperLogDQL) (list []*monitorModels.SysOperLog) {
	whereSql := ``
	if openLog.Title != "" {
		whereSql += " AND title like concat('%', :title, '%')"
	}
	if openLog.BusinessType != nil {
		whereSql += " AND business_type = :business_type"
	}
	if openLog.Status != "" {
		whereSql += " AND status = :status"
	}
	if openLog.OperName != "" {
		whereSql += " AND oper_name like concat('%', :oper_name, '%')"
	}
	if openLog.BeginTime != "" {
		whereSql += " AND date_format(oper_time,'%y%m%d') >= :begin_time"
	}
	if openLog.EndTime != "" {
		whereSql += " AND date_format(oper_time,'%y%m%d') >= :end_time"
	}
	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	list = make([]*monitorModels.SysOperLog, 0)
	err := db.NamedSelectContext(ctx, &list, operLogDao.selectSql+whereSql, openLog)
	if err != nil {
		panic(err)
	}
	return
}

func (operLogDao *OperLogDao) DeleteOperLogByIds(ctx context.Context, db sqly.SqlyContext, operIds []int64) {
	query, i, err := sqly.In("delete from sys_oper_log where oper_id in (?)", operIds)
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}

func (operLogDao *OperLogDao) SelectOperLogById(ctx context.Context, db sqly.SqlyContext, operId int64) (operLog *monitorModels.SysOperLog) {
	whereSql := `  where oper_id = ?`
	operLog = new(monitorModels.SysOperLog)
	err := db.GetContext(ctx, operLog, operLogDao.selectSql+whereSql, operId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return
}

func (operLogDao *OperLogDao) CleanOperLog(ctx context.Context, db sqly.SqlyContext) {
	_, err := db.ExecContext(ctx, "truncate table sys_oper_log")
	if err != nil {
		panic(err)
	}
}
