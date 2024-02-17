package controller

import (
	"baize/app/bzMonitor/service"
	"baize/app/bzMonitor/service/serviceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type UserOnline struct {
	uos service.IUserOnlineService
}

func NewUserOnline(uos *serviceImpl.UserOnlineService) *UserOnline {
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
