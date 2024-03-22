package toolDaoImpl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewGenTableColumnDao, GetGenTableDao)
