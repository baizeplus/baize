package monitorController

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewInfoServer, NewUserOnline, NewLogininfor, NewOperLog, wire.Struct(new(Monitor), "*"))

type Monitor struct {
	Server     *InfoServer
	UserOnline *UserOnline
	Logfor     *Logininfor
	Oper       *OperLog
}
