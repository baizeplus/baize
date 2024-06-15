package genUtils

import (
	"regexp"
	"strings"
)

// ColumnTypeStr 数据库字符串类型
var ColumnTypeStr = []string{"char", "varchar", "narchar", "varchar2", "tinytext", "text", "mediumtext", "longtext"}

// ColumnTypeTime 数据库时间类型
var ColumnTypeTime = []string{"datetime", "time", "date", "timestamp"}

// ColumnTypeNumber 数据库数字类型
var ColumnTypeNumber = []string{"tinyint", "smallint", "mediumint", "int", "number", "integer", "bigint", "float", "double", "decimal"}

// ColumnNameBaseEntity 页面基础字段
var ColumnNameBaseEntity = []string{"create_by", "create_time", "update_by", "update_time"}

// IsExistInArray 正则判断string 是否存在在数组中
func IsExistInArray(value string, array []string) bool {
	re, err := regexp.Compile(strings.Join(array, "|"))
	if err != nil {
		return false
	}
	return re.MatchString(value)
}

// IsStringObject 判断是否是数据库字符串类型
func IsStringObject(value string) bool {
	return IsExistInArray(value, ColumnTypeStr)
}

// IsTimeObject 判断是否是数据库时间类型
func IsTimeObject(value string) bool {
	return IsExistInArray(value, ColumnTypeTime)
}

// IsNumberObject 判断是否是数据库数字类型
func IsNumberObject(value string) bool {
	return IsExistInArray(value, ColumnTypeNumber)
}

// IsNotEntity 页面不需要编辑字段
func IsNotEntity(value string) bool {
	return !IsExistInArray(value, ColumnNameBaseEntity)
}
