package datasource

import (
	"baize/app/setting"
	"fmt"

	"github.com/baizeplus/sqly"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

// ProviderSet is datasource providers.
var ProviderSet = wire.NewSet(NewData)

// NewData .
func NewData(d *setting.Datasource) (*sqly.DB, func(), error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", d.Mysql.User, d.Mysql.Password, d.Mysql.Host, d.Mysql.Port, d.Mysql.DB)
	db, err := sqly.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(d.Mysql.MaxOpenConns)
	db.SetMaxIdleConns(d.Mysql.MaxIdleConns)
	return db, nil, err
}
