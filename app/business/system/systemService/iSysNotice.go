package systemService

import (
	"baize/app/business/system/systemModels"
	"github.com/gin-gonic/gin"
)

type ISysNoticeService interface {
	SelectNoticeList(c *gin.Context, notice *systemModels.NoticeDQL) (list []*systemModels.SysNoticeVo, total int64)
	SelectNoticeById(c *gin.Context, id int64) *systemModels.SysNoticeVo
	InsertNotice(c *gin.Context, notice *systemModels.SysNoticeVo)
	NewMessAge(c *gin.Context, userId int64) int64
	SelectConsumptionNoticeById(c *gin.Context, noticeId int64) *systemModels.ConsumptionNoticeVo
	SelectConsumptionNoticeList(c *gin.Context, notice *systemModels.ConsumptionNoticeDQL) (list []*systemModels.ConsumptionNoticeVo, total int64)
	UpdateNoticeRead(c *gin.Context, noticeId, userId int64)
	UpdateNoticeReadAll(c *gin.Context, userId int64)
	DeleteConsumptionNotice(c *gin.Context, noticeId []int64, userId int64)
}
