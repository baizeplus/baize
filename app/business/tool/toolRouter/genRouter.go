package toolRouter

import (
	"baize/app/business/tool/toolController"
	"baize/app/middlewares"
	"github.com/gin-gonic/gin"
)

func InitGenTableRouter(router *gin.RouterGroup, gt *toolController.GenTable) {
	genTable := router.Group("/tool/gen")
	genTable.GET("/list", middlewares.HasPermission("tool:gen:list"), gt.GenTableList)
	genTable.GET(":tableId", middlewares.HasPermission("tool:gen:query"), gt.GenTableGetInfo)
	genTable.GET("/db/list", middlewares.HasPermission("tool:gen:list"), gt.DataList)
	genTable.GET("/column/:talbleId", middlewares.HasPermission("tool:gen:list"), gt.ColumnList)
	genTable.POST("/importTable", middlewares.HasPermission("tool:gen:list"), gt.ImportTable)
	genTable.PUT("", middlewares.HasPermission("tool:gen:edit"), gt.EditSave)
	genTable.DELETE("/:tableIds", middlewares.HasPermission("tool:gen:remove"), gt.GenTableRemove)
	genTable.GET("/preview/:tableId", middlewares.HasPermission("tool:gen:code"), gt.Preview)
	genTable.GET("/genCode/:tableId", middlewares.HasPermission("tool:gen:code"), gt.GenCode)
}
