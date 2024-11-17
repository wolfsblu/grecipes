package main

import (
	"context"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gorilla/sessions"
	"github.com/swaggest/swgui/v5emb"
	"github.com/wolfsblu/grecipes/api"
	"github.com/wolfsblu/grecipes/db"
	"github.com/wolfsblu/grecipes/env"
	"github.com/wolfsblu/grecipes/middleware"
	"github.com/wolfsblu/grecipes/routes"
	"github.com/wolfsblu/grecipes/service"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func main() {
	env.Load()

	var sessionStore = sessions.NewCookieStore([]byte(env.MustGet("SESSION_KEY")))

	dbPath := env.MustGet("DB_PATH")
	query, err := db.Connect(dbPath)
	if err != nil {
		log.Fatalln("failed to connect to the database:", err)
	}
	svc := service.New(query, sessionStore)

	handler, err := api.NewServer(svc)
	if err != nil {
		log.Fatalln("failed to start api server:", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/api/docs/", v5emb.New("OpenAPI Docs", "/api/openapi.yml", "/api/docs/"))
	mux.Handle("/api/", middleware.Authenticate(sessionStore, http.StripPrefix("/api", handler)))

	provider, err := oidc.NewProvider(context.Background(), env.MustGet("OIDC_ISSUER_URL"))
	if err != nil {
		log.Fatalln("failed to create oidc provider:", err)
	}
	oauth2Config := oauth2.Config{
		ClientID:     env.MustGet("OIDC_CLIENT_ID"),
		ClientSecret: env.MustGet("OIDC_CLIENT_SECRET"),
		RedirectURL:  fmt.Sprintf("http://%s/oidc/login/callback/", env.MustGet("HOST")),
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
	verifier := provider.Verifier(&oidc.Config{ClientID: env.MustGet("OIDC_CLIENT_ID")})
	mux.HandleFunc("/oidc/login/callback/", func(w http.ResponseWriter, r *http.Request) {
		// Verify state and errors.

		ctx := context.Background()
		oauth2Token, err := oauth2Config.Exchange(ctx, r.URL.Query().Get("code"))
		if err != nil {
			log.Fatalln("failed to get access token:", err)
		}

		// Extract the ID Token from OAuth2 token.
		rawIDToken, ok := oauth2Token.Extra("id_token").(string)
		if !ok {
			log.Fatalln("failed to get id token:", err)
		}

		// Parse and verify ID Token payload.
		idToken, err := verifier.Verify(ctx, rawIDToken)
		if err != nil {
			log.Fatalln("failed to verify token:", err)
		}

		// Extract custom claims
		var claims struct {
			Email    string `json:"email"`
			Verified bool   `json:"email_verified"`
		}
		if err := idToken.Claims(&claims); err != nil {
			log.Fatalln("failed to verify token claims:", err)
		}
	})
	mux.HandleFunc("/oidc/login/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, oauth2Config.AuthCodeURL("test"), http.StatusFound)
	})
	routes.Register(mux)

	host := env.MustGet("HOST")
	err = http.ListenAndServe(host, mux)
	if err != nil {
		log.Fatalln("failed to start web server:", err)
	}
}
