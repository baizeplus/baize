package systemController

import (
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/business/system/systemService/systemServiceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type NoticeController struct {
	ns systemService.ISysNoticeService
}

func NewNoticeController(ns *systemServiceImpl.NoticeService) *NoticeController {
	return &NoticeController{ns: ns}
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
func (nc *NoticeController) NoticeList(c *gin.Context) {

	n := new(systemModels.NoticeDQL)
	_ = c.ShouldBind(n)

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
func (nc *NoticeController) NoticeGetInfo(c *gin.Context) {
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
func (nc *NoticeController) NoticeAdd(c *gin.Context) {
	na := new(systemModels.SysNoticeVo)
	if err := c.ShouldBindJSON(na); err != nil {
		baizeContext.ParameterError(c)
		return
	}

	na.SetCreateBy(baizeContext.GetUserId(c))
	nc.ns.InsertNotice(c, na)
	baizeContext.Success(c)
}

// IsNewMessage 消费方是否有新消息 1有 2没有
// @Summary 消费方是否有新消息
// @Description 消费方是否有新消息
// @Tags 消息通知
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object} response.ResponseData
// @Router /system/consumption/isNewMessage [get]
func (nc *NoticeController) IsNewMessage(c *gin.Context) {
	baizeContext.SuccessData(c, nc.ns.IsNewMessAge(c, baizeContext.GetUserId(c)))
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
func (nc *NoticeController) UserNoticeList(c *gin.Context) {

	n := new(systemModels.ConsumptionNoticeDQL)
	_ = c.ShouldBind(n)
	n.UserId = baizeContext.GetUserId(c)
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
func (nc *NoticeController) NoticeRead(c *gin.Context) {
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
func (nc *NoticeController) NoticeReadAll(c *gin.Context) {
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
// @Router /system/consumption/noticeDelete/{id} [delete]
func (nc *NoticeController) NoticeDelete(c *gin.Context) {
	id := baizeContext.ParamInt64(c, "id")
	if id == 0 {
		baizeContext.ParameterError(c)
		return
	}
	nc.ns.DeleteConsumptionNotice(c, id, baizeContext.GetUserId(c))
	baizeContext.Success(c)
}
