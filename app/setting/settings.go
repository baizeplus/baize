package setting

import (
	"fmt"
	"github.com/spf13/viper"
)

var Conf = new(AppConfig)

type AppConfig struct {
	Name         string `mapstructure:"name"`
	Mode         string `mapstructure:"mode"`
	Version      string `mapstructure:"version"`
	StartTime    string `mapstructure:"start_time"`
	Port         int    `mapstructure:"port"`
	Host         string `mapstructure:"host"`
	Cluster      bool   `mapstructure:"cluster"`
	*TokenConfig `mapstructure:"token"`
	*LogConfig   `mapstructure:"log"`
	*Datasource  `mapstructure:"datasource"`
	*UploadFile  `mapstructure:"upload_file"`
}

type TokenConfig struct {
	ExpireTime int64  `mapstructure:"expire_time"`
	Secret     string `mapstructure:"secret"`
	Issuer     string `mapstructure:"issuer"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type Datasource struct {
	Mysql *Mysql `mapstructure:"mysql"`
}
type Mysql struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DB           string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

func init() {

	viper.SetConfigFile("./config/config.yaml")

	err := viper.ReadInConfig() // 读取配置信息
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		panic(err)
		return
	}

	// 把读取到的配置信息反序列化到 Conf 变量中
	if err = viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		panic(err)
	}

	if Conf.Cluster && viper.GetString("cache.type") != "redis" {

		panic("cluster mode must use redis")

	}

	//viper.WatchConfig()   //监视文件更改

	return
}
