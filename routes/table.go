package routes

import (
	"github.com/gorilla/sessions"
	swagger "github.com/swaggest/swgui/v5emb"
	"net/http"
)

func RegisterApi(mux *http.ServeMux, apiServer http.Handler) {
	mux.Handle("/api/docs/", swagger.New("OpenAPI Docs", "/api/openapi.yml", "/api/docs/"))
	mux.Handle("/api/", http.StripPrefix("/api", apiServer))
	mux.HandleFunc("/api/openapi.yml", apiDocs)
}

func RegisterApp(mux *http.ServeMux) {
	mux.HandleFunc("/assets/", assets)
	mux.HandleFunc("/", app)
}

func RegisterAuth(mux *http.ServeMux, store *sessions.CookieStore) {
	mux.HandleFunc("/auth/login", func(w http.ResponseWriter, r *http.Request) {

	})
}
