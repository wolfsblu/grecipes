package main

import (
	"fmt"
	"github.com/wolfsblu/go-chef/api"
	"github.com/wolfsblu/go-chef/domain"
	"github.com/wolfsblu/go-chef/infra/env"
	"github.com/wolfsblu/go-chef/infra/handlers"
	"github.com/wolfsblu/go-chef/infra/routing"
	"github.com/wolfsblu/go-chef/infra/smtp"
	"github.com/wolfsblu/go-chef/infra/sqlite"
	"log"
	"net/http"
)

func main() {
	env.Load()

	dbPath := env.MustGet("DB_PATH")
	query, err := sqlite.Connect(dbPath)
	if err != nil {
		log.Fatalln("failed to connect to the database:", err)
	}

	err = sqlite.Migrate(fmt.Sprintf("sqlite://%s", dbPath))
	if err != nil {
		log.Fatalln("failed to apply database migrations:", err)
	}

	recipeService := domain.NewRecipeService(
		&sqlite.Store{DB: query},
		&smtp.Mailer{},
	)
	rh := handlers.NewRecipeHandler(recipeService)
	sh := handlers.NewSecurityHandler(recipeService)
	apiServer, err := api.NewServer(rh, sh)
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
