package routes

import (
	"github.com/wolfsblu/grecipes/templates"
	"net/http"
)

func Assets(w http.ResponseWriter, r *http.Request) {
	h := http.FileServer(http.FS(templates.PublicFS))
	h.ServeHTTP(w, r)
}
