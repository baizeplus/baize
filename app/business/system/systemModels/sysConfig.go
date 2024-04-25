package systemModels

import "baize/app/baize"

type SysConfigDQL struct {
	ConfigName string `form:"configName" db:"config_name"` //参数名称
	ConfigKey  string `form:"configKey" db:"config_key"`   //参数键名
	ConfigType string `form:"configType" db:"config_type"` //系统内置（Y是 N否）
	baize.BaseEntityDQL
}

type SysConfigVo struct {
	ConfigId    int64  `json:"configId,string" db:"config_id"`             //参数主键
	ConfigName  string `json:"configName" db:"config_name" bze:"1,参数名称"`   //参数名称
	ConfigKey   string `json:"configKey" db:"config_key" bze:"2,参数键名"`     //参数键名
	ConfigValue string `json:"configValue" db:"config_value" bze:"3,参数键值"` //参数键值
	ConfigType  string `json:"configType" db:"config_type" bze:"4,系统内置"`   //系统内置（Y是 N否）
	Remark      string `json:"remark" db:"remark" bze:"5,备注"`              //备注
	baize.BaseEntity
}
