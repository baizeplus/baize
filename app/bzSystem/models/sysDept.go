package models

import "baize/app/baize"

type SysDeptDQL struct {
	ParentId int64  `form:"parentId,string" db:"parent_id"` //上级id
	DeptName string `form:"deptName" db:"dept_name"`        //部门名称
	Status   string `form:"status" db:"status"`             //状态
}

type SysDeptVo struct {
	DeptId    int64  `json:"deptId,string" db:"dept_id"`     //id
	ParentId  int64  `json:"parentId,string" db:"parent_id"` //上级id
	Ancestors string `json:"ancestors" db:"ancestors"`       //祖级列表
	DeptName  string `json:"deptName" db:"dept_name"`        //部门名称
	OrderNum  string `json:"orderNum" db:"order_num"`        //排序
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
