package routes

import (
	"github.com/wolfsblu/grecipes/docs"
	"net/http"
)

func Docs(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, docs.APIDocsFS, "openapi.yml")
}
