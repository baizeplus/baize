package monitorController

import (
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/monitor/monitorService"
	"baize/app/business/monitor/monitorService/monitorServiceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type UserOnline struct {
	uos monitorService.IUserOnlineService
}

func NewUserOnline(uos *monitorServiceImpl.UserOnlineService) *UserOnline {
	return &UserOnline{
		uos: uos,
	}
}

// UserOnlineList 查询在线用户列表查询
// @Summary 查询在线用户列表查询
// @Description 查询在线用户列表查询
// @Tags 在线用户
// @Param  object query monitorModels.SysUserOnlineDQL true "查询信息"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData{data=response.ListData{Rows=[]monitorModels.SysUserOnline}}  "成功"
// @Router /monitor/online/list  [get]
func (uoc *UserOnline) UserOnlineList(c *gin.Context) {
	ol := new(monitorModels.SysUserOnlineDQL)
	_ = c.ShouldBind(ol)
	list, total := uoc.uos.SelectUserOnlineList(c, ol)
	baizeContext.SuccessListData(c, list, total)
}

// ForceLogout 强退在线用户列表
// @Summary 强退在线用户列表
// @Description 强退在线用户列表
// @Tags 在线用户
// @Param tokenId path string true "tokenId"
// @Security BearerAuth
// @Produce application/json
// @Success 200 {object}  response.ResponseData "成功"
// @Router /monitor/online/{tokenId}  [delete]
func (uoc *UserOnline) ForceLogout(c *gin.Context) {
	uoc.uos.ForceLogout(c, c.Param("tokenId"))
	baizeContext.Success(c)
}
