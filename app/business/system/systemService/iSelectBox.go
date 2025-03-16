package systemService

import (
	"baize/app/business/system/systemModels"
	"github.com/gin-gonic/gin"
)

type ISelectBoxService interface {
	SelectPermissionBox(c *gin.Context) (list []*systemModels.SelectPermission)
}
