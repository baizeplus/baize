package monitorController

import (
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

func (uoc *UserOnline) UserOnlineList(c *gin.Context) {
	list, total := uoc.uos.SelectUserOnlineList(c)
	baizeContext.SuccessListData(c, list, total)
}
func (uoc *UserOnline) ForceLogout(c *gin.Context) {
	uoc.uos.ForceLogout(c, c.Param("tokenId"))
	baizeContext.Success(c)
}
