package genTableController

import (
	"baize/app/business/tool/toolModels"
	"baize/app/business/tool/toolService"
	"baize/app/business/tool/toolService/toolServiceImpl"
	"baize/app/utils/baizeContext"
	"github.com/gin-gonic/gin"
	"strings"
)

type GenTable struct {
	gt toolService.IGenTableService
}

func NewGenTable(ls *toolServiceImpl.GenTabletService) *GenTable {
	return &GenTable{gt: ls}
}

func (gc *GenTable) GenTableList(c *gin.Context) {
	getTable := new(toolModels.GenTableDQL)
	c.ShouldBind(getTable)
	list, count := gc.gt.SelectGenTableList(c, getTable)
	baizeContext.SuccessListData(c, list, count)
}

func (gc *GenTable) GenTableGetInfo(c *gin.Context) {
	tableId := baizeContext.ParamInt64(c, "tableId")
	genTable := gc.gt.SelectGenTableById(c, tableId)
	tables := gc.gt.SelectGenTableAll(c)
	list := gc.gt.SelectGenTableColumnListByTableId(c, tableId)
	data := make(map[string]interface{})
	data["info"] = genTable
	data["rows"] = list
	data["tables"] = tables
	baizeContext.SuccessData(c, data)
}

func (gc *GenTable) DataList(c *gin.Context) {
	getTable := new(toolModels.GenTableDQL)
	c.ShouldBind(getTable)
	list, count := gc.gt.SelectDbTableList(c, getTable)
	baizeContext.SuccessListData(c, list, count)

}
func (gc *GenTable) ColumnList(c *gin.Context) {
	tableId := baizeContext.ParamInt64(c, "tableId")
	list := gc.gt.SelectGenTableColumnListByTableId(c, tableId)
	total := int64(len(list))
	baizeContext.SuccessListData(c, list, &total)
}
func (gc *GenTable) ImportTable(c *gin.Context) {
	gc.gt.ImportTableSave(c, strings.Split(c.Query("tables"), ","), "")
	baizeContext.SuccessMsg(c, "导入成功")
}
func (gc *GenTable) EditSave(c *gin.Context) {
	genTable := new(toolModels.GenTableDML)
	c.ShouldBindJSON(genTable)
	//genTable.SetUpdateBy(bzc.GetCurrentUserName())
	gc.gt.UpdateGenTable(c, genTable)
	baizeContext.Success(c)
}
func (gc *GenTable) GenTableRemove(c *gin.Context) {
	gc.gt.DeleteGenTableByIds(c, baizeContext.ParamInt64Array(c, "tableIds"))
	baizeContext.Success(c)
}
func (gc *GenTable) Preview(c *gin.Context) {
	gc.gt.PreviewCode(c, baizeContext.ParamInt64(c, "tableId"))
	baizeContext.Success(c)
}
func (gc *GenTable) Download(c *gin.Context) {
	panic("等待开发")
}

func (gc *GenTable) GenCode(c *gin.Context) {
	panic("等待开发")
}

func (gc *GenTable) SynchDb(c *gin.Context) {
	panic("等待开发")

}

func (gc *GenTable) BatchGenCode(c *gin.Context) {
	panic("等待开发")

}
