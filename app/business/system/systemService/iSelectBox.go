package systemService

import (
	"baize/app/baize"
	"baize/app/business/system/systemModels"
	"github.com/gin-gonic/gin"
)

type ISelectBoxService interface {
	SelectPermissionBox(c *gin.Context) (list []*systemModels.SelectPermission)
	SelectDeptBox(c *gin.Context, be *baize.BaseEntityDQL) (list []*systemModels.SelectDept)
}
