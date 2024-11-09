package main

import (
	"github.com/swaggest/swgui/v5emb"
	"github.com/wolfsblu/grecipes/api"
	"github.com/wolfsblu/grecipes/routes"
	"github.com/wolfsblu/grecipes/service"
	"log"
	"net/http"
)

func main() {
	svc := &service.RecipesService{
		Recipes: map[int64]api.Recipe{},
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
