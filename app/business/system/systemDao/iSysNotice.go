package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type ISysNoticeDao interface {
	SelectNoticeList(ctx context.Context, notice *systemModels.NoticeDQL) (list []*systemModels.SysNoticeVo, total int64)
	SelectNoticeById(ctx context.Context, id int64) *systemModels.SysNoticeVo
	InsertNotice(ctx context.Context, notice *systemModels.SysNoticeVo)
	DeleteNoticeById(ctx context.Context, id int64)
	BatchSysNoticeUsers(ctx context.Context, notice []*systemModels.NoticeUser)
	SelectNewMessageCountByUserId(ctx context.Context, userId int64) int64
	SelectConsumptionNoticeById(ctx context.Context, userId, noticeId int64) *systemModels.ConsumptionNoticeVo
	SelectConsumptionNoticeList(ctx context.Context, notice *systemModels.ConsumptionNoticeDQL) (list []*systemModels.ConsumptionNoticeVo, total int64)
	SelectNoticeStatusByNoticeIdAndUserId(ctx context.Context, noticeId, userId int64) int
	SelectNoticeStatusByNoticeIdsAndUserId(ctx context.Context, noticeId []int64, userId int64) int
	UpdateNoticeRead(ctx context.Context, noticeId int64, userId int64)
	UpdateNoticeReadAll(ctx context.Context, userId int64)
	DeleteConsumptionNotice(ctx context.Context, noticeId []int64, userId int64)
}
