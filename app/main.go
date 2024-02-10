package main

import (
	"baize/app/setting"
	"fmt"
)

// @title baize
// @version 2.0.x
// @description baize接口文档

// @contact.name danny
// @contact.email zhao_402295440@126.com

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app, cleanup, err := wireApp(setting.Conf.Datasource)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	app.Run(fmt.Sprintf(":%d", setting.Conf.Port))

}
