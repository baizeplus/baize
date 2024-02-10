package baize

import (
	"baize/app/utils/stringUtils"
	"github.com/baizeplus/sqly"
)

type BaseEntity struct {
	CreateBy   string `json:"createBy" db:"create_by"`                           //创建人
	CreateTime *Time  `json:"createTime" db:"create_time" swaggertype:"integer"` //创建时间
	UpdateBy   string `json:"updateBy" db:"update_by"`                           //修改人
	UpdateTime *Time  `json:"updateTime" db:"update_time" swaggertype:"integer"` //修改时间
}

func (b *BaseEntity) SetCreateBy(userName string) {
	b.CreateBy = userName
	kt := NewTime()
	b.CreateTime = kt
	b.UpdateBy = userName
	b.UpdateTime = kt
}

func (b *BaseEntity) SetUpdateBy(userName string) {
	b.UpdateBy = userName
	b.UpdateTime = NewTime()
}

type BaseEntityDQL struct {
	DataScope string `swaggerignore:"true"`
	OrderBy   string `form:"orderBy" `          //排序字段
	IsAsc     string `form:"isAsc" `            //排序规则  降序desc   asc升序
	Page      int64  `form:"page" default:"1"`  //第几页
	Size      int64  `form:"size" default:"10"` //数量
}

func (b *BaseEntityDQL) GetOrder() string {
	if b.OrderBy != "" {
		return " order by " + stringUtils.ToUnderScoreCase(b.OrderBy) + " " + b.IsAsc
	}
	return ""
}
func (b *BaseEntityDQL) ToPage() *sqly.Page {
	s := new(sqly.Page)
	s.Page = b.Page
	s.Size = b.Size
	s.OrderBy = b.GetOrder()
	return s
}

type User interface {
	GetRoles() []*Role
	GetDeptId() int64
}
type Role struct {
	RoleId    int64  `db:"role_id"`
	DataScope string `db:"data_scope"`
}
