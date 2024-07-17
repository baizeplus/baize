package systemModels

import (
	"baize/app/baize"
	"baize/app/constant/dataScopeAspect"
	"baize/app/utils/snowflake"
	"strconv"
)

type SysUserDQL struct {
	UserName    string `form:"userName" db:"user_name"`      //用户名
	Status      string `form:"status" db:"status"`           //状态
	Phonenumber string `form:"phonenumber" db:"phonenumber"` //电话
	BeginTime   string `form:"beginTime" db:"begin_time"`    //注册开始时间
	EndTime     string `form:"endTime" db:"end_time"`        //注册结束时间
	DeptId      int64  `form:"deptId" db:"dept_id"`          //部门ID
	baize.BaseEntityDQL
}

type SysUserDML struct {
	UserId      int64    `json:"userId,string" db:"user_id"swaggerignore:"true"` //用户ID
	DeptId      int64    `json:"deptId,string" db:"dept_id" binding:"required"`  //部门ID
	UserName    string   `json:"userName" db:"user_name" binding:"required"`     //用户名
	NickName    string   `json:"nickName" db:"nick_name" binding:"required"`     //用户昵称
	Email       string   `json:"email" db:"email"`                               //邮箱
	Avatar      string   `json:"avatar" db:"avatar"`                             //头像
	Phonenumber string   `json:"phonenumber" db:"phonenumber"`                   //手机号
	Sex         string   `json:"sex" db:"sex"  binding:"required"`               //性别
	Password    string   `json:"password" db:"password" binding:"required"`      //密码
	DataScope   string   `json:"dataScope" db:"data_scope"`                      //权限范围
	Status      string   `json:"status" db:"status"`                             //状态
	Remark      string   `json:"remark" db:"remark"`                             //备注
	PostIds     []string `json:"postIds"`                                        //岗位IDS
	RoleIds     []string `json:"roleIds"`                                        //角色IDS
	baize.BaseEntity
}

type SysUserVo struct {
	UserId      int64   `json:"userId,string" db:"user_id"`
	UserName    string  `json:"userName" db:"user_name" bze:"1,用户名"`
	NickName    string  `json:"nickName" db:"nick_name" bze:"2,用户昵称"`
	Sex         string  `json:"sex" db:"sex" bze:"3,性别"`
	Status      string  `json:"status" db:"status"`
	DelFlag     string  `json:"delFlag" db:"del_flag"`
	DeptId      int64   `json:"deptId,string" db:"dept_id"`
	DeptName    *string `json:"deptName" db:"dept_name" bze:"4,部门名称"`
	Leader      string  `json:"leader" db:"leader"`
	Email       string  `json:"email" db:"email"`
	Phonenumber string  `json:"phonenumber"db:"phonenumber" bze:"5,电话"`
	Avatar      string  `json:"avatar" db:"avatar"`
	DataScope   string  `json:"dataScope" db:"data_scope"`
	RoleId      *int64  `json:"roleId" db:"role_id"`
	Remark      string  `json:"remark" db:"remark"`
	baize.BaseEntity
}
type SysUserDataScope struct {
	UserId    int64    `json:"userId,string"  binding:"required"` //用户ID
	DataScope string   `json:"dataScope"  binding:"required"`     //数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限,无任何）权限
	DeptIds   []string `json:"deptIds"`                           //如果是自定义就是部门ID 其他不填
}

type ResetPwd struct {
	UserId   int64  `json:"userId,string" db:"user_id" binding:"required"` //用户ID
	Password string `json:"password" db:"password" binding:"required"`     //新密码
}
type EditUserStatus struct {
	UserId int64  `json:"userId,string" binding:"required"` //用户id
	Status string `json:"status" binding:"required"`        //状态
	baize.BaseEntity
}

func RowsToSysUserDMLList(rows [][]string, str string, failureNum int, dept map[string]int64, password string, userId int64) ([]*SysUserDML, string, int) {
	list := make([]*SysUserDML, 0, len(rows)-1)
	for i, row := range rows {
		if i == 0 {
			continue
		}
		if row[0] == "" {
			str += "<br/>第" + strconv.Itoa(i+1) + "行用户名为空"
			failureNum++
			continue
		}
		sysUser := new(SysUserDML)
		sysUser.UserId = snowflake.GenID()
		sysUser.UserName = row[0]
		sysUser.NickName = row[1]
		sysUser.DeptId = dept[row[2]]
		if sysUser.DeptId == 0 {
			str += "<br/>第" + strconv.Itoa(i+1) + "部门错误"
			failureNum++
			continue
		}
		sysUser.Email = row[3]
		sysUser.Phonenumber = row[4]
		sex := row[4]
		if sex == "男" {
			sysUser.Sex = "0"
		} else if sex == "女" {
			sysUser.Sex = "1"
		} else {
			sysUser.Sex = "2"
		}
		sysUser.Status = "0"
		sysUser.Password = password
		sysUser.DataScope = dataScopeAspect.NoDataScope
		sysUser.SetCreateBy(userId)
		list = append(list, sysUser)
	}
	return list, str, failureNum
}

type Accredit struct {
	Posts []*SysPostVo `json:"posts"` //岗位
	Roles []*SysRoleVo `json:"roles"` //角色
}
type UserAndAccredit struct {
	User    *SysUserVo   `json:"user"`    //user
	Roles   []*SysRoleVo `json:"roles"`   //角色
	RoleIds []string     `json:"roleIds"` //选择的角色Id
	Posts   []*SysPostVo `json:"posts"`   //岗位
	PostIds []string     `json:"postIds"` //选择的岗位Id
}
type UserAndRoles struct {
	User    *SysUserVo   `json:"user"`    //user
	Roles   []*SysRoleVo `json:"roles"`   //角色
	RoleIds []string     `json:"roleIds"` //选择的角色Id
}

type UserProfile struct {
	User      *SysUserVo `json:"user"`      //user
	RoleGroup string     `json:"roleGroup"` //角色
	PostGroup string     `json:"postGroup"` //选择的角色Id
}
type SysUserDeptScope struct {
	UserId int64 `db:"user_id"`
	DeptId int64 `db:"dept_id"`
}
