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
	"baize/app/business/tool/toolController"
	"baize/app/business/tool/toolDao/toolDaoImpl"
	"baize/app/business/tool/toolService/toolServiceImpl"
	"baize/app/datasource/cache"
	"baize/app/datasource/mysql"
	"baize/app/datasource/objectFile"
	"baize/app/routes"
	"github.com/gin-gonic/gin"
)

// Injectors from wire.go:

func wireApp() (*gin.Engine, func(), error) {
	cacheCache := cache.NewCache()
	sqlyContext, cleanup, err := mysql.NewData()
	if err != nil {
		return nil, nil, err
	}
	iUserDao := systemDaoImpl.NewSysUserDao(sqlyContext)
	iPermissionDao := systemDaoImpl.NewSysPermissionDao(sqlyContext)
	iRoleDao := systemDaoImpl.NewSysRoleDao(sqlyContext)
	iLogininforDao := monitorDaoImpl.NewLogininforDao(sqlyContext)
	iConfigDao := systemDaoImpl.NewSysConfigDao(sqlyContext)
	iConfigService := systemServiceImpl.NewConfigService(iConfigDao, cacheCache)
	iLoginService := systemServiceImpl.NewLoginService(cacheCache, iUserDao, iPermissionDao, iRoleDao, iLogininforDao, iConfigService)
	iUserPostDao := systemDaoImpl.NewSysUserPostDao(sqlyContext)
	iUserRoleDao := systemDaoImpl.NewSysUserRoleDao(sqlyContext)
	objectFileObjectFile := objectFile.NewConfig()
	iDeptDao := systemDaoImpl.NewSysDeptDao(sqlyContext)
	iPostDao := systemDaoImpl.NewSysPostDao(sqlyContext)
	iUserDeptScopeDao := systemDaoImpl.NewSysUserDeptScopeDao(sqlyContext)
	iUserService := systemServiceImpl.NewUserService(sqlyContext, iUserDao, iUserPostDao, iUserRoleDao, objectFileObjectFile, iDeptDao, iRoleDao, iPostDao, iUserDeptScopeDao, iConfigService)
	iLogininforService := monitorServiceImpl.NewLogininforService(iLogininforDao)
	login := systemController.NewLogin(iLoginService, iUserService, iConfigService, iLogininforService)
	iPostService := systemServiceImpl.NewPostService(iPostDao)
	iRolePermissionDao := systemDaoImpl.NewSysRolePermissionDao(sqlyContext)
	iRoleService := systemServiceImpl.NewRoleService(sqlyContext, iRoleDao, iRolePermissionDao, iUserRoleDao)
	user := systemController.NewUser(iUserService, iPostService, iRoleService)
	iDeptService := systemServiceImpl.NewDeptService(iDeptDao, iRoleDao)
	dept := systemController.NewDept(iDeptService)
	iDictTypeDao := systemDaoImpl.NewSysDictTypeDao(sqlyContext)
	iDictTypeService := systemServiceImpl.NewDictTypeService(iDictTypeDao, cacheCache)
	iDictDataDao := systemDaoImpl.NewSysDictDataDao(sqlyContext)
	iDictDataService := systemServiceImpl.NewDictDataService(iDictDataDao, cacheCache)
	dictType := systemController.NewDictType(iDictTypeService, iDictDataService)
	dictData := systemController.NewDictData(iDictDataService)
	role := systemController.NewRole(iRoleService)
	post := systemController.NewPost(iPostService)
	profile := systemController.NewProfile(iUserService)
	config := systemController.NewConfig(iConfigService)
	iFileService := systemServiceImpl.NewFileService(objectFileObjectFile)
	file := systemController.NewFile(iFileService)
	iSseService := systemServiceImpl.NewSseService(cacheCache)
	sse := systemController.NewSse(iSseService)
	iSysNoticeDao := systemDaoImpl.NewSysNoticeDao(sqlyContext)
	iSysNoticeService := systemServiceImpl.NewNoticeService(iSysNoticeDao, iUserDao, iSseService)
	notice := systemController.NewNotice(iSysNoticeService)
	iSysPermissionService := systemServiceImpl.NewPermissionService(iPermissionDao)
	permission := systemController.NewPermission(iSysPermissionService)
	iSelectBoxService := systemServiceImpl.NewSelectService(iPermissionDao, iDeptDao)
	selectBox := systemController.NewSelectBox(iSelectBoxService)
	system := &systemController.System{
		Login:      login,
		User:       user,
		Dept:       dept,
		DictType:   dictType,
		DictData:   dictData,
		Role:       role,
		Post:       post,
		Profile:    profile,
		Config:     config,
		File:       file,
		Sse:        sse,
		Notice:     notice,
		Permission: permission,
		SelectBox:  selectBox,
	}
	infoServer := monitorController.NewInfoServer()
	iUserOnlineService := monitorServiceImpl.NewUserOnlineService(cacheCache)
	userOnline := monitorController.NewUserOnline(iUserOnlineService)
	logininfor := monitorController.NewLogininfor(iLogininforService)
	iOperLog := monitorDaoImpl.NewOperLog(sqlyContext)
	iSysOperLogService := monitorServiceImpl.NewOperLog(iOperLog)
	operLog := monitorController.NewOperLog(iSysOperLogService)
	iJobDao := monitorDaoImpl.NewJobDao(sqlyContext)
	iJobService := monitorServiceImpl.NewJobService(cacheCache, iJobDao)
	job := monitorController.NewJob(iJobService)
	monitor := &monitorController.Monitor{
		Server:     infoServer,
		UserOnline: userOnline,
		Logfor:     logininfor,
		Oper:       operLog,
		Job:        job,
	}
	iGenTableColumn := toolDaoImpl.NewGenTableColumnDao(sqlyContext)
	iGenTable := toolDaoImpl.GetGenTableDao(sqlyContext)
	iGenTableService := toolServiceImpl.NewGenTabletService(iGenTableColumn, iGenTable)
	genTable := toolController.NewGenTable(iGenTableService)
	tool := &toolController.Tool{
		GenTable: genTable,
	}
	engine := routes.NewGinEngine(cacheCache, sqlyContext, system, monitor, tool)
	return engine, func() {
		cleanup()
	}, nil
}
