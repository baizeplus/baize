package systemServiceImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemDao/systemDaoImpl"
	"baize/app/business/system/systemModels"
	"baize/app/utils/snowflake"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
)

type NoticeService struct {
	data *sqly.DB
	nd   systemDao.ISysNoticeDao
	sud  systemDao.IUserDao
	sdd  systemDao.IDeptDao
}

func NewNoticeService(data *sqly.DB, nd *systemDaoImpl.SysNoticeDao,
	sud *systemDaoImpl.SysUserDao, sdd *systemDaoImpl.SysDeptDao) *NoticeService {
	return &NoticeService{data: data, nd: nd, sud: sud, sdd: sdd}
}

func (n *NoticeService) SelectNoticeList(c *gin.Context, notice *systemModels.NoticeDQL) (list []*systemModels.SysNoticeVo, total *int64) {
	return n.nd.SelectNoticeList(c, n.data, notice)
}

func (n *NoticeService) SelectNoticeById(c *gin.Context, id int64) *systemModels.SysNoticeVo {
	return n.nd.SelectNoticeById(c, n.data, id)
}

func (n *NoticeService) InsertNotice(c *gin.Context, notice *systemModels.SysNoticeVo) {
	noticeId := snowflake.GenID()
	notice.Id = noticeId
	n.nd.InsertNotice(c, n.data, notice)
}

func (n *NoticeService) IsNewMessAge(c *gin.Context, userId int64) string {
	if n.nd.SelectNewMessageCountByUserId(c, n.data, userId) > 0 {
		return "1"
	}
	return "2"
}

func (n *NoticeService) SelectConsumptionNoticeList(c *gin.Context, notice *systemModels.ConsumptionNoticeDQL) (list []*systemModels.ConsumptionNoticeVo, total *int64) {
	return n.nd.SelectConsumptionNoticeList(c, n.data, notice)
}
func (n *NoticeService) UpdateNoticeRead(c *gin.Context, noticeId, userId int64) {
	n.nd.UpdateNoticeRead(c, n.data, noticeId, userId)
}
func (n *NoticeService) UpdateNoticeReadAll(c *gin.Context, userId int64) {
	n.nd.UpdateNoticeReadAll(c, n.data, userId)
}
func (n *NoticeService) DeleteConsumptionNotice(c *gin.Context, noticeId, userId int64) {
	n.nd.DeleteConsumptionNotice(c, n.data, noticeId, userId)
}
