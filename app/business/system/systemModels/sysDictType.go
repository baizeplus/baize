package systemModels

import "baize/app/baize"

type SysDictTypeVo struct {
	DictId   int64  `json:"dictId,string" db:"dict_id"`
	DictName string `json:"dictName" db:"dict_name"`
	DictType string `json:"dictType" db:"dict_type"`
	Status   string `json:"status" db:"status"`
	Remark   string `json:"remark" db:"remark"`
	baize.BaseEntity
}

type SysDictTypeDQL struct {
	DictName string `form:"dictName" db:"dict_name"`
	Status   string `form:"status" db:"status"`
	DictType string `form:"dictType" db:"dict_type"`
	baize.BaseEntityDQL
}
