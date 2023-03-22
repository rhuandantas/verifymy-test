package main

type WebServer struct {
}

func NewWebServer() (*WebServer, error) {
	return &WebServer{}, nil
}

func (ws *WebServer) Start() {

}

func main() {
	server, err := InitializeWebServer()
	if err != nil {
		panic(err)
	}

	server.Start()
}
