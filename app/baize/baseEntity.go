package baize

import (
	"baize/app/utils/stringUtils"
	"github.com/baizeplus/sqly"
)

type BaseEntity struct {
	CreateBy   int64 `json:"createBy" db:"create_by"`                           //创建人
	CreateTime *Time `json:"createTime" db:"create_time" swaggertype:"integer"` //创建时间
	UpdateBy   int64 `json:"updateBy" db:"update_by"`                           //修改人
	UpdateTime *Time `json:"updateTime" db:"update_time" swaggertype:"integer"` //修改时间
}

func (b *BaseEntity) SetCreateBy(userId int64) {
	b.CreateBy = userId
	kt := NewTime()
	b.CreateTime = kt
	b.UpdateBy = userId
	b.UpdateTime = kt
}

func (b *BaseEntity) SetUpdateBy(userId int64) {
	b.UpdateBy = userId
	b.UpdateTime = NewTime()
}

type BaseEntityDQL struct {
	DataScope string `swaggerignore:"true"`
	OrderBy   string `form:"orderBy" `              //排序字段
	IsAsc     string `form:"isAsc" `                //排序规则  降序desc   asc升序
	Page      int64  `form:"pageNum" default:"1"`   //第几页
	Size      int64  `form:"pageSize" default:"10"` //数量
}

func (b *BaseEntityDQL) GetOrder() string {
	if b.OrderBy != "" {
		return " " + stringUtils.ToUnderScoreCase(b.OrderBy) + " " + b.IsAsc
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
