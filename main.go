package main

import (
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("failed to load .env file", err)
	}

	dbPath := getEnv("DB_PATH", "tmp/db.sqlite")
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

	host := getEnv("HOST", "127.0.0.1:8080")
	err = http.ListenAndServe(host, mux)
	if err != nil {
		log.Fatalln("failed to start web server:", err)
	}
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		value = fallback
	}
	return value
}
