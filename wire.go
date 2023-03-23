// wire.go
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rhuandantas/verifymy-test/internal"
	"github.com/rhuandantas/verifymy-test/internal/handlers"
	"github.com/rhuandantas/verifymy-test/internal/repo"
	"github.com/rhuandantas/verifymy-test/internal/server"
	"github.com/rhuandantas/verifymy-test/log"
)

func InitializeWebServer() (*server.HttpServer, error) {
	wire.Build(internal.NewLocalConfigProvider,
		log.NewLogger,
		repo.NewUserRepo,
		handlers.NewUserHandler,
		server.NewAPIServer)
	return &server.HttpServer{}, nil
}
