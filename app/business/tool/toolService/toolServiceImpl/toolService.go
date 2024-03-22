package toolServiceImpl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewGenTabletService)
