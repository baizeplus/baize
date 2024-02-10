package serviceImpl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewLoginService, NewUserService, NewMenuService, NewRoleService, NewPostService)
