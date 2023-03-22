// wire.go
//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rhuandantas/verifymy-test/internal/server"
)

func InitializeWebServer() (*server.HttpServer, error) {
	wire.Build(server.NewAPIServer)
	return &server.HttpServer{}, nil
}
