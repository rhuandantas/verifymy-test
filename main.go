package main

import (
	"os"
	"os/signal"
	"syscall"
)

//go:generate wire
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
