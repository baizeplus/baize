package monitorModels

import (
	"baize/app/baize"
	"time"
)

type LogininforDQL struct {
	Status   string `form:"status" db:"status"`
	UserName string `form:"userName" db:"user_name"`
	IpAddr   string `form:"ipaddr" db:"ipaddr"`
	baize.BaseEntityDQL
}

type Logininfor struct {
	InfoId        int64     `json:"infoId,string" db:"info_id"`
	Status        int8      `json:"status" db:"status"`
	UserName      string    `json:"userName" db:"user_name"`
	Msg           string    `json:"msg" db:"msg"`
	IpAddr        string    `json:"ipAddr" db:"ipaddr"`
	LoginLocation string    `json:"loginLocation" db:"login_location"`
	Browser       string    `json:"browser" db:"browser"`
	Os            string    `json:"os" db:"os"`
	LoginTime     time.Time `json:"loginTime" db:"login_time"`
}
