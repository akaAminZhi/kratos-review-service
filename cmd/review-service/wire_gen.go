// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"review-service/internal/biz"
	"review-service/internal/conf"
	"review-service/internal/data"
	"review-service/internal/server"
	"review-service/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewDB(confData)
	dataData, cleanup, err := data.NewData(db, logger)
	if err != nil {
		return nil, nil, err
	}
	reviewRepo := data.NewReviewRepo(dataData, logger)
	reviewUsecase := biz.NewReviewUsecase(reviewRepo, logger)
	reviewService := service.NewReviewService(reviewUsecase)
	grpcServer := server.NewGRPCServer(confServer, reviewService, logger)
	httpServer := server.NewHTTPServer(confServer, reviewService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
