package main

import (
	"github.com/KevinAnthony/minerva-client/config"
	"github.com/KevinAnthony/minerva-client/server"

	"github.com/kevinanthony/gorps/v2/http"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	reqh := http.NewRequestHandler(http.NewRequestHandlerHelper())
	_ = cfg

	server.NewServer(reqh).Run()
}
