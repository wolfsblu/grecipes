package routing

import (
	"net/http"
	"os"
)

func cors(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin, ok := os.LookupEnv("CORS_ORIGIN")
		if ok {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
		}
		h.ServeHTTP(w, r)
	}
}
