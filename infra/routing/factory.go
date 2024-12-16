package routing

import (
	swagger "github.com/swaggest/swgui/v5emb"
	"github.com/wolfsblu/go-chef/api"
	"net/http"
)

func NewServeMux(server *api.Server) *http.ServeMux {
	mux := http.NewServeMux()
	handleFrontend(mux)
	handleAPI(mux, server)
	return mux
}

func handleFrontend(mux *http.ServeMux) {
	mux.HandleFunc("/assets/", assets)
	mux.HandleFunc("/", index)
}

func handleAPI(mux *http.ServeMux, apiServer http.Handler) {
	mux.Handle("/api/docs/", swagger.New("OpenAPI Docs", "/api/openapi.yml", "/api/docs/"))
	mux.Handle("/api/", cors(http.StripPrefix("/api", apiServer)))
	mux.HandleFunc("/api/openapi.yml", apiDocs)
}
