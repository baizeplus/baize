//go:build wireinject
// +build wireinject

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
	"baize/app/datasource"
	"baize/app/routes"
	"baize/app/setting"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func wireApp(*setting.Datasource) (*gin.Engine, func(), error) {
	panic(wire.Build(
		toolDaoImpl.ProviderSet,
		toolServiceImpl.ProviderSet,
		toolController.ProviderSet,
		systemDaoImpl.ProviderSet,
		systemServiceImpl.ProviderSet,
		systemController.ProviderSet,
		monitorDaoImpl.ProviderSet,
		monitorServiceImpl.ProviderSet,
		monitorController.ProviderSet,
		datasource.ProviderSet,
		routes.ProviderSet,
	))
}
