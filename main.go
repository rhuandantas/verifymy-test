package main

import (
	"os"
	"os/signal"
	"syscall"
)

//go:generate wire

//	@title			VerifyMy API
//	@version		3.0.0
//	@description	This is a documentation of all endpoints in the API.

//	@host		localhost:3000
//	@BasePath	/
//  @schemes http
//  @produce json
//	@securityDefinitions.apikey	JWT
//	@in							header
//	@name						token
func main() {
	server, err := InitializeWebServer()
	server.RegisterHandlers()
	if err != nil {
		server.Server.Logger.Error(err.Error())
		panic(err)
	}
	server.Start()

	// listens for system signals to gracefully shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	switch <-signalChannel {
	case os.Interrupt:
		server.Server.Logger.Info("Received SIGINT, stopping...")
	case syscall.SIGTERM:
		server.Server.Logger.Info("Received SIGINT, stopping...")
	}
}
