package controller

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewInfoServer, NewUserOnline, NewLogininfor)
