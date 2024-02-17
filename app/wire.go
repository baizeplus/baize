//go:build wireinject
// +build wireinject

package main

import (
	controllerM "baize/app/bzMonitor/controller"
	daoImplM "baize/app/bzMonitor/dao/daoImpl"
	serviceImplM "baize/app/bzMonitor/service/serviceImpl"
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
		daoImplM.ProviderSet,
		daoImpl.ProviderSet,
		datasource.ProviderSet,
		serviceImplM.ProviderSet,
		serviceImpl.ProviderSet,
		controller.ProviderSet,
		controllerM.ProviderSet,
		routes.ProviderSet,
	))
}
