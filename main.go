package main

import (
	"github.com/wolfsblu/go-chef/factories"
	"github.com/wolfsblu/go-chef/infra/env"
	"github.com/wolfsblu/go-chef/infra/routing"
	"log"
	"net/http"
)

func main() {
	env.Load()

	apiServer, err := factories.NewApiServer()
	if err != nil {
		log.Fatalln("failed to start api server:", err)
	}

	mux := http.NewServeMux()
	routing.RegisterApp(mux)
	routing.RegisterApi(mux, apiServer)

	host := env.MustGet("HOST")
	err = http.ListenAndServe(host, mux)
	if err != nil {
		log.Fatalln("failed to start web server:", err)
	}
}
