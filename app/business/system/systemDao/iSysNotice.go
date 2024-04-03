package systemDao

import (
	"baize/app/business/system/systemModels"
	"context"
	"github.com/baizeplus/sqly"
)

type ISysNoticeDao interface {
	SelectNoticeList(ctx context.Context, db sqly.SqlyContext, notice *systemModels.NoticeDQL) (list []*systemModels.SysNoticeVo, total *int64)
	SelectNoticeById(ctx context.Context, db sqly.SqlyContext, id int64) *systemModels.SysNoticeVo
	InsertNotice(ctx context.Context, db sqly.SqlyContext, notice *systemModels.SysNoticeVo)
	DeleteNoticeById(ctx context.Context, db sqly.SqlyContext, id int64)
	BatchSysNoticeUsers(ctx context.Context, db sqly.SqlyContext, notice []*systemModels.NoticeUser)
	SelectNewMessageCountByUserId(ctx context.Context, db sqly.SqlyContext, userId int64) int64
	SelectConsumptionNoticeList(ctx context.Context, db sqly.SqlyContext, notice *systemModels.ConsumptionNoticeDQL) (list []*systemModels.ConsumptionNoticeVo, total *int64)
	UpdateNoticeRead(ctx context.Context, db sqly.SqlyContext, noticeId, userId int64)
	UpdateNoticeReadAll(ctx context.Context, db sqly.SqlyContext, userId int64)
	DeleteConsumptionNotice(ctx context.Context, db sqly.SqlyContext, noticeId, userId int64)
}