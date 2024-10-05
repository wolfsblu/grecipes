package routes

import (
	"github.com/wolfsblu/grecipes/templates"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	err := templates.Template["index.html"].ExecuteTemplate(w, "base.html", nil)
	if err != nil {
		log.Println("failed to execute template:", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
