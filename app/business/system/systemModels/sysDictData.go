package systemModels

import "baize/app/baize"

type SysDictDataVo struct {
	DictCode  int64  `json:"dictCode,string" db:"dict_code"`
	DictSort  int32  `json:"dictSort" db:"dict_sort"`
	DictLabel string `json:"dictLabel" db:"dict_label"  bze:"1,字典标签"`
	DictValue string `json:"dictValue" db:"dict_value"  bze:"2,参数键值"`
	DictType  string `json:"dictType" db:"dict_type"  bze:"3,字典类型"`
	CssClass  string `json:"cssClass" db:"css_class"`
	ListClass string `json:"listClass" db:"list_class"`
	IsDefault string `json:"isDefault" db:"is_default"`
	Status    string `json:"status" db:"status" bze:"4,状态"`
	Remark    string `json:"remark" db:"remark" bze:"5,备注"`
	baize.BaseEntity
}

type SysDictDataDQL struct {
	DictType  string `form:"dictType" db:"dict_type"`
	DictLabel string `form:"dictLabel" db:"dict_label"`
	Status    string `form:"status" db:"status"`
	baize.BaseEntityDQL
}
