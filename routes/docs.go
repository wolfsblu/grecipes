package routes

import (
	"github.com/wolfsblu/go-chef/api"
	"net/http"
)

func apiDocs(w http.ResponseWriter, r *http.Request) {
	http.ServeFileFS(w, r, api.DocsFS, "openapi.yml")
}
