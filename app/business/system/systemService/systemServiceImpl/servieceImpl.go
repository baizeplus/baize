package systemServiceImpl

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewLoginService, NewSseService, NewNoticeService, NewUserService, NewMenuService, NewRoleService, NewPostService, NewDeptService, NewDictTypeService, NewDictDataService, NewConfigService, NewFileService)
