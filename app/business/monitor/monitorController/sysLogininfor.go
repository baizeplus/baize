package monitorController

import (
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/monitor/monitorService"
	"baize/app/business/monitor/monitorService/monitorServiceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type Logininfor struct {
	ls monitorService.ILogininforService
}

func NewLogininfor(ls *monitorServiceImpl.LogininforService) *Logininfor {
	return &Logininfor{ls: ls}
}

func (lc *Logininfor) LogininforList(c *gin.Context) {
	loginfor := new(monitorModels.LogininforDQL)
	_ = c.ShouldBind(loginfor)
	if loginfor.OrderBy == "" {
		loginfor.OrderBy = "info_id"
		loginfor.IsAsc = "desc"
	}
	list, count := lc.ls.SelectLogininforList(c, loginfor)
	baizeContext.SuccessListData(c, list, count)

}

func (lc *Logininfor) LogininforExport(c *gin.Context) {
	loginfor := new(monitorModels.LogininforDQL)
	_ = c.ShouldBind(loginfor)
	if loginfor.OrderBy == "" {
		loginfor.OrderBy = "info_id"
		loginfor.IsAsc = "desc"
	}
	data := lc.ls.ExportLogininfor(c, loginfor)
	baizeContext.DataPackageExcel(c, data)
}

func (lc *Logininfor) LogininforRemove(c *gin.Context) {
	lc.ls.DeleteLogininforByIds(c, baizeContext.ParamInt64Array(c, "infoIds"))
	baizeContext.Success(c)
}

func (lc *Logininfor) LogininforClean(c *gin.Context) {
	lc.ls.CleanLogininfor(c)
	baizeContext.Success(c)

}
