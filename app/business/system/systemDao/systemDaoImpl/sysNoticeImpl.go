package systemDaoImpl

import (
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"errors"
	"github.com/baizeplus/sqly"
)

type SysNoticeDao struct {
}

func NewSysNoticeDao() *SysNoticeDao {
	return &SysNoticeDao{}
}

func (s *SysNoticeDao) SelectNoticeList(ctx context.Context, db sqly.SqlyContext, notice *systemModels.NoticeDQL) (list []*systemModels.SysNoticeVo, total int64) {
	selectSql := `select id,title,type,txt,create_by,create_time,create_name,dept_ids from sys_notice  `
	whereSql := ""
	if notice.NoticeTitle != "" {
		whereSql += " AND title like concat('%', :notice_title, '%')"
	}
	if notice.NoticeType != "" {
		whereSql += " AND type = :notice_type"
	}
	if notice.CreateBy != "" {
		whereSql += " AND create_name like concat('%', :create_by, '%')"
	}
	if whereSql != "" {
		whereSql = " where " + whereSql[4:]
	}
	err := db.NamedSelectPageContext(ctx, &list, &total, selectSql+whereSql, notice)
	if err != nil {
		panic(err)
	}
	return
}

func (s *SysNoticeDao) SelectNoticeById(ctx context.Context, db sqly.SqlyContext, id int64) *systemModels.SysNoticeVo {
	n := new(systemModels.SysNoticeVo)
	sqlStr := `select id,title,type,txt,create_by,create_time,create_name,dept_ids from sys_notice where id=?`
	err := db.GetContext(ctx, n, sqlStr, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return n
}

func (s *SysNoticeDao) InsertNotice(ctx context.Context, db sqly.SqlyContext, notice *systemModels.SysNoticeVo) {
	insertSQL := `insert into sys_notice(id,title,type,txt,create_name,dept_ids,dept_id,create_by,create_time)
					values(:id,:title,:type,:txt,:create_name,:dept_ids,:dept_id,:create_by,:create_time)`
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
	err := db.GetContext(ctx, &count, `select count(*) from  sys_notice_user  where user_id=? and status='1' `, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return count
}

func (s *SysNoticeDao) SelectConsumptionNoticeById(ctx context.Context, db sqly.SqlyContext, userId, noticeId int64) *systemModels.ConsumptionNoticeVo {
	vo := new(systemModels.ConsumptionNoticeVo)
	err := db.GetContext(ctx, vo, `select sn.id,sn.title,sn.txt,sn.create_name, sn.type,sn.create_time,snu.status from sys_notice sn left join sys_notice_user snu on sn.id = snu.notice_id where snu.user_id=? and snu.notice_id=?`,
		userId, noticeId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return vo
}

func (s *SysNoticeDao) SelectConsumptionNoticeList(ctx context.Context, db sqly.SqlyContext, notice *systemModels.ConsumptionNoticeDQL) (list []*systemModels.ConsumptionNoticeVo, total int64) {
	selectSql := `select sn.id,sn.title,sn.txt,sn.create_name,sn.create_time, sn.type,snu.status from sys_notice sn
left join sys_notice_user snu on sn.id = snu.notice_id
where snu.user_id=:user_id `
	if notice.Title != "" {
		selectSql += " AND sn.title like concat('%', :title, '%')"
	}
	if notice.Status != "" {
		selectSql += " AND snu.status=:status"
	}
	if notice.Type != "" {
		selectSql += " AND sn.type=:type"
	}
	err := db.NamedSelectPageContext(ctx, &list, &total, selectSql, notice)
	if err != nil {
		panic(err)
	}
	return list, total
}
func (s *SysNoticeDao) SelectNoticeStatusByNoticeIdAndUserId(ctx context.Context, db sqly.SqlyContext, noticeId, userId int64) int {
	count := 0
	err := db.GetContext(ctx, &count, "SELECT EXISTS( SELECT 1 FROM sys_notice_user where user_id = ? and status='1' and notice_id =?)", userId, noticeId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return count
}
func (s *SysNoticeDao) SelectNoticeStatusByNoticeIdsAndUserId(ctx context.Context, db sqly.SqlyContext, noticeId []int64, userId int64) int {
	query, i, err := sqly.In("SELECT EXISTS( SELECT 1 FROM sys_notice_user where user_id = ? and status='1' and notice_id in (?)) ", userId, noticeId)
	count := 0
	err = db.GetContext(ctx, &count, query, i...)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return count
}
func (s *SysNoticeDao) UpdateNoticeRead(ctx context.Context, db sqly.SqlyContext, noticeId int64, userId int64) {

	_, err := db.ExecContext(ctx, "update sys_notice_user set status = '2'  where user_id = ? and notice_id = ?", userId, noticeId)
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
func (s *SysNoticeDao) DeleteConsumptionNotice(ctx context.Context, db sqly.SqlyContext, noticeId []int64, userId int64) {
	query, i, err := sqly.In("delete from sys_notice_user where  user_id = ? and notice_id in(?) ", userId, noticeId)
	if err != nil {
		panic(err)
	}
	_, err = db.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}
