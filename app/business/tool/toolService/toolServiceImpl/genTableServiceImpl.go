package toolServiceImpl

import (
	"archive/zip"
	"baize/app/business/tool/toolDao"
	"baize/app/business/tool/toolDao/toolDaoImpl"
	"baize/app/business/tool/toolModels"
	genUtils "baize/app/business/tool/utils"
	"baize/app/utils/baizeContext"
	"baize/app/utils/snowflake"
	"baize/app/utils/zipUtils"
	"bytes"
	"fmt"
	"github.com/baizeplus/sqly"
	"github.com/gin-gonic/gin"
	"go/format"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
)

type GenTabletService struct {
	data               *sqly.DB
	genTabletDao       toolDao.IGenTable
	genTabletColumnDao toolDao.IGenTableColumn
}

func NewGenTabletService(data *sqly.DB, gtc *toolDaoImpl.GenTableColumnDao, gt *toolDaoImpl.GenTableDao,
) *GenTabletService {
	return &GenTabletService{
		data:               data,
		genTabletDao:       gt,
		genTabletColumnDao: gtc,
	}
}

func (genTabletService *GenTabletService) SelectGenTableList(c *gin.Context, getTable *toolModels.GenTableDQL) (list []*toolModels.GenTableVo, total int64) {
	return genTabletService.genTabletDao.SelectGenTableList(c, genTabletService.data, getTable)
}
func (genTabletService *GenTabletService) SelectDbTableList(c *gin.Context, getTable *toolModels.GenTableDQL) (list []*toolModels.DBTableVo, total int64) {
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
		genTableList = append(genTableList, toolModels.GetGenTableDML(genTable, tableId, baizeContext.GetUserId(c)))
		list := genTabletService.genTabletColumnDao.SelectDbTableColumnsByName(c, genTabletService.data, genTable.TableName)
		for _, column := range list {
			genTableColumnList = append(genTableColumnList, toolModels.GetGenTableColumnDML(column, tableId, baizeContext.GetUserId(c)))
		}
	}
	genTabletService.genTabletDao.BatchInsertGenTable(c, genTabletService.data, genTableList)
	genTabletService.genTabletColumnDao.BatchInsertGenTableColumn(c, genTabletService.data, genTableColumnList)

}
func (genTabletService *GenTabletService) UpdateGenTable(c *gin.Context, genTable *toolModels.GenTableDML) {
	genTabletService.genTabletDao.UpdateGenTable(c, genTabletService.data, genTable)
	for _, cenTableColumn := range genTable.Columns {
		genTabletService.genTabletColumnDao.UpdateGenTableColumn(c, genTabletService.data, cenTableColumn)
	}
	return
}

func (genTabletService *GenTabletService) DeleteGenTableByIds(c *gin.Context, ids []int64) {
	genTabletService.genTabletDao.DeleteGenTableByIds(c, genTabletService.data, ids)
	genTabletService.genTabletColumnDao.DeleteGenTableColumnByIds(c, genTabletService.data, ids)
	return
}
func (genTabletService *GenTabletService) PreviewCode(c *gin.Context, tableId int64) (m map[string]string) {
	data := make(map[string]any)
	data["Table"] = genTabletService.genTabletDao.SelectGenTableById(c, genTabletService.data, tableId)
	data["Columns"] = genTabletService.genTabletColumnDao.SelectGenTableColumnListByTableId(c, genTabletService.data, tableId)
	m = make(map[string]string)
	root := "./template/go/"
	var files []string
	err := filepath.Walk(root, visit(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		formattedCode, err := format.Source(genTabletService.loadTemplate("./"+file, data))
		if err != nil {
			panic(err)
		}
		m[filepath.Base(file)] = string(formattedCode)
	}
	root = "./template/vue"
	files = files[:0]
	err = filepath.Walk(root, visit(&files))
	if err != nil {
		panic(err)
	}
	files = append(files, "template/sql/sql.sql.tmpl")
	for _, file := range files {
		loadTemplate := genTabletService.loadTemplate("./"+file, data)
		m[filepath.Base(file)] = string(loadTemplate)
	}

	file := "/template/sql/sql.sql.tmpl"
	loadTemplate := genTabletService.loadTemplate("./"+file, data)
	m[filepath.Base(file)] = string(loadTemplate)

	return m
}
func (genTabletService *GenTabletService) GenCode(c *gin.Context, tableId int64) []byte {
	// 创建一个内存缓冲区
	buffer := new(bytes.Buffer)
	// 创建一个新的 zip Writer
	zipWriter := zip.NewWriter(buffer)
	data := make(map[string]any)
	data["Table"] = genTabletService.genTabletDao.SelectGenTableById(c, genTabletService.data, tableId)
	data["Columns"] = genTabletService.genTabletColumnDao.SelectGenTableColumnListByTableId(c, genTabletService.data, tableId)
	root := "./template/go/"
	var files []string
	err := filepath.Walk(root, visit(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		formattedCode, err := format.Source(genTabletService.loadTemplate("./"+file, data))
		if err != nil {
			fmt.Println(err)
		}
		if err := zipUtils.AddFileToZip(zipWriter, strings.TrimSuffix(strings.TrimPrefix(file, "template\\"), ".tmpl"), string(formattedCode)); err != nil {
			panic(err)
		}
	}

	root = "./template/vue"
	files = files[:0]
	err = filepath.Walk(root, visit(&files))
	if err != nil {
		panic(err)
	}
	files = append(files, "template/sql/sql.sql.tmpl")
	for _, file := range files {
		loadTemplate := genTabletService.loadTemplate("./"+file, data)
		if err := zipUtils.AddFileToZip(zipWriter, strings.TrimSuffix(strings.TrimPrefix(file, "template\\"), ".tmpl"), string(loadTemplate)); err != nil {
			panic(err)
		}
	}

	// 关闭压缩包
	if err := zipWriter.Close(); err != nil {
		panic(err)
	}
	// 将缓冲区的内容写入到返回的字节切片中
	return buffer.Bytes()
}
func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if !info.IsDir() {
			*files = append(*files, path)
		}
		return nil
	}
}

func (genTabletService *GenTabletService) SelectGenTableColumnListByTableId(c *gin.Context, tableId int64) (list []*toolModels.GenTableColumnVo) {
	return genTabletService.genTabletColumnDao.SelectGenTableColumnListByTableId(c, genTabletService.data, tableId)
}

func (genTabletService *GenTabletService) loadTemplate(templateName string, data map[string]any) []byte {
	genTabletService.setTemplateData(data)
	b, err := os.ReadFile(templateName)
	if err != nil {
		panic(err)
	}
	templateStr := string(b)
	tmpl := template.New(templateName)
	tmpl.Funcs(template.FuncMap{"Contains": genUtils.Contains, "CaseCamelLower": genUtils.CaseCamelLower, "HasSuffix": strings.HasSuffix})
	// 解析模板字符串
	tmpl, err = tmpl.Parse(templateStr)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBufferString("")
	err = tmpl.Execute(buffer, data) //将string与模板合成，变量name的内容会替换掉{{.}}
	if err != nil {
		print(err)
	}

	return buffer.Bytes()
}

func (genTabletService *GenTabletService) setTemplateData(data map[string]any) {
	data["GenerateTime"] = time.Now()
	column := data["Columns"].([]*toolModels.GenTableColumnVo)
	data["ColumnsLastIndex"] = len(column) - 1
	for _, vo := range column {
		if vo.IsPk == "1" {
			data["IdField"] = vo.HtmlField
			data["IdGoField"] = vo.GoField
			data["IdType"] = vo.GoType
			data["IdColumnName"] = vo.ColumnName
			break
		}
	}
	for _, vo := range column {
		if vo.IsRequired == "1" && vo.GoType == "Time" {
			data["ContainsTimeType"] = true
			break
		}
	}
}
