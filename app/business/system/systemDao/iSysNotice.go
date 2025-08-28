package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
)

type ISysNoticeDao interface {
	SelectNoticeList(ctx context.Context, notice *systemModels.NoticeDQL) (list []*systemModels.SysNoticeVo, total int64)
	SelectNoticeById(ctx context.Context, id string) *systemModels.SysNoticeVo
	InsertNotice(ctx context.Context, notice *systemModels.SysNoticeVo)
	DeleteNoticeById(ctx context.Context, id string)
	BatchSysNoticeUsers(ctx context.Context, notice []*systemModels.NoticeUser)
	SelectNewMessageCountByUserId(ctx context.Context, userId string) int64
	SelectConsumptionNoticeById(ctx context.Context, userId, noticeId string) *systemModels.ConsumptionNoticeVo
	SelectConsumptionNoticeList(ctx context.Context, notice *systemModels.ConsumptionNoticeDQL) (list []*systemModels.ConsumptionNoticeVo, total int64)
	SelectNoticeStatusByNoticeIdAndUserId(ctx context.Context, noticeId, userId string) int
	SelectNoticeStatusByNoticeIdsAndUserId(ctx context.Context, noticeId []string, userId string) int
	UpdateNoticeRead(ctx context.Context, noticeId string, userId string)
	UpdateNoticeReadAll(ctx context.Context, userId string)
	DeleteConsumptionNotice(ctx context.Context, noticeId []string, userId string)
}
