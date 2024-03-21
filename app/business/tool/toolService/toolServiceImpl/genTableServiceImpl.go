package toolServiceImpl

import (
	"baize/app/business/tool/toolDao"
	"baize/app/business/tool/toolModels"
	"baize/app/utils/snowflake"
	"bytes"
	"fmt"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
	"html/template"
	"os"
	"time"
)

type GenTabletService struct {
	data               *sqly.DB
	genTabletDao       toolDao.IGenTable
	genTabletColumnDao toolDao.IGenTableColumn
}

func GetGenTabletService() *GenTabletService {
	return &GenTabletService{}
}

func (genTabletService *GenTabletService) SelectGenTableList(c *gin.Context, getTable *toolModels.GenTableDQL) (list []*toolModels.GenTableVo, total *int64) {
	return genTabletService.genTabletDao.SelectGenTableList(c, genTabletService.data, getTable)
}
func (genTabletService *GenTabletService) SelectDbTableList(c *gin.Context, getTable *toolModels.GenTableDQL) (list []*toolModels.DBTableVo, total *int64) {
	return genTabletService.genTabletDao.SelectDbTableList(c, genTabletService.data, getTable)
}
func (genTabletService *GenTabletService) SelectGenTableAll(c *gin.Context) (list []*toolModels.GenTableVo) {
	return genTabletService.genTabletDao.SelectGenTableAll(c, genTabletService.data)
}
func (genTabletService *GenTabletService) SelectGenTableById(c *gin.Context, id int64) (genTable *toolModels.GenTableVo) {
	return genTabletService.genTabletDao.SelectGenTableById(c, genTabletService.data, id)
}
func (genTabletService *GenTabletService) ImportTableSave(c *gin.Context, table []string, userName string) {
	tableList := genTabletService.genTabletDao.SelectDbTableListByNames(c, genTabletService.data, table)
	genTableList := make([]*toolModels.GenTableDML, 0, len(tableList))
	genTableColumnList := make([]*toolModels.GenTableColumnDML, 0, len(tableList)*2)
	for _, genTable := range tableList {
		tableId := snowflake.GenID()
		genTableList = append(genTableList, toolModels.GetGenTableDML(genTable, tableId, userName))
		list := genTabletService.genTabletColumnDao.SelectDbTableColumnsByName(c, genTabletService.data, genTable.TableName)
		for _, column := range list {
			genTableColumnList = append(genTableColumnList, toolModels.GetGenTableColumnDML(column, tableId, userName))
		}
	}
	genTabletService.genTabletDao.BatchInsertGenTable(c, genTabletService.data, genTableList)
	genTabletService.genTabletColumnDao.BatchInsertGenTableColumn(c, genTabletService.data, genTableColumnList)

}
func (genTabletService *GenTabletService) UpdateGenTable(c *gin.Context, genTable *toolModels.GenTableDML) (err error) {
	genTabletService.genTabletDao.UpdateGenTable(c, genTabletService.data, genTable)
	for _, cenTableColumn := range genTable.Columns {
		genTabletService.genTabletColumnDao.UpdateGenTableColumn(c, genTabletService.data, cenTableColumn)
	}
	return
}

func (genTabletService *GenTabletService) DeleteGenTableByIds(c *gin.Context, ids []int64) (err error) {
	genTabletService.genTabletDao.DeleteGenTableByIds(c, genTabletService.data, ids)
	genTabletService.genTabletColumnDao.DeleteGenTableColumnByIds(c, genTabletService.data, ids)
	return nil
}
func (genTabletService *GenTabletService) PreviewCode(c *gin.Context, tableId int64) (genTable *toolModels.GenTableVo, err error) {
	genTable = genTabletService.genTabletDao.SelectGenTableById(c, genTabletService.data, tableId)
	genTable.Columns = genTabletService.genTabletColumnDao.SelectGenTableColumnListByTableId(c, genTabletService.data, tableId)
	genTable.GenerateTime = time.Now()
	s := genTabletService.loadTemplate("./template/vm/go/model/model.tmpl", genTable)
	fmt.Println(s)
	return genTable, nil
}
func (genTabletService *GenTabletService) SelectGenTableColumnListByTableId(c *gin.Context, tableId int64) (list []*toolModels.GenTableColumnVo) {
	return genTabletService.genTabletColumnDao.SelectGenTableColumnListByTableId(c, genTabletService.data, tableId)
}

func (genTabletService *GenTabletService) loadTemplate(templateName string, data interface{}) string {
	b, err := os.ReadFile(templateName)
	if err != nil {
		panic(err)
	}
	templateStr := string(b)
	tmpl, err := template.New(templateName).Parse(templateStr) //建立一个模板，内容是"hello, {{.}}"
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBufferString("")
	err = tmpl.Execute(buffer, gin.H{"table": data}) //将string与模板合成，变量name的内容会替换掉{{.}}
	if err != nil {
		print(err)
	}
	return buffer.String()
}
