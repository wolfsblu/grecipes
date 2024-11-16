package main

import (
	"github.com/swaggest/swgui/v5emb"
	"github.com/wolfsblu/grecipes/api"
	"github.com/wolfsblu/grecipes/db"
	"github.com/wolfsblu/grecipes/routes"
	"github.com/wolfsblu/grecipes/service"
	"log"
	"net/http"
	"os"
)

func main() {
	dbPath, ok := os.LookupEnv("DB_PATH")
	if !ok {
		log.Fatalln("env variable DB_PATH needs to point to a sqlite database path")
	}
	query, err := db.Connect(dbPath)
	if err != nil {
		log.Fatalln("could not connect to database", err)
	}
	svc := &service.RecipesService{
		Db: query,
	}
	handler, err := api.NewServer(svc)
	if err != nil {
		log.Fatalln("failed to start API server:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/docs/", v5emb.New("OpenAPI Docs", "/api/openapi.yml", "/api/docs/"))
	mux.Handle("/api/", http.StripPrefix("/api", handler))
	routes.Register(mux)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln("failed to start server:", err)
	}
}
