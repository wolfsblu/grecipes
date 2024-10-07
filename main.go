package main

import (
	"github.com/wolfsblu/grecipes/routes"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	routes.Register(mux)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln("failed to start server:", err)
	}
}
