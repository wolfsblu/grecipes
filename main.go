package main

import (
	"github.com/wolfsblu/grecipes/routes"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.Index)
	err := http.ListenAndServe("127.0.0.1:8080", mux)
	if err != nil {
		log.Fatalln("Failed to start server:", err)
	}
}
