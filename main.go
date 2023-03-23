package main

//go:generate wire
func main() {
	server, err := InitializeWebServer()
	server.RegisterHandlers()
	if err != nil {
		panic(err)
	}
	server.Start()
}
