package routes

import (
	"context"
	"github.com/wolfsblu/grecipes/db"
	"github.com/wolfsblu/grecipes/templates/components"
	"github.com/wolfsblu/grecipes/templates/pages"
	"log"
	"net/http"
)

func CreateRecipe(w http.ResponseWriter, _ *http.Request) {
	c := pages.Create()
	_ = c.Render(context.Background(), w)
}

func GetRecipes(w http.ResponseWriter, _ *http.Request) {
	recipes, err := db.Query.ListRecipes(context.Background())
	if err != nil {
		log.Println("failed to load recipes:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	c := components.RecipeList(recipes)
	_ = c.Render(context.Background(), w)
}
