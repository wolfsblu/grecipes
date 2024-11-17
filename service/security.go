package service

import (
	"context"
	"github.com/gorilla/sessions"
	"github.com/wolfsblu/grecipes/api"
	"github.com/wolfsblu/grecipes/db"
	"log"
)

type SecurityService struct {
	Db      *db.Queries
	Session *sessions.CookieStore
}

func NewSecurity(query *db.Queries, session *sessions.CookieStore) *SecurityService {
	return &SecurityService{
		Db:      query,
		Session: session,
	}
}

func (s *SecurityService) HandleCookieAuth(ctx context.Context, operationName string, t api.CookieAuth) (context.Context, error) {
	log.Println(operationName)
	log.Println(t.GetAPIKey())
	return ctx, nil
}
