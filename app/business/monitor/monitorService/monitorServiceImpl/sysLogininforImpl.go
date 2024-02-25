package monitorServiceImpl

import (
	"baize/app/business/monitor/monitorDao"
	"baize/app/business/monitor/monitorDao/monitorDaoImpl"
	"baize/app/business/monitor/monitorModels"
	"baize/app/utils/snowflake"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
)

type LogininforService struct {
	data *sqly.DB
	ld   monitorDao.ILogininforDao
}

func NewLogininforService(data *sqly.DB, ld *monitorDaoImpl.LogininforDao) *LogininforService {
	return &LogininforService{data: data, ld: ld}
}

func (ls *LogininforService) SelectLogininforList(c *gin.Context, logininfor *monitorModels.LogininforDQL) (list []*monitorModels.Logininfor, total *int64) {
	return ls.ld.SelectLogininforList(c, ls.data, logininfor)

}
func (ls *LogininforService) ExportLogininfor(c *gin.Context, logininfor *monitorModels.LogininforDQL) (data []byte) {
	//list, _ := ls.ld.SelectLogininforList(ls.data.GetSlaveDb(), logininfor)
	//return exceLize.SetRows(systemModels.SysLogininforToRows(list))
	return nil
}

func (ls *LogininforService) InserLogininfor(c *gin.Context, loginUser *monitorModels.Logininfor) {
	loginUser.InfoId = snowflake.GenID()
	ls.ld.InserLogininfor(c, ls.data, loginUser)
}

func (ls *LogininforService) DeleteLogininforByIds(c *gin.Context, infoIds []int64) {
	ls.ld.DeleteLogininforByIds(c, ls.data, infoIds)

}

func (ls *LogininforService) CleanLogininfor(c *gin.Context) {
	ls.ld.CleanLogininfor(c, ls.data)

}
