package service

import (
	"context"
	"fmt"
	"github.com/wolfsblu/go-chef/api"
	"github.com/wolfsblu/go-chef/db"
	"github.com/wolfsblu/go-chef/security"
)

type ContextKey string

const (
	CtxKeyUser = ContextKey("User")
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
	user, err := s.Db.GetUser(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", &ErrSecurity, err)
	}
	return context.WithValue(ctx, CtxKeyUser, user), nil
}
