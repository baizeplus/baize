package mysql

import (
	"baize/app/utils/logger"
	"fmt"
	"github.com/baizeplus/sqly"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"time"
)

//type Sql interface {
//	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
//	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
//	NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqly.Rows, error)
//	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
//	NamedSelectContext(ctx context.Context, dest interface{}, query string, arg interface{}) error
//	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
//	NamedGetContext(ctx context.Context, dest interface{}, query string, args interface{}) error
//	NamedSelectPageContext(ctx context.Context, dest interface{}, total *int64, query string, page sqly.Page) error
//	MustBegin() *sqly.Tx
//}

// NewData .
func NewData() (sqly.SqlyContext, func(), error) {
	type Mysql struct {
		Host         string `mapstructure:"host"`
		User         string `mapstructure:"user"`
		Password     string `mapstructure:"password"`
		DB           string `mapstructure:"dbname"`
		Port         int    `mapstructure:"port"`
		MaxOpenConns int    `mapstructure:"max_open_conns"`
		MaxIdleConns int    `mapstructure:"max_idle_conns"`
	}
	// 把读取到的配置信息反序列化到 Conf 变量中
	var d Mysql
	if err := viper.UnmarshalKey("mysql", &d); err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", d.User, d.Password, d.Host, d.Port, d.DB) + "?parseTime=true&loc=Asia%2FShanghai"
	db, err := sqly.Connect("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(d.MaxOpenConns)
	db.SetMaxIdleConns(d.MaxIdleConns)
	db.SetConnMaxLifetime(time.Minute * 5)
	sqly.SetLog(new(logger.SqlyLog))
	return db, nil, err
}
