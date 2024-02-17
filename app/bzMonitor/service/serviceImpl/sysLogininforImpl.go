package serviceImpl

import (
	"baize/app/bzMonitor/dao"
	"baize/app/bzMonitor/dao/daoImpl"
	"baize/app/bzMonitor/models"
	"baize/app/utils/snowflake"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
)

type LogininforService struct {
	data *sqly.DB
	ld   dao.ILogininforDao
}

func NewLogininforService(data *sqly.DB, ld *daoImpl.LogininforDao) *LogininforService {
	return &LogininforService{data: data, ld: ld}
}

func (ls *LogininforService) SelectLogininforList(c *gin.Context, logininfor *models.LogininforDQL) (list []*models.Logininfor, total *int64) {
	return ls.ld.SelectLogininforList(c, ls.data, logininfor)

}
func (ls *LogininforService) ExportLogininfor(c *gin.Context, logininfor *models.LogininforDQL) (data []byte) {
	//list, _ := ls.ld.SelectLogininforList(ls.data.GetSlaveDb(), logininfor)
	//return exceLize.SetRows(models.SysLogininforToRows(list))
	return nil
}

func (ls *LogininforService) InserLogininfor(c *gin.Context, loginUser *models.Logininfor) {
	loginUser.InfoId = snowflake.GenID()
	ls.ld.InserLogininfor(c, ls.data, loginUser)
}

func (ls *LogininforService) DeleteLogininforByIds(c *gin.Context, infoIds []int64) {
	ls.ld.DeleteLogininforByIds(c, ls.data, infoIds)

}

func (ls *LogininforService) CleanLogininfor(c *gin.Context) {
	ls.ld.CleanLogininfor(c, ls.data)

}
