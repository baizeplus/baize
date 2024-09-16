package systemServiceImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemDao/systemDaoImpl"
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/utils/baizeContext"
	"baize/app/utils/snowflake"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
	"strconv"
)

type NoticeService struct {
	data *sqly.DB
	nd   systemDao.ISysNoticeDao
	sud  systemDao.IUserDao
	ss   systemService.ISseService
	sss  *systemModels.SseType
}

func NewNoticeService(data *sqly.DB, nd *systemDaoImpl.SysNoticeDao,
	sud *systemDaoImpl.SysUserDao, ss *SseService) *NoticeService {
	return &NoticeService{data: data, nd: nd, sud: sud, ss: ss, sss: &systemModels.SseType{Key: "notice", Value: "1"}}
}

func (n *NoticeService) SelectNoticeList(c *gin.Context, notice *systemModels.NoticeDQL) (list []*systemModels.SysNoticeVo, total int64) {
	return n.nd.SelectNoticeList(c, n.data, notice)
}

func (n *NoticeService) SelectNoticeById(c *gin.Context, id int64) *systemModels.SysNoticeVo {
	return n.nd.SelectNoticeById(c, n.data, id)
}

func (n *NoticeService) InsertNotice(c *gin.Context, notice *systemModels.SysNoticeVo) {
	noticeId := snowflake.GenID()
	notice.Id = noticeId
	ids := notice.DeptIds
	notice.SetCreateBy(baizeContext.GetUserId(c))
	notice.DeptId = baizeContext.GetDeptId(c)
	notice.CreateName = baizeContext.GetUserName(c)
	deptIds := make([]int64, 0, len(ids.Data))
	for _, id := range ids.Data {
		i, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			continue
		}
		deptIds = append(deptIds, i)
	}
	userIds := n.sud.SelectUserIdsByDeptIds(c, n.data, deptIds)
	tx, err := n.data.Beginx()
	if err != nil {
		panic(err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else {
			err = tx.Commit()
			go n.ss.SendNotification(c, &systemModels.Sse{UserIds: userIds, Sse: n.sss})
		}
	}()
	n.nd.InsertNotice(c, tx, notice)
	if len(userIds) == 0 {
		return
	}
	users := make([]*systemModels.NoticeUser, 0, len(userIds))
	for _, id := range userIds {
		s := new(systemModels.NoticeUser)
		s.NoticeId = noticeId
		s.UserId = id
		s.Status = "1"
		users = append(users, s)
	}
	n.nd.BatchSysNoticeUsers(c, tx, users)

}

func (n *NoticeService) NewMessAge(c *gin.Context, userId int64) int64 {
	return n.nd.SelectNewMessageCountByUserId(c, n.data, userId)
}

func (n *NoticeService) SelectConsumptionNoticeById(c *gin.Context, noticeId int64) *systemModels.ConsumptionNoticeVo {
	userId := baizeContext.GetUserId(c)
	status := n.nd.SelectNoticeStatusByNoticeIdAndUserId(c, n.data, noticeId, userId)
	if status == 1 {
		n.nd.UpdateNoticeRead(c, n.data, noticeId, userId)
		go n.ss.SendNotification(c, &systemModels.Sse{UserIds: []int64{userId}, Sse: n.sss})
	}
	return n.nd.SelectConsumptionNoticeById(c, n.data, userId, noticeId)
}

func (n *NoticeService) SelectConsumptionNoticeList(c *gin.Context, notice *systemModels.ConsumptionNoticeDQL) (list []*systemModels.ConsumptionNoticeVo, total int64) {
	return n.nd.SelectConsumptionNoticeList(c, n.data, notice)
}
func (n *NoticeService) UpdateNoticeRead(c *gin.Context, noticeId, userId int64) {
	status := n.nd.SelectNoticeStatusByNoticeIdAndUserId(c, n.data, noticeId, userId)
	if status == 0 {
		return
	}
	n.nd.UpdateNoticeRead(c, n.data, noticeId, userId)
	go n.ss.SendNotification(c, &systemModels.Sse{UserIds: []int64{userId}, Sse: n.sss})
}
func (n *NoticeService) UpdateNoticeReadAll(c *gin.Context, userId int64) {
	n.nd.UpdateNoticeReadAll(c, n.data, userId)
	go n.ss.SendNotification(c, &systemModels.Sse{UserIds: []int64{userId}, Sse: n.sss})
}
func (n *NoticeService) DeleteConsumptionNotice(c *gin.Context, noticeId []int64, userId int64) {
	i := n.nd.SelectNoticeStatusByNoticeIdsAndUserId(c, n.data, noticeId, userId)
	n.nd.DeleteConsumptionNotice(c, n.data, noticeId, userId)
	if i == 1 {
		go n.ss.SendNotification(c, &systemModels.Sse{UserIds: []int64{userId}, Sse: n.sss})
	}
}
