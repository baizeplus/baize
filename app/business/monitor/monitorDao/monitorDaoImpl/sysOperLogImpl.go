package monitorDaoImpl

import (
	"baize/app/business/monitor/monitorDao"
	"baize/app/business/monitor/monitorModels"
	"context"
	"database/sql"
	"errors"
	"github.com/baizeplus/sqly"
)

type OperLogDao struct {
	ms        sqly.SqlyContext
	selectSql string
}

func NewOperLog(ms sqly.SqlyContext) monitorDao.IOperLog {
	return &OperLogDao{
		ms:        ms,
		selectSql: ` select oper_id, title, business_type, method, request_method, user_id, oper_name, oper_url, oper_ip, oper_param, json_result, status,  oper_time,cost_time from sys_oper_log`,
	}
}

func (operLogDao *OperLogDao) InsertOperLog(ctx context.Context, operLog *monitorModels.SysOperLog) {
	_, err := operLogDao.ms.NamedExecContext(ctx, "insert into sys_oper_log(oper_id,title, business_type, method, request_method, user_id, oper_name, oper_url, oper_ip, oper_param, json_result, status,  oper_time,cost_time)  values (:oper_id,:title, :business_type, :method, :request_method, :user_id, :oper_name, :oper_url, :oper_ip,  :oper_param, :json_result, :status, :oper_time,:cost_time)", operLog)
	if err != nil {
		panic(err)
	}
	return
}
func (operLogDao *OperLogDao) SelectOperLogList(ctx context.Context, openLog *monitorModels.SysOperLogDQL) (list []*monitorModels.SysOperLog, total int64) {
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
	err := operLogDao.ms.NamedSelectPageContext(ctx, &list, &total, operLogDao.selectSql+whereSql, openLog)
	if err != nil {
		panic(err)
	}
	return
}
func (operLogDao *OperLogDao) SelectOperLogListAll(ctx context.Context, openLog *monitorModels.SysOperLogDQL) (list []*monitorModels.SysOperLog) {
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
	err := operLogDao.ms.NamedSelectContext(ctx, &list, operLogDao.selectSql+whereSql, openLog)
	if err != nil {
		panic(err)
	}
	return
}

func (operLogDao *OperLogDao) DeleteOperLogByIds(ctx context.Context, operIds []string) {
	query, i, err := sqly.In("delete from sys_oper_log where oper_id in (?)", operIds)
	if err != nil {
		panic(err)
	}
	_, err = operLogDao.ms.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}

func (operLogDao *OperLogDao) SelectOperLogById(ctx context.Context, operId string) (operLog *monitorModels.SysOperLog) {
	whereSql := `  where oper_id = ?`
	operLog = new(monitorModels.SysOperLog)
	err := operLogDao.ms.GetContext(ctx, operLog, operLogDao.selectSql+whereSql, operId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return
}

func (operLogDao *OperLogDao) CleanOperLog(ctx context.Context) {
	_, err := operLogDao.ms.ExecContext(ctx, "truncate table sys_oper_log")
	if err != nil {
		panic(err)
	}
}
