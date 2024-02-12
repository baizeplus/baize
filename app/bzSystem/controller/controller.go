package controller

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewUserController, NewRoleController, NewPostController, NewMenuController, NewLoginController, NewDictTypeController,
	NewProfileController, NewDictDataController, NewDeptController)
