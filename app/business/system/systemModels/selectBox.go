package systemModels

type SelectPermission struct {
	PermissionId   int64  `json:"permissionId,string" db:"permission_id"`
	PermissionName string `json:"permissionName" db:"permission_name"`
	ParentId       int64  `json:"parentId,string" db:"parent_id"`
}

type SelectDept struct {
	DeptId   int64  `json:"deptId,string" db:"dept_id"`
	DeptName string `json:"deptName" db:"dept_name"`
	ParentId int64  `json:"parentId,string" db:"parent_id"`
}
