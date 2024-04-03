package systemDaoImpl

import (
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"github.com/baizeplus/sqly"
)

type SysNoticeDao struct {
}

func NewSysNoticeDao() *SysNoticeDao {
	return &SysNoticeDao{}
}

func (s *SysNoticeDao) SelectNoticeList(ctx context.Context, db sqly.SqlyContext, notice *systemModels.NoticeDQL) (list []*systemModels.SysNoticeVo, total *int64) {
	//selectSql := `select id,case user_ids when''then'1' else'2'end  receiver,publish_time,title,txt,concat(su.first_name,' ',su.last_name) create_name,sn.status from sys_notice sn left join sys_user su on sn.update_by=su.user_id  `

	selectSql := `select id,publish_time,title,concat(su.first_name,' ',su.last_name) create_name,sn.status from sys_notice sn left join sys_user su on sn.update_by=su.user_id  `
	whereSql := ""
	if notice.Title != "" {
		whereSql += " AND title like concat('%', :title, '%')"
	}
	if notice.TimeRangeStart != "" {
		whereSql += " AND publish_time > :time_start"
	}
	if notice.TimeRangeEnd != "" {
		whereSql += " AND publish_time < concat( :time_end , ' 23::59::59::999')"
	}
	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	list = make([]*systemModels.SysNoticeVo, 0, 16)
	total = new(int64)
	err := db.NamedSelectPageContext(ctx, &list, total, selectSql+whereSql, notice, notice.ToPage())
	if err != nil {
		panic(err)
	}
	return
}

func (s *SysNoticeDao) SelectNoticeById(ctx context.Context, db sqly.SqlyContext, id int64) *systemModels.SysNoticeVo {
	n := new(systemModels.SysNoticeVo)
	sqlStr := `select id,publish_time,title,status from sys_notice where id=?`
	err := db.GetContext(ctx, n, sqlStr, id)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return n
}

func (s *SysNoticeDao) InsertNotice(ctx context.Context, db sqly.SqlyContext, notice *systemModels.SysNoticeVo) {
	insertSQL := `insert into sys_notice(id,publish_time,title,status,create_by,create_time,update_by,update_time)
					values(:id,:publish_time,:title,:status,:create_by,:create_time,:update_by,:update_time)`
	_, err := db.NamedExecContext(ctx, insertSQL, notice)
	if err != nil {
		panic(err)
	}

	return
}

func (s *SysNoticeDao) DeleteNoticeById(ctx context.Context, db sqly.SqlyContext, id int64) {

	_, err := db.ExecContext(ctx, `delete from sys_notice where id = ? `, id)
	if err != nil {
		panic(err)
	}
}

func (s *SysNoticeDao) BatchSysNoticeUsers(ctx context.Context, db sqly.SqlyContext, notice []*systemModels.NoticeUser) {
	insertSQL := `insert into sys_notice_user(notice_id,user_id,status)
					values(:notice_id,:user_id,:status)`
	_, err := db.NamedExecContext(ctx, insertSQL, notice)
	if err != nil {
		panic(err)
	}
	return
}

func (s *SysNoticeDao) SelectNewMessageCountByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) int64 {
	count := int64(0)
	err := db.GetContext(ctx, &count, `select count(*) from sys_notice n left join sys_notice_user nu on n.id=nu.notice_id where nu.user_id=? and nu.status='1' and  n.status='0' AND publish_time< now()`, userId)
	if err != nil {
		panic(err)
	}
	return count
}

func (s *SysNoticeDao) SelectConsumptionNoticeList(ctx context.Context, db sqly.SqlyContext, notice *systemModels.ConsumptionNoticeDQL) (list []*systemModels.ConsumptionNoticeVo, total *int64) {
	selectSql := `select sn.id,sn.publish_time,sn.title, snu.status from sys_notice sn
left join sys_notice_user snu on sn.id = snu.notice_id
where sn.status='0' and snu.user_id=:user_id AND sn.publish_time< now()`
	if notice.Title != "" {
		selectSql += " AND sn.title like concat('%', :title, '%')"
	}
	if notice.Unread != "" {
		selectSql += " AND snu.status=:status"
	}
	list = make([]*systemModels.ConsumptionNoticeVo, 0)
	total = new(int64)
	err := db.NamedSelectPageContext(ctx, &list, total, selectSql, notice, notice.ToPage())
	if err != nil {
		panic(err)
	}
	return list, total
}
func (s *SysNoticeDao) UpdateNoticeRead(ctx context.Context, db sqly.SqlyContext, noticeId, userId int64) {
	_, err := db.ExecContext(ctx, `update sys_notice_user set status = '2'  where notice_id=? and user_id = ?`, noticeId, userId)
	if err != nil {
		panic(err)
	}
}
func (s *SysNoticeDao) UpdateNoticeReadAll(ctx context.Context, db sqly.SqlyContext, userId int64) {
	_, err := db.ExecContext(ctx, `update sys_notice_user set status = '2'  where user_id = ?`, userId)
	if err != nil {
		panic(err)
	}
}
func (s *SysNoticeDao) DeleteConsumptionNotice(ctx context.Context, db sqly.SqlyContext, noticeId, userId int64) {
	_, err := db.ExecContext(ctx, "delete from sys_notice_user where notice_id=? and user_id = ? ", noticeId, userId)
	if err != nil {
		panic(err)
	}
}
