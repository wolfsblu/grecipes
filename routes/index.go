package routes

import (
	"context"
	"github.com/wolfsblu/grecipes/db"
	"github.com/wolfsblu/grecipes/templates"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	getRecipes()
	err := templates.Template["index.html"].ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		log.Println("failed to execute template:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func getRecipes() {
	recipes, err := db.Query.ListRecipes(context.Background())
	if err != nil {
		log.Println(err)
	}
	log.Println(recipes)
}
