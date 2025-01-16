package systemDaoImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemModels"
	"context"
	"database/sql"
	"errors"
	"github.com/baizeplus/sqly"
)

type sysNoticeDao struct {
	ms sqly.SqlyContext
}

func NewSysNoticeDao(ms sqly.SqlyContext) systemDao.ISysNoticeDao {
	return &sysNoticeDao{ms: ms}
}

func (s *sysNoticeDao) SelectNoticeList(ctx context.Context, notice *systemModels.NoticeDQL) (list []*systemModels.SysNoticeVo, total int64) {
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
	err := s.ms.NamedSelectPageContext(ctx, &list, &total, selectSql+whereSql, notice)
	if err != nil {
		panic(err)
	}
	return
}

func (s *sysNoticeDao) SelectNoticeById(ctx context.Context, id int64) *systemModels.SysNoticeVo {
	n := new(systemModels.SysNoticeVo)
	sqlStr := `select id,title,type,txt,create_by,create_time,create_name,dept_ids from sys_notice where id=?`
	err := s.ms.GetContext(ctx, n, sqlStr, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return n
}

func (s *sysNoticeDao) InsertNotice(ctx context.Context, notice *systemModels.SysNoticeVo) {
	insertSQL := `insert into sys_notice(id,title,type,txt,create_name,dept_ids,dept_id,create_by,create_time)
					values(:id,:title,:type,:txt,:create_name,:dept_ids,:dept_id,:create_by,:create_time)`
	_, err := s.ms.NamedExecContext(ctx, insertSQL, notice)
	if err != nil {
		panic(err)
	}

	return
}

func (s *sysNoticeDao) DeleteNoticeById(ctx context.Context, id int64) {

	_, err := s.ms.ExecContext(ctx, `delete from sys_notice where id = ? `, id)
	if err != nil {
		panic(err)
	}
}

func (s *sysNoticeDao) BatchSysNoticeUsers(ctx context.Context, notice []*systemModels.NoticeUser) {
	insertSQL := `insert into sys_notice_user(notice_id,user_id,status)
					values(:notice_id,:user_id,:status)`
	_, err := s.ms.NamedExecContext(ctx, insertSQL, notice)
	if err != nil {
		panic(err)
	}
	return
}

func (s *sysNoticeDao) SelectNewMessageCountByUserId(ctx context.Context, userId int64) int64 {
	count := int64(0)
	err := s.ms.GetContext(ctx, &count, `select count(*) from  sys_notice_user  where user_id=? and status='1' `, userId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return count
}

func (s *sysNoticeDao) SelectConsumptionNoticeById(ctx context.Context, userId, noticeId int64) *systemModels.ConsumptionNoticeVo {
	vo := new(systemModels.ConsumptionNoticeVo)
	err := s.ms.GetContext(ctx, vo, `select sn.id,sn.title,sn.txt,sn.create_name, sn.type,sn.create_time,snu.status from sys_notice sn left join sys_notice_user snu on sn.id = snu.notice_id where snu.user_id=? and snu.notice_id=?`,
		userId, noticeId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return vo
}

func (s *sysNoticeDao) SelectConsumptionNoticeList(ctx context.Context, notice *systemModels.ConsumptionNoticeDQL) (list []*systemModels.ConsumptionNoticeVo, total int64) {
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
	err := s.ms.NamedSelectPageContext(ctx, &list, &total, selectSql, notice)
	if err != nil {
		panic(err)
	}
	return list, total
}
func (s *sysNoticeDao) SelectNoticeStatusByNoticeIdAndUserId(ctx context.Context, noticeId, userId int64) int {
	count := 0
	err := s.ms.GetContext(ctx, &count, "SELECT EXISTS( SELECT 1 FROM sys_notice_user where user_id = ? and status='1' and notice_id =?)", userId, noticeId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return count
}
func (s *sysNoticeDao) SelectNoticeStatusByNoticeIdsAndUserId(ctx context.Context, noticeId []int64, userId int64) int {
	query, i, err := sqly.In("SELECT EXISTS( SELECT 1 FROM sys_notice_user where user_id = ? and status='1' and notice_id in (?)) ", userId, noticeId)
	count := 0
	err = s.ms.GetContext(ctx, &count, query, i...)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		panic(err)
	}
	return count
}
func (s *sysNoticeDao) UpdateNoticeRead(ctx context.Context, noticeId int64, userId int64) {

	_, err := s.ms.ExecContext(ctx, "update sys_notice_user set status = '2'  where user_id = ? and notice_id = ?", userId, noticeId)
	if err != nil {
		panic(err)
	}
}
func (s *sysNoticeDao) UpdateNoticeReadAll(ctx context.Context, userId int64) {
	_, err := s.ms.ExecContext(ctx, `update sys_notice_user set status = '2'  where user_id = ?`, userId)
	if err != nil {
		panic(err)
	}
}
func (s *sysNoticeDao) DeleteConsumptionNotice(ctx context.Context, noticeId []int64, userId int64) {
	query, i, err := sqly.In("delete from sys_notice_user where  user_id = ? and notice_id in(?) ", userId, noticeId)
	if err != nil {
		panic(err)
	}
	_, err = s.ms.ExecContext(ctx, query, i...)
	if err != nil {
		panic(err)
	}
}
