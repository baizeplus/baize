package systemModels

import "baize/app/baize"

type SysRole struct {
	RoleId   string `db:"role_id"`
	RoleName string `db:"role_name"`
}

type SysRoleDQL struct {
	RoleName  string `form:"roleName" db:"role_name"`
	Status    string `form:"status" db:"status"`
	BeginTime string `form:"beginTime" db:"begin_time"`
	EndTime   string `form:"endTime" db:"end_time"`
	CreateBy  string `db:"create_by" swaggerignore:"true"` //创建人
	baize.BaseEntityDQL
}

type SysRoleVo struct {
	RoleId        string   `json:"roleId" db:"role_id"`
	RoleName      string   `json:"roleName" db:"role_name" bze:"1,角色名称"`
	RoleSort      int      `json:"roleSort" db:"role_sort"`
	Status        string   `json:"status"  db:"status"`
	Remake        string   `json:"remark" db:"remark"`
	PermissionIds []string `json:"permissionIds"`
	baize.BaseEntity
}

type SysRoleDML struct {
	RoleId        string   `json:"RoleId" db:"role_id"`
	RoleName      string   `json:"roleName" db:"role_name"`
	RoleSort      int      `json:"roleSort" db:"role_sort"`
	Status        string   `json:"status" db:"status"`
	DelFlag       string   `json:"delFlag" db:"del_flag"`
	Remake        string   `json:"remark" db:"remark"`
	PermissionIds []string `json:"permissionIds"`
	baize.BaseEntity
}

type SysRolePermission struct {
	RoleId       string `db:"role_id"`
	PermissionId string `db:"permission_id"`
}

type SysRoleAndUserDQL struct {
	RoleId      string `form:"roleId" db:"role_id" binding:"required"`
	UserName    string `form:"userName" db:"user_name"`
	Phonenumber string `form:"phonenumber" db:"phonenumber"`
	baize.BaseEntityDQL
}

type SysRoleIdAndName struct {
	RoleId   string `json:"roleId" db:"role_id"`
	RoleName string `json:"roleName" db:"role_name" `
}
