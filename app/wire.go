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
	"baize/app/datasource"
	"baize/app/routes"
	"baize/app/setting"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func wireApp(*setting.Datasource) (*gin.Engine, func(), error) {
	panic(wire.Build(
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
