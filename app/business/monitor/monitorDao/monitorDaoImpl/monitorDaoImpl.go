package monitorDaoImpl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewLogininforDao, NewOperLog)
