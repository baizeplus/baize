package systemModels

import "baize/app/baize"

type SysRole struct {
	RoleId   int64  `db:"role_id"`
	RoleName string `db:"role_name"`
	RoleKey  string `db:"role_key"`
}

type SysRoleDQL struct {
	RoleName  string `form:"roleName" db:"role_name"`
	Status    string `form:"status" db:"status"`
	RoleKey   string `form:"roleKey" db:"role_key"`
	BeginTime string `form:"beginTime" db:"begin_time"`
	EndTime   string `form:"endTime" db:"end_time"`
	CreateBy  int64  `db:"create_by" swaggerignore:"true"` //创建人
	baize.BaseEntityDQL
}

type SysRoleVo struct {
	RoleId   int64  `json:"roleId,string" db:"role_id"`
	RoleName string `json:"roleName" db:"role_name" bze:"1,角色名称"`
	RoleKey  string `json:"roleKey" db:"role_key" bze:"2,权限字符"`
	RoleSort int    `json:"roleSort" db:"role_sort"`
	Status   string `json:"status"  db:"status"`
	DelFlag  string `json:"delFlag" db:"del_flag"`
	Remake   string `json:"remark" db:"remark"`
	baize.BaseEntity
}

type SysRoleDML struct {
	RoleId   int64    `json:"RoleId,string" db:"role_id"`
	RoleName string   `json:"roleName" db:"role_name"`
	RoleKey  string   `json:"roleKey" db:"role_key"`
	RoleSort int      `json:"roleSort" db:"role_sort"`
	Status   string   `json:"status" db:"status"`
	Remake   string   `json:"remark" db:"remark"`
	MenuIds  []string `json:"menuIds"`
	baize.BaseEntity
}

type SysRoleMenu struct {
	RoleId int64 `db:"role_id"`
	MenuId int64 `db:"menu_id"`
}

type SysRoleAndUserDQL struct {
	RoleId      string `form:"roleId" db:"role_id" binding:"required"`
	UserName    string `form:"userName" db:"user_name"`
	Phonenumber string `form:"phonenumber" db:"phonenumber"`

	baize.BaseEntityDQL
}
