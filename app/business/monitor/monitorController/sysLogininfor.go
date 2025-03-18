package monitorController

import (
	"baize/app/business/monitor/monitorModels"
	"baize/app/business/monitor/monitorService"
	"baize/app/middlewares"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type Logininfor struct {
	ls monitorService.ILogininforService
}

func NewLogininfor(ls monitorService.ILogininforService) *Logininfor {
	return &Logininfor{ls: ls}
}

func (lc *Logininfor) PrivateRoutes(router *gin.RouterGroup) {
	logininfor := router.Group("/monitor/logininfor")
	logininfor.GET("/list", middlewares.HasPermission("system:monitor:logininfor"), lc.LogininforList)
	logininfor.GET("/export", middlewares.HasPermission("monitor:logininfor:export"), lc.LogininforExport)
	logininfor.DELETE("/:infoIds", middlewares.SetLog("登录日志", middlewares.Delete), middlewares.HasPermission("monitor:logininfor:remove"), lc.LogininforRemove)
	logininfor.DELETE("/clean", middlewares.SetLog("登录日志", middlewares.Clear), middlewares.HasPermission("monitor:logininfor:remove"), lc.LogininforClean)
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
