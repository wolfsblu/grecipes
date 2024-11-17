package middleware

import (
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

func Authenticate(store *sessions.CookieStore, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware called")
		next.ServeHTTP(w, r)
	})
}
