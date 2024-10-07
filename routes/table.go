package routes

import "net/http"

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/public/", Assets)
	mux.HandleFunc("/recipes/create/", CreateRecipe)
	mux.HandleFunc("/recipes/", GetRecipes)
	mux.HandleFunc("/{$}", Index)
}
