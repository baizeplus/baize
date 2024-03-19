package monitorDaoImpl

import (
	"baize/app/business/monitor/monitorModels"
	"context"
	"github.com/baizeplus/sqly"
)

type LogininforDao struct {
	selectSql string
}

func NewLogininforDao() *LogininforDao {
	return &LogininforDao{
		selectSql: `select info_id, user_name, ipaddr, login_location, browser, os, status, msg, login_time  from sys_logininfor`,
	}
}

func (ld *LogininforDao) InserLogininfor(ctx context.Context, db sqly.SqlyContext, logininfor *monitorModels.Logininfor) {
	_, err := db.NamedExecContext(ctx, "insert into sys_logininfor (info_id,user_name, status, ipaddr, login_location, browser, os, msg, login_time) values (:info_id,:user_name, :status, :ipaddr, :login_location, :browser, :os, :msg, sysdate())", logininfor)
	if err != nil {
		panic(err)
	}
	return
}
func (ld *LogininforDao) SelectLogininforList(ctx context.Context, db sqly.SqlyContext, logininfor *monitorModels.LogininforDQL) (list []*monitorModels.Logininfor, total *int64) {
	whereSql := ``
	if logininfor.IpAddr != "" {
		whereSql += " AND ipaddr like concat('%', :ipaddr, '%')"
	}
	if logininfor.Status != "" {
		whereSql += " AND  status = :status"
	}
	if logininfor.UserName != "" {
		whereSql += " AND user_name like concat('%', :userName, '%')"
	}

	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	logininfor.OrderBy = " info_id DESC"
	list = make([]*monitorModels.Logininfor, 0)
	total = new(int64)
	err := db.NamedSelectPageContext(ctx, &list, total, ld.selectSql+whereSql, logininfor, logininfor.ToPage())
	if err != nil {
		panic(err)
	}
	return

}
func (ld *LogininforDao) DeleteLogininforByIds(ctx context.Context, db sqly.SqlyContext, infoIds []int64) {
	query, i, err := sqly.In("delete from sys_logininfor where info_id in (?)", infoIds)
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}

func (ld *LogininforDao) CleanLogininfor(ctx context.Context, db sqly.SqlyContext) {
	_, err := db.ExecContext(ctx, "truncate table sys_logininfor")
	if err != nil {
		panic(err)
	}
}
