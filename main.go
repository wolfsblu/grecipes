package main

import (
	"github.com/wolfsblu/grecipes/routes"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/public/", routes.Assets)
	mux.HandleFunc("/recipes/create/", routes.CreateRecipe)
	mux.HandleFunc("/recipes/", routes.GetRecipes)
	mux.HandleFunc("/{$}", routes.Index)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln("failed to start server:", err)
	}
}
