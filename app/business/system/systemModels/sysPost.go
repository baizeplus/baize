package systemModels

import "baize/app/baize"

type SysPostDQL struct {
	PostCode string `form:"postCode" db:"post_code"`
	Status   string `form:"status" db:"status"`
	PostName string `form:"postName" db:"post_name"`
	baize.BaseEntityDQL
}

type SysPostVo struct {
	PostId   int64  `json:"postId,string" db:"post_id"`
	PostSort int64  `json:"postSort" db:"post_sort"`
	PostCode string `json:"postCode" db:"post_code"`
	PostName string `json:"postName" db:"post_name"`
	Status   string `json:"status" db:"status"`
	Remark   string `json:"remark" db:"remark"`
	baize.BaseEntity
}
