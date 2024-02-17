package controller

import (
	"baize/app/bzMonitor/models"
	"baize/app/bzMonitor/service"
	"baize/app/bzMonitor/service/serviceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
)

type Logininfor struct {
	ls service.ILogininforService
}

func NewLogininfor(ls *serviceImpl.LogininforService) *Logininfor {
	return &Logininfor{ls: ls}
}

func (lc *Logininfor) LogininforList(c *gin.Context) {
	loginfor := new(models.LogininforDQL)
	c.ShouldBind(loginfor)
	list, count := lc.ls.SelectLogininforList(c, loginfor)
	baizeContext.SuccessListData(c, list, count)

}

func (lc *Logininfor) LogininforExport(c *gin.Context) {

}

func (lc *Logininfor) LogininforRemove(c *gin.Context) {

	lc.ls.DeleteLogininforByIds(c, baizeContext.ParamInt64Array(c, "infoIds"))
	baizeContext.Success(c)
}

func (lc *Logininfor) LogininforClean(c *gin.Context) {
	lc.ls.CleanLogininfor(c)
	baizeContext.Success(c)

}
