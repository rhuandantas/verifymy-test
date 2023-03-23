package main

//go:generate wire
func main() {
	server, err := InitializeWebServer()
	server.RegisterHandlers()
	if err != nil {
		server.Server.Logger.Error(err.Error())
		panic(err)
	}
	server.Start()
}
