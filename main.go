package main

import (
	"fmt"
	"github.com/wolfsblu/go-chef/api"
	"github.com/wolfsblu/go-chef/db"
	"github.com/wolfsblu/go-chef/domain"
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
		log.Fatalln("failed to connect to the database:", err)
	}

	err = db.Migrate(fmt.Sprintf("sqlite://%s", dbPath))
	if err != nil {
		log.Fatalln("failed to apply database migrations:", err)
	}

	store := &db.SqliteStore{DB: query}
	recipeService := domain.NewRecipeService(store)
	rh := handlers.NewRecipeHandler(recipeService)
	sh := handlers.NewSecurityHandler(recipeService)
	apiServer, err := api.NewServer(rh, sh)
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
