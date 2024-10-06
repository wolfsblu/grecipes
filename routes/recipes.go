package routes

import (
	"context"
	"github.com/wolfsblu/grecipes/db"
	"github.com/wolfsblu/grecipes/templates"
	"log"
	"net/http"
)

func GetRecipes(w http.ResponseWriter, _ *http.Request) {
	recipes, err := db.Query.ListRecipes(context.Background())
	if err != nil {
		log.Println("failed to load recipes:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	data := map[string]interface{}{"Recipes": recipes}
	_ = templates.Template["index.html"].ExecuteTemplate(w, "recipe-list", data)
}
