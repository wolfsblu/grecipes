package service

import (
	"context"
	"fmt"
	"github.com/wolfsblu/go-chef/api"
	"github.com/wolfsblu/go-chef/db"
	"github.com/wolfsblu/go-chef/security"
	"log"
)

type SecurityService struct {
	Db *db.Queries
}

func NewSecurity(query *db.Queries) *SecurityService {
	return &SecurityService{
		Db: query,
	}
}

func (s *SecurityService) HandleCookieAuth(ctx context.Context, operationName string, t api.CookieAuth) (context.Context, error) {
	userId, err := security.GetUserFromSessionCookie(t.APIKey)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", &ErrSecurity, err)
	}
	log.Println("Logged in as user with ID", userId)
	return ctx, nil
}
