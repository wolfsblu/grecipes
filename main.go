package main

import (
	"github.com/wolfsblu/grecipes/api"
	"github.com/wolfsblu/grecipes/routes"
	"github.com/wolfsblu/grecipes/service"
	"log"
	"net/http"
)

func main() {
	petsService := &service.PetsService{
		Pets: map[int64]api.Pet{},
	}
	petsServer, err := api.NewServer(petsService)
	if err != nil {
		log.Fatalln("failed to start API server:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/", http.StripPrefix("/api", petsServer))
	routes.Register(mux)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln("failed to start server:", err)
	}
}
