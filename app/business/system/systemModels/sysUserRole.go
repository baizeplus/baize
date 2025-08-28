package systemModels

type SysUserRole struct {
	UserId string `json:"userId" db:"user_id"`
	RoleId string `json:"roleId" db:"role_id"`
}

func NewSysUserRole(userId string, roleId string) *SysUserRole {
	return &SysUserRole{UserId: userId, RoleId: roleId}
}
