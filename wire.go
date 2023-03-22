//go:build wireinject
// +build wireinject

// wire.go

package main

import "github.com/google/wire"

func InitializeWebServer() (*WebServer, error) {
	wire.Build(NewWebServer)
	return &WebServer{}, nil
}
