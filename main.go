package main

import (
	"github.com/gorilla/sessions"
	"github.com/wolfsblu/grecipes/api"
	"github.com/wolfsblu/grecipes/db"
	"github.com/wolfsblu/grecipes/env"
	"github.com/wolfsblu/grecipes/routes"
	"github.com/wolfsblu/grecipes/service"
	"log"
	"net/http"
)

func main() {
	env.Load()

	var sessionStore = sessions.NewCookieStore([]byte(env.MustGet("SESSION_KEY")))

	dbPath := env.MustGet("DB_PATH")
	query, err := db.Connect(dbPath)
	if err != nil {
		log.Fatalln("failed to connect to the database", err)
	}
	svc := service.New(query, sessionStore)
	sec := service.NewSecurity(query, sessionStore)

	apiServer, err := api.NewServer(svc, sec)
	if err != nil {
		log.Fatalln("failed to start api server:", err)
	}

	mux := http.NewServeMux()
	routes.RegisterApp(mux)
	routes.RegisterAuth(mux, sessionStore)
	routes.RegisterApi(mux, apiHandler)

	host := env.MustGet("HOST")
	err = http.ListenAndServe(host, mux)
	if err != nil {
		log.Fatalln("failed to start web server:", err)
	}
}
