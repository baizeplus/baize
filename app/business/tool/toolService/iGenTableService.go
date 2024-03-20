package toolService

import (
	"baize/app/business/tool/toolModels"
	"github.com/gin-gonic/gin"
)

type IGenTableService interface {
	SelectGenTableList(c *gin.Context, getTable *toolModels.GenTableDQL) (list []*toolModels.GenTableVo, total *int64)
	SelectDbTableList(c *gin.Context, getTable *toolModels.GenTableDQL) (list []*toolModels.DBTableVo, total *int64)
	SelectGenTableAll(c *gin.Context) (list []*toolModels.GenTableVo)
	SelectGenTableById(c *gin.Context, id int64) (genTable *toolModels.GenTableVo)
	ImportTableSave(c *gin.Context, table []string, userName string)
	UpdateGenTable(c *gin.Context, genTable *toolModels.GenTableDML) (err error)
	DeleteGenTableByIds(c *gin.Context, ids []int64) (err error)
	PreviewCode(c *gin.Context, tableId int64) (genTable *toolModels.GenTableVo, err error)
	SelectGenTableColumnListByTableId(c *gin.Context, tableId int64) (list []*toolModels.GenTableColumnVo)
}
