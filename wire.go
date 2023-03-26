// wire.go
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rhuandantas/verifymy-test/internal/config"
	"github.com/rhuandantas/verifymy-test/internal/log"
	"github.com/rhuandantas/verifymy-test/internal/repo"
	"github.com/rhuandantas/verifymy-test/internal/server"
	"github.com/rhuandantas/verifymy-test/internal/server/handlers"
	"github.com/rhuandantas/verifymy-test/internal/server/middlewares/auth"
	"github.com/rhuandantas/verifymy-test/internal/util"
)

func InitializeWebServer() (*server.HttpServer, error) {
	wire.Build(config.NewLocalConfigProvider,
		util.NewCustomValidator,
		log.NewLogger,
		repo.NewMysqlORMConn,
		auth.NewJwtToken,
		repo.NewUserRepo,
		handlers.NewUserHandler,
		handlers.NewHealthCheck,
		server.NewAPIServer)
	return &server.HttpServer{}, nil
}
