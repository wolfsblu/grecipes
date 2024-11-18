package routes

import (
	"github.com/wolfsblu/go-chef/docs"
	"net/http"
)

func apiDocs(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, docs.APIDocsFS, "openapi.yml")
}
