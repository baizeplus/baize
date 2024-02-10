package models

type LoginBody struct {
	Username string `json:"username" binding:"required"` //用户名
	Password string `json:"password" binding:"required"` //密码
	Code     string `json:"code" binding:"required"`     //验证码
	Uuid     string `json:"uuid" binding:"required"`     //uuid
}

//
//type LoginUser struct {
//	Token         string
//	LoginTime     int64
//	ExpireTime    int64
//	IpAddr        string
//	LoginLocation string
//	Browser       string
//	Os            string
//	User          *User
//	RolePerms     []string
//	Permissions   []string
//}

type User struct {
	UserId   int64  `json:"userId,string" db:"user_id"`
	DeptId   int64  `json:"-" db:"dept_id"`
	UserName string `json:"userName" db:"user_name"`
	//NickName    string        `json:"nickName" db:"nick_name"`
	//Email       string        `json:"email" db:"email"`
	//Phonenumber string        `json:"phonenumber" db:"phonenumber"`
	//Sex         string        `json:"sex" db:"sex"`
	Avatar   string `json:"avatar" db:"avatar" `
	Password string `json:"-" db:"password"`
	Status   string `json:"-" db:"status"`
	DelFlag  string `json:"-" db:"del_flag"`
	//ParentId    int64         `json:"parentId" db:"parent_id"`
	//DeptName    string        `json:"deptName" db:"dept_name"`
	//CreateTime  *baize.Time   `json:"createTime" db:"create_time" db:"create_time"  swaggertype:"integer"`
	//Roles       []*baize.Role `json:"roles"`
}

type GetInfo struct {
	User        *User    `json:"user"`
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}
