// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"baize/app/business/monitor/monitorController"
	"baize/app/business/monitor/monitorDao/monitorDaoImpl"
	"baize/app/business/monitor/monitorService/monitorServiceImpl"
	"baize/app/business/system/systemController"
	"baize/app/business/system/systemDao/systemDaoImpl"
	"baize/app/business/system/systemService/systemServiceImpl"
	"baize/app/datasource"
	"baize/app/routes"
	"baize/app/setting"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func wireApp(settingDatasource *setting.Datasource) (*gin.Engine, func(), error) {
	db, cleanup, err := datasource.NewData(settingDatasource)
	if err != nil {
		return nil, nil, err
	}
	sysUserDao := systemDaoImpl.NewSysUserDao()
	sysMenuDao := systemDaoImpl.NewSysMenuDao()
	sysRoleDao := systemDaoImpl.NewSysRoleDao()
	logininforDao := monitorDaoImpl.NewLogininforDao()
	sysConfigDao := systemDaoImpl.NewSysConfigDao()
	configService := systemServiceImpl.NewConfigService(db, sysConfigDao)
	loginService := systemServiceImpl.NewLoginService(db, sysUserDao, sysMenuDao, sysRoleDao, logininforDao, configService)
	sysUserPostDao := systemDaoImpl.NewSysUserPostDao()
	sysUserRoleDao := systemDaoImpl.NewSysUserRoleDao()
	sysPostDao := systemDaoImpl.NewSysPostDao()
	sysUserDeptScopeDao := systemDaoImpl.NewSysUserDeptScopeDao()
	userService := systemServiceImpl.NewUserService(db, sysUserDao, sysUserPostDao, sysUserRoleDao, sysRoleDao, sysPostDao, sysUserDeptScopeDao)
	sysRoleMenuDao := systemDaoImpl.NewSysRoleMenuDao()
	menuService := systemServiceImpl.NewMenuService(db, sysMenuDao, sysRoleMenuDao, sysRoleDao)
	login := systemController.NewLogin(loginService, userService, menuService, configService)
	postService := systemServiceImpl.NewPostService(db, sysPostDao)
	roleService := systemServiceImpl.NewRoleService(db, sysRoleDao, sysRoleMenuDao, sysUserRoleDao)
	user := systemController.NewUser(userService, postService, roleService)
	sysDeptDao := systemDaoImpl.NewSysDeptDao()
	deptService := systemServiceImpl.NewDeptService(db, sysDeptDao, sysRoleDao)
	dept := systemController.NewDept(deptService)
	sysDictTypeDao := systemDaoImpl.NewSysDictTypeDao()
	dictTypeService := systemServiceImpl.NewDictTypeService(db, sysDictTypeDao)
	sysDictDataDao := systemDaoImpl.NewSysDictDataDao()
	dictDataService := systemServiceImpl.NewDictDataService(db, sysDictDataDao)
	dictType := systemController.NewDictType(dictTypeService, dictDataService)
	dictData := systemController.NewDictData(dictDataService)
	menu := systemController.NewMenu(menuService)
	role := systemController.NewRole(roleService)
	post := systemController.NewPost(postService)
	profile := systemController.NewProfile(userService)
	config := systemController.NewConfig(configService)
	infoServer := monitorController.NewInfoServer()
	userOnlineService := monitorServiceImpl.NewUserOnlineService()
	userOnline := monitorController.NewUserOnline(userOnlineService)
	logininforService := monitorServiceImpl.NewLogininforService(db, logininforDao)
	logininfor := monitorController.NewLogininfor(logininforService)
	engine := routes.NewGinEngine(login, user, dept, dictType, dictData, menu, role, post, profile, config, infoServer, userOnline, logininfor)
	return engine, func() {
		cleanup()
	}, nil
}
