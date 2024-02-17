package service

import (
	"baize/app/bzMonitor/models"
	"github.com/gin-gonic/gin"
)

type IUserOnlineService interface {
	SelectUserOnlineList(c *gin.Context) (list []*models.SysUserOnline, total *int64)
	ForceLogout(c *gin.Context, tokenId string)
}
