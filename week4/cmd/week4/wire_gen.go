// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"week4/internal/biz"
	"week4/internal/conf"
	"week4/internal/data"
	"week4/internal/server"
	"week4/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	dramaRepo := data.NewDramaRepo(dataData, logger)
	dramaUserCase := biz.NewDramaUserCase(dramaRepo, logger)
	dramaService := service.NewDramaService(dramaUserCase, logger)
	httpServer := server.NewHTTPServer(confServer, dramaService, logger)
	grpcServer := server.NewGRPCServer(confServer, dramaService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}