package main

//go:generate wire
func main() {
	server, err := InitializeWebServer()
	if err != nil {
		panic(err)
	}

	server.Start()
}
