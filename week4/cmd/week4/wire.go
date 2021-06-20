// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"week4/internal/biz"
	"week4/internal/conf"
	"week4/internal/data"
	"week4/internal/server"
	"week4/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderServerSet, data.ProviderDataSet, biz.ProviderBizSet, service.ProviderSetService, newApp))
}
