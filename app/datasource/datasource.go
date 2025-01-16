package datasource

import (
	"baize/app/datasource/cache"
	"baize/app/datasource/mysql"
	"baize/app/datasource/objectFile"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(mysql.NewData, objectFile.NewConfig, cache.NewCache, cache.NewRedisSubscribe)
