package baize

import (
	"baize/app/utils/stringUtils"
	"time"
)

type BaseEntity struct {
	CreateBy   int64      `json:"createBy" db:"create_by"`                           //创建人
	CreateTime *time.Time `json:"createTime" db:"create_time" swaggertype:"integer"` //创建时间
	UpdateBy   int64      `json:"updateBy" db:"update_by"`                           //修改人
	UpdateTime *time.Time `json:"updateTime" db:"update_time" swaggertype:"integer"` //修改时间
}

func (b *BaseEntity) SetCreateBy(userId int64) {
	b.CreateBy = userId
	//kt := NewTime()
	kt := time.Now()
	b.CreateTime = &kt
	b.UpdateBy = userId
	b.UpdateTime = &kt
}

func (b *BaseEntity) SetUpdateBy(userId int64) {
	b.UpdateBy = userId
	kt := time.Now()
	b.UpdateTime = &kt
}

type BaseEntityDQL struct {
	DataScope string `swaggerignore:"true"`
	OrderBy   string `form:"orderBy" `                 //排序字段
	IsAsc     string `form:"isAsc" `                   //排序规则  降序desc   asc升序
	Page      int64  `form:"pageNum" default:"1"`      //第几页
	Size      int64  `form:"pageSize" default:"10000"` //数量
}

func (b *BaseEntityDQL) GetOrderBy() string {
	if b.OrderBy != "" {
		return " " + stringUtils.ToUnderScoreCase(b.OrderBy) + " " + b.IsAsc
	}
	return ""
}
func (b *BaseEntityDQL) GetSize() int64 {
	if b.Size < 1 {
		b.Size = 10
	}
	if b.Size > 10000 {
		b.Size = 10000
	}
	return b.Size
}
func (b *BaseEntityDQL) GetPage() int64 {
	if b.Page < 1 {
		return 1
	}
	return b.Page
}
