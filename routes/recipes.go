package routes

import (
	"context"
	"encoding/json"
	"github.com/wolfsblu/grecipes/db"
	"log"
	"net/http"
)

func GetRecipes(w http.ResponseWriter, _ *http.Request) {
	recipes, err := db.Query.ListRecipes(context.Background())
	if err != nil {
		log.Println("failed to load recipes:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(recipes)
	if err != nil {
		log.Println("failed to deserialize recipes:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
