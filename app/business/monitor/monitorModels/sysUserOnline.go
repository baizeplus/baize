package monitorModels

type SysUserOnline struct {
	TokenId   string `json:"tokenId"`
	UserName  string `json:"userName"`
	Ipaddr    string `json:"ipaddr"`
	Browser   string `json:"browser"`
	Os        string `json:"os"`
	LoginTime int64  `json:"loginTime"`
	DeptName  string `json:"deptName"`
}
type SysUserOnlineDQL struct {
	UserName string `form:"userName" db:"user_name"`
	Ipaddr   string `form:"ipaddr" db:"ipaddr"`
}
