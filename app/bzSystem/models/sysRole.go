package models

import "baize/app/baize"

type SysRole struct {
	RoleId    int64  `db:"role_id"`
	RoleName  string `db:"role_name"`
	RoleKey   string `db:"role_key"`
	DataScope string `db:"data_scope"`
}

type SysRoleDQL struct {
	RoleName  string `form:"roleName" db:"role_name"`
	Status    string `form:"status" db:"status"`
	RoleKey   string `form:"roleKey" db:"role_key"`
	BeginTime string `form:"beginTime" db:"begin_time"`
	EndTime   string `form:"endTime" db:"end_time"`
	baize.BaseEntityDQL
}

type SysRoleVo struct {
	RoleId            int64  `json:"roleId,string" db:"role_id"`
	RoleName          string `json:"roleName" db:"role_name"`
	RoleKey           string `json:"roleKey" db:"role_key"`
	RoleSort          int    `json:"roleSort" db:"role_sort"`
	DataScope         string `json:"dataScope" db:"data_scope"`
	MenuCheckStrictly bool   `json:"permissionCheckStrictly" db:"menu_check_strictly"`
	DeptCheckStrictly bool   `json:"deptCheckStrictly" db:"dept_check_strictly"`
	Status            string `json:"status"  db:"status"`
	DelFlag           string `json:"delFlag" db:"del_flag"`
	Remake            string `json:"remark" db:"remark"`
	baize.BaseEntity
}

type SysRoleDML struct {
	RoleId                  int64    `json:"RoleId,string" db:"role_id"`
	RoleName                string   `json:"roleName" db:"role_name"`
	RoleKey                 string   `json:"roleKey" db:"role_key"`
	RoleSort                int      `json:"roleSort" db:"role_sort"`
	DataScope               string   `json:"dataScope" db:"data_scope"`
	PermissionCheckStrictly bool     `json:"PermissionCheckStrictly" db:"Permission_check_strictly"`
	DeptCheckStrictly       bool     `json:"deptCheckStrictly" db:"dept_check_strictly" `
	Status                  string   `json:"status" db:"status"`
	Remake                  string   `json:"remark" db:"remark "`
	PermissionIds           []string `json:"PermissionIds"`
	DeptIds                 []string `json:"deptIds"`
	baize.BaseEntity
}

type SysRoleMenu struct {
	RoleId int64 `db:"role_id"`
	MenuId int64 `db:"menu_id"`
}

type SysRoleDept struct {
	RoleId int64 `db:"role_id"`
	DeptId int64 `db:"dept_id"`
}

type SysRoleAndUserDQL struct {
	RoleId      string `form:"roleId" db:"role_id" binding:"required"`
	UserName    string `form:"userName" db:"user_name"`
	Phonenumber string `form:"phonenumber" db:"phonenumber"`

	baize.BaseEntityDQL
}
