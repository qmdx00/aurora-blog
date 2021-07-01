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
	logger := conf.NewLogger()
	repo, cleanup, err := data.NewData(logger)
	if err != nil {
		return nil, nil, err
	}
	bizBiz := biz.NewBiz(repo)
	serviceService := service.NewService(bizBiz)
	engine := server.NewHttpRouter(serviceService, logger)
	transportServer := server.NewHttpServer(engine, logger)
	app := newApp(transportServer, logger)
	return app, func() {
		cleanup()
	}, nil
}
