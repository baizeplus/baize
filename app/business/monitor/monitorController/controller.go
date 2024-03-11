package monitorController

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewInfoServer, NewUserOnline, NewLogininfor, NewOperLog)
