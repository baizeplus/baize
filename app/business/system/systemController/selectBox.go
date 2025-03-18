package systemController

import (
	"baize/app/business/system/systemService"
	"baize/app/middlewares"
	"baize/app/utils/baizeContext"
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
	sb.GET("/permission", middlewares.HasPermissions([]string{"system:permission:add", "system:permission:edit"}), s.SelectPermission)
}

func (s *SelectBox) SelectPermission(c *gin.Context) {
	baizeContext.SuccessData(c, s.sbs.SelectPermissionBox(c))
}
