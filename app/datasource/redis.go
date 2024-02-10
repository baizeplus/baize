package datasource

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var RedisDb *redis.Client

func NewRedis() {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("datasource.redis.host"), viper.GetInt("datasource.redis.port")),
		Password: viper.GetString("datasource.redis.password"),
		DB:       viper.GetInt("datasource.redis.db"),
	})
}
