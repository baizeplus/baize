package systemModels

import "baize/app/baize"

type SysDictTypeVo struct {
	DictId   int64  `json:"dictId,string" db:"dict_id"`
	DictName string `json:"dictName" db:"dict_name"  bze:"1,字典名称"`
	DictType string `json:"dictType" db:"dict_type"  bze:"2,字典类型"`
	Status   string `json:"status" db:"status"  bze:"3,状态"`
	Remark   string `json:"remark" db:"remark"  bze:"4,备注"`
	baize.BaseEntity
}

type SysDictTypeDQL struct {
	DictName string `form:"dictName" db:"dict_name"`
	Status   string `form:"status" db:"status"`
	DictType string `form:"dictType" db:"dict_type"`
	baize.BaseEntityDQL
}
