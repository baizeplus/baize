// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"baize/app/bzMonitor/dao/monitorDaoImpl"
	"baize/app/bzSystem/controller"
	"baize/app/bzSystem/dao/daoImpl"
	"baize/app/bzSystem/service/serviceImpl"
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
	sysUserDao := daoImpl.NewSysUserDao()
	sysMenuDao := daoImpl.NewSysMenuDao()
	sysRoleDao := daoImpl.NewSysRoleDao()
	logininforDao := monitorDaoImpl.NewLogininforDao()
	loginService := serviceImpl.NewLoginService(db, sysUserDao, sysMenuDao, sysRoleDao, logininforDao)
	sysUserPostDao := daoImpl.NewSysUserPostDao()
	sysUserRoleDao := daoImpl.NewSysUserRoleDao()
	sysPostDao := daoImpl.NewSysPostDao()
	userService := serviceImpl.NewUserService(db, sysUserDao, sysUserPostDao, sysUserRoleDao, sysRoleDao, sysPostDao)
	sysRoleMenuDao := daoImpl.NewSysRoleMenuDao()
	menuService := serviceImpl.NewMenuService(db, sysMenuDao, sysRoleMenuDao, sysRoleDao)
	loginController := controller.NewLoginController(loginService, userService, menuService)
	postService := serviceImpl.NewPostService(db, sysPostDao)
	sysRoleDeptDao := daoImpl.NewSysRoleDeptDao()
	roleService := serviceImpl.NewRoleService(db, sysRoleDao, sysRoleMenuDao, sysRoleDeptDao, sysUserRoleDao)
	userController := controller.NewUserController(userService, postService, roleService)
	dictDataController := controller.NewDictDataController()
	sysDictTypeDao := daoImpl.NewSysDictTypeDao()
	dictTypeService := serviceImpl.NewDictTypeService(db, sysDictTypeDao)
	dictTypeController := controller.NewDictTypeController(dictTypeService)
	engine := routes.NewGinEngine(loginController, userController, dictDataController, dictTypeController)
	return engine, func() {
		cleanup()
	}, nil
}
