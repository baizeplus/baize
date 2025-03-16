package systemModels

type SelectPermission struct {
	PermissionId   int64  `json:"permissionId,string" db:"permission_id"`
	PermissionName string `json:"permissionName" db:"permission_name"`
	ParentId       int64  `json:"parentId" db:"parent_id"`
}
