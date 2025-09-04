package systemModels

type SelectPermission struct {
	PermissionId   string `json:"permissionId" db:"permission_id"`
	PermissionName string `json:"permissionName" db:"permission_name"`
	ParentId       string `json:"parentId" db:"parent_id"`
}

type SelectDept struct {
	DeptId   string `json:"deptId" db:"dept_id"`
	DeptName string `json:"deptName" db:"dept_name"`
	ParentId string `json:"parentId" db:"parent_id"`
}
