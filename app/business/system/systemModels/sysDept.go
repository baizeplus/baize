package systemModels

import (
	"baize/app/baize"
	"sort"
)

type SysDeptDQL struct {
	ParentId int64  `form:"parentId,string" db:"parent_id"` //上级id
	DeptName string `form:"deptName" db:"dept_name"`        //部门名称
	Status   string `form:"status" db:"status"`             //状态
	baize.BaseEntityDQL
}

type SysDeptVo struct {
	DeptId    int64  `json:"deptId,string" db:"dept_id"`     //id
	ParentId  int64  `json:"parentId,string" db:"parent_id"` //上级id
	Ancestors string `json:"ancestors" db:"ancestors"`       //祖级列表
	DeptName  string `json:"deptName" db:"dept_name"`        //部门名称
	OrderNum  int64  `json:"orderNum" db:"order_num"`        //排序
	Leader    string `json:"leader" db:"leader"`             //负责人
	Phone     string `json:"phone" db:"phone"`               //电话
	Email     string `json:"email" db:"email"`               //邮箱
	Status    string `json:"status" db:"status"`             //状态
	DelFlag   string `json:"delFag" db:"del_flag"`           //删除标志
	baize.BaseEntity
}

type RoleDeptTree struct {
	CheckedKeys []string     `json:"checkedKeys"` //keys
	Depts       []*SysDeptVo `json:"depts"`       //部门
}

func GetParentNameAll(items []*SysDeptVo) []string {
	ss := make([]string, len(items))
	for _, item := range items {
		ss = append(ss, GetParentName(items, item.ParentId, item.DeptName))
	}
	sort.Strings(ss)
	return ss
}

func GetParentName(items []*SysDeptVo, parentId int64, name string) string {

	for _, item := range items {
		if item.DeptId == parentId {
			pName := item.DeptName + "/" + name
			return GetParentName(items, item.ParentId, pName)
		}
	}
	return name
}

func GetParentNameAndIds(items []*SysDeptVo) map[string]int64 {
	ss := make(map[string]int64)
	for _, item := range items {
		ss[GetParentName(items, item.ParentId, item.DeptName)] = item.DeptId
	}
	return ss
}
