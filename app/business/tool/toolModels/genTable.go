package toolModels

import (
	"baize/app/baize"
	genUtils "baize/app/business/tool/utils"
	"baize/app/utils/stringUtils"
	"strings"
)

type GenTableDQL struct {
	TableName    string `form:"tableName" db:"table_name"`
	TableComment string `form:"tableComment" db:"table_comment"`
	BeginTime    string `form:"beginTime" db:"begin_time"`
	EndTime      string `form:"endTime" db:"end_time"`
	baize.BaseEntityDQL
}

type GenTableDML struct {
	TableId        int64                `json:"tableId,string" db:"table_id"`
	TableName      string               `json:"tableName" db:"table_name"`
	TableComment   string               `json:"tableComment" db:"table_comment"`
	SubTableName   string               `json:"subTableName" db:"sub_table_name"`
	SubTableFkName string               `json:"subTableFkName" db:"sub_table_fk_name"`
	StructName     string               `json:"structName" db:"struct_name"`
	TplCategory    string               `json:"tplCategory" db:"tpl_category"`
	PackageName    string               `json:"packageName" db:"package_name"`
	ModuleName     string               `json:"moduleName" db:"module_name"`
	BusinessName   string               `json:"businessName" db:"business_name"`
	FunctionName   string               `json:"functionName" db:"function_name"`
	FunctionAuthor string               `json:"functionAuthor" db:"function_author"`
	Options        string               `json:"options" db:"options"`
	Remark         string               `json:"remark" db:"remark"`
	Columns        []*GenTableColumnDML `json:"columns"`
	baize.BaseEntity
}

func GetGenTableDML(table *DBTableVo, tableId int64, userId int64) *GenTableDML {
	gen := new(GenTableDML)
	gen.TableId = tableId
	gen.TableName = table.TableName
	gen.TableComment = table.TableComment
	gen.StructName = stringUtils.ConvertToBigCamelCase(genUtils.ConvertClassName(table.TableName))
	gen.PackageName = "baize/app/business/tool/"
	gen.ModuleName = "module"
	gen.BusinessName = genUtils.GetBusinessName(table.TableName)
	gen.FunctionName = strings.ReplaceAll(table.TableComment, "表", "")
	gen.FunctionAuthor = "baize"
	gen.TplCategory = "crud"
	gen.SetCreateBy(userId)
	return gen
}

type GenTableVo struct {
	TableId        int64   `json:"tableId,string" db:"table_id"`
	TableName      string  `json:"tableName" db:"table_name"`
	TableComment   string  `json:"tableComment" db:"table_comment"`
	SubTableName   *string `json:"subTableName" db:"sub_table_name"`
	SubTableFkName *string `json:"subTableFkName" db:"sub_table_fk_name"`
	StructName     string  `json:"structName" db:"struct_name"`
	TplCategory    string  `json:"tplCategory" db:"tpl_category"`
	PackageName    string  `json:"packageName" db:"package_name"`
	ModuleName     string  `json:"moduleName" db:"module_name"`
	BusinessName   string  `json:"businessName" db:"business_name"`
	FunctionName   string  `json:"functionName" db:"function_name"`
	FunctionAuthor string  `json:"functionAuthor" db:"function_author"`
	Options        *string `json:"options" db:"options"`
	Remark         string  `json:"remark" db:"remark"`
	baize.BaseEntity
}

type DBTableVo struct {
	TableName    string      `json:"tableName" db:"TABLE_NAME"`
	TableComment string      `json:"tableComment" db:"TABLE_COMMENT"`
	CreateTime   *baize.Time `json:"createTime" db:"CREATE_TIME"`
	UpdateTime   *baize.Time `json:"updateTime" db:"UPDATE_TIME"`
}
