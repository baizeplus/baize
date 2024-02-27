package systemController

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUser, NewRole, NewPost, NewMenu, NewLogin, NewDictType,
	NewProfile, NewDictData, NewDept)