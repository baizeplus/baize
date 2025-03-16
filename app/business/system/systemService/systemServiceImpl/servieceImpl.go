package systemServiceImpl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewLoginService, NewSseService, NewNoticeService, NewUserService, NewPermissionService, NewRoleService, NewPostService,
	NewSelectService, NewDeptService, NewDictTypeService, NewDictDataService, NewConfigService, NewFileService)
