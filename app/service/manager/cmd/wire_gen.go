// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"aurora/blog/api/app/service/manager/internal/biz"
	"aurora/blog/api/app/service/manager/internal/conf"
	"aurora/blog/api/app/service/manager/internal/data"
	"aurora/blog/api/app/service/manager/internal/server"
	"aurora/blog/api/app/service/manager/internal/service"
	"aurora/blog/api/pkg/lifecycle"
)

// Injectors from wire.go:

func initApp() (*lifecycle.App, func(), error) {
	logger := conf.NewZapLogger()
	configData := conf.NewDataConfig()
	dataData, cleanup, err := data.NewData(logger, configData)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userBusiness := biz.NewUserBusiness(userRepo)
	userService := service.NewUserService(userBusiness, logger)
	engine := server.NewHttpRouter(userService, logger)
	configServer := conf.NewServerConfig()
	transportServer := server.NewHttpServer(engine, logger, configServer)
	app := newApp(transportServer, logger)
	return app, func() {
		cleanup()
	}, nil
}
