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
