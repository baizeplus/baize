package systemService

import (
	"baize/app/business/system/systemModels"

	"github.com/gin-gonic/gin"
)

type ISysNoticeService interface {
	SelectNoticeList(c *gin.Context, notice *systemModels.NoticeDQL) (list []*systemModels.SysNoticeVo, total int64)
	SelectNoticeById(c *gin.Context, id string) *systemModels.SysNoticeVo
	InsertNotice(c *gin.Context, notice *systemModels.SysNoticeVo)
	NewMessAge(c *gin.Context, userId string) int64
	SelectConsumptionNoticeById(c *gin.Context, noticeId string) *systemModels.ConsumptionNoticeVo
	SelectConsumptionNoticeList(c *gin.Context, notice *systemModels.ConsumptionNoticeDQL) (list []*systemModels.ConsumptionNoticeVo, total int64)
	UpdateNoticeRead(c *gin.Context, noticeId, userId string)
	UpdateNoticeReadAll(c *gin.Context, userId string)
	DeleteConsumptionNotice(c *gin.Context, noticeId []string, userId string)
}
