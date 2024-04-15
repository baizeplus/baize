package systemController

import (
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/business/system/systemService/systemServiceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type Notice struct {
	ns systemService.ISysNoticeService
}

func NewNotice(ns *systemServiceImpl.NoticeService) *Notice {
	return &Notice{ns: ns}
}

// NoticeList 消息通知列表
// @Summary 消息通知列表
// @Description 消息通知列表
// @Tags 消息通知
// @Param object query systemModels.NoticeDQL false "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ListData{rows=[]systemModels.SysNoticeVo}} "成功"
// @Router /system/notice/list [get]
func (nc *Notice) NoticeList(c *gin.Context) {
	n := new(systemModels.NoticeDQL)
	_ = c.ShouldBind(n)
	n.DataScope = baizeContext.GetDataScope(c, "sys_notice")
	if n.OrderBy == "" {
		n.OrderBy = "id"
		n.IsAsc = "desc"
	}
	list, count := nc.ns.SelectNoticeList(c, n)
	baizeContext.SuccessListData(c, list, count)

}

// NoticeGetInfo 根据id消息通知
// @Summary 根据id消息通知
// @Description 根据id消息通知
// @Tags 消息通知
// @Param id path string true "id"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=systemModels.SysNoticeVo} "成功"
// @Router /system/notice/{id}  [get]
func (nc *Notice) NoticeGetInfo(c *gin.Context) {
	id := baizeContext.ParamInt64(c, "id")
	if id == 0 {
		baizeContext.ParameterError(c)
		return
	}
	notice := nc.ns.SelectNoticeById(c, id)
	baizeContext.SuccessData(c, notice)
}

// NoticeAdd 新增消息通知
// @Summary 新增消息通知
// @Description 新增消息通知
// @Tags 消息通知
// @Param  object body systemModels.SysNoticeVo true "系统角色"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object} response.ResponseData
// @Router /system/notice [post]
func (nc *Notice) NoticeAdd(c *gin.Context) {
	na := new(systemModels.SysNoticeVo)
	if err := c.ShouldBindJSON(na); err != nil {
		baizeContext.ParameterError(c)
		return
	}

	na.SetCreateBy(baizeContext.GetUserId(c))
	nc.ns.InsertNotice(c, na)
	baizeContext.Success(c)
}

// NewMessage 未读消息通知数量
// @Summary 未读消息通知数量
// @Description 未读消息通知数量
// @Tags 消息通知
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object} response.ResponseData
// @Router /system/consumption/newMessage [get]
func (nc *Notice) NewMessage(c *gin.Context) {
	baizeContext.SuccessData(c, nc.ns.NewMessAge(c, baizeContext.GetUserId(c)))
}

// UserNoticeGetInfo 消费方获取消息根据ID
// @Summary 消费方获取消息根据ID
// @Description 消费方获取消息根据ID
// @Tags 消息通知
// @Param id path string true "id"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=systemModels.ConsumptionNoticeVo} "成功"
// @Router /system/consumption/{id}  [get]
func (nc *Notice) UserNoticeGetInfo(c *gin.Context) {
	id := baizeContext.ParamInt64(c, "id")
	if id == 0 {
		baizeContext.ParameterError(c)
		return
	}
	baizeContext.SuccessData(c, nc.ns.SelectConsumptionNoticeById(c, id))
}

// UserNoticeList 消费方获取消息列表
// @Summary 消费方获取消息列表
// @Description 消费方获取消息列表
// @Tags 消息通知
// @Param object query systemModels.ConsumptionNoticeDQL false "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object} response.ResponseData{data=response.ListData{rows=[]systemModels.ConsumptionNoticeVo}}
// @Router /system/consumption/userNoticeList [get]
func (nc *Notice) UserNoticeList(c *gin.Context) {
	n := new(systemModels.ConsumptionNoticeDQL)
	_ = c.ShouldBind(n)
	n.UserId = baizeContext.GetUserId(c)
	if n.OrderBy == "" {
		n.OrderBy = "sn.id"
		n.IsAsc = "desc"
	}
	list, total := nc.ns.SelectConsumptionNoticeList(c, n)
	baizeContext.SuccessListData(c, list, total)
}

// NoticeRead 消费方消息修改为已读
// @Summary 消费方消息修改为已读
// @Description 消费方消息修改为已读
// @Tags 消息通知
// @Param id path string true "id"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object} response.ResponseData
// @Router /system/consumption/noticeRead/{id} [put]
func (nc *Notice) NoticeRead(c *gin.Context) {
	id := baizeContext.ParamInt64(c, "id")
	if id == 0 {
		baizeContext.ParameterError(c)
		return
	}
	nc.ns.UpdateNoticeRead(c, id, baizeContext.GetUserId(c))
	baizeContext.Success(c)
}

// NoticeReadAll 消费方所有消息修改为已读
// @Summary 消费方所有消息修改为已读
// @Description 消费方所有消息修改为已读
// @Tags 消息通知
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object} response.ResponseData
// @Router /system/consumption/noticeReadAll [put]
func (nc *Notice) NoticeReadAll(c *gin.Context) {
	nc.ns.UpdateNoticeReadAll(c, baizeContext.GetUserId(c))
	baizeContext.Success(c)
}

// NoticeDelete 消费方消息删除消息
// @Summary 消费方消息删除消息
// @Description 消费方消息删除消息
// @Tags 消息通知
// @Param id path string true "id"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object} response.ResponseData
// @Router /system/consumption/noticeDelete/{ids} [delete]
func (nc *Notice) NoticeDelete(c *gin.Context) {
	ids := baizeContext.ParamInt64Array(c, "ids")
	if len(ids) == 0 {
		baizeContext.ParameterError(c)
		return
	}
	nc.ns.DeleteConsumptionNotice(c, ids, baizeContext.GetUserId(c))
	baizeContext.Success(c)
}
