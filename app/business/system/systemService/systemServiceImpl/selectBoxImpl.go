package systemServiceImpl

import (
	"baize/app/baize"
	"baize/app/business/system/systemDao"
	"baize/app/business/system/systemModels"
	"baize/app/business/system/systemService"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type SelectService struct {
	pd systemDao.IPermissionDao
	dd systemDao.IDeptDao
}

func NewSelectService(pd systemDao.IPermissionDao, dd systemDao.IDeptDao) systemService.ISelectBoxService {
	return &SelectService{pd: pd, dd: dd}
}

func (cs *SelectService) SelectPermissionBox(c *gin.Context) (list []*systemModels.SelectPermission) {
	return cs.pd.SelectPermissionListSelectBoxByPerm(c, baizeContext.GetPermission(c))
}

func (cs *SelectService) SelectDeptBox(c *gin.Context, be *baize.BaseEntityDQL) (list []*systemModels.SelectDept) {
	return cs.dd.SelectDeptListSelectBox(c, be)
}
