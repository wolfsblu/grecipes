package main

import (
	"github.com/swaggest/swgui/v5emb"
	"github.com/wolfsblu/grecipes/api"
	"github.com/wolfsblu/grecipes/db"
	"github.com/wolfsblu/grecipes/env"
	"github.com/wolfsblu/grecipes/routes"
	"github.com/wolfsblu/grecipes/service"
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
	svc := &service.RecipesService{
		Db: query,
	}
	handler, err := api.NewServer(svc)
	if err != nil {
		log.Fatalln("failed to start api server:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/docs/", v5emb.New("OpenAPI Docs", "/api/openapi.yml", "/api/docs/"))
	mux.Handle("/api/", http.StripPrefix("/api", handler))
	routes.Register(mux)

	host := env.MustGet("HOST")
	err = http.ListenAndServe(host, mux)
	if err != nil {
		log.Fatalln("failed to start web server:", err)
	}
}
