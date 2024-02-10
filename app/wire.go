//go:build wireinject
// +build wireinject

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
	"github.com/google/wire"
)

func wireApp(*setting.Datasource) (*gin.Engine, func(), error) {
	panic(wire.Build(
		monitorDaoImpl.ProviderSet,
		daoImpl.ProviderSet,
		datasource.ProviderSet,
		serviceImpl.ProviderSet,
		controller.ProviderSet,
		routes.ProviderSet,
	))
}
