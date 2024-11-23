package main

import (
	"github.com/wolfsblu/go-chef/api"
	"github.com/wolfsblu/go-chef/db"
	"github.com/wolfsblu/go-chef/env"
	"github.com/wolfsblu/go-chef/handlers"
	"github.com/wolfsblu/go-chef/routes"
	"log"
	"net/http"
)

func main() {
	env.Load()

	dbPath := env.MustGet("DB_PATH")
	query, err := db.Connect(dbPath)
	if err != nil {
		log.Fatalln("failed to connect to the database", err)
	}

	recipes := handlers.NewRecipeHandler(query)
	security := handlers.NewSecurityHandler(query)
	apiServer, err := api.NewServer(recipes, security)
	if err != nil {
		log.Fatalln("failed to start api server:", err)
	}

	mux := http.NewServeMux()
	routes.RegisterApp(mux)
	routes.RegisterApi(mux, apiServer)

	host := env.MustGet("HOST")
	err = http.ListenAndServe(host, mux)
	if err != nil {
		log.Fatalln("failed to start web server:", err)
	}
}
