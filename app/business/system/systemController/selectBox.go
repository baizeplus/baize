package systemController

import (
	"baize/app/business/system/systemService"
	"github.com/gin-gonic/gin"
)

type SelectBox struct {
	sbs systemService.ISelectBoxService
}

func NewSelectBox(sbs systemService.ISelectBoxService) *SelectBox {
	return &SelectBox{sbs: sbs}
}
func (s *SelectBox) PrivateRoutes(router *gin.RouterGroup) {
	sb := router.Group("/system/selectBox")
	sb.GET("/permission", s.SelectPermission)
}

func (s *SelectBox) SelectPermission(c *gin.Context) {
	s.sbs.SelectPermissionBox(c)
}
