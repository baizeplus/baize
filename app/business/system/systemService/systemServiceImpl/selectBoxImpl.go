package systemServiceImpl

import (
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type SelectService struct {
	pd systemDao.IPermissionDao
}

func NewSelectService(pd systemDao.IPermissionDao) systemService.ISelectBoxService {
	return &SelectService{pd: pd}
}

func (cs *SelectService) SelectPermissionBox(c *gin.Context) (list []*systemModels.SelectPermission) {
	return cs.pd.SelectPermissionListSelectBoxByPerm(c, baizeContext.GetPermission(c))
}
