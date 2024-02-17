package serviceImpl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUserOnlineService, NewLogininforService)
