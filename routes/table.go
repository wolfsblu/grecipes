package routes

import "net/http"

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/assets/", Assets)
	mux.HandleFunc("/", Index)
}
