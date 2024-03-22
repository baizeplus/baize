package toolController

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewGenTable, wire.Struct(new(Tool), "*"))

type Tool struct {
	GenTable *GenTable
}
