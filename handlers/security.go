package handlers

import (
	"context"
	"fmt"
	"github.com/wolfsblu/go-chef/api"
	"github.com/wolfsblu/go-chef/domain"
	"github.com/wolfsblu/go-chef/security"
)

type contextKey string

const (
	ctxKeyUser = contextKey("User")
)

type SecurityHandler struct {
	Recipes *domain.RecipeService
}

func NewSecurityHandler(service *domain.RecipeService) *SecurityHandler {
	return &SecurityHandler{
		Recipes: service,
	}
}

func (h *SecurityHandler) HandleCookieAuth(ctx context.Context, _ string, t api.CookieAuth) (context.Context, error) {
	userId, err := security.GetUserFromSessionCookie(t.APIKey)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", &domain.ErrSecurity, err)
	}
	user, err := h.Recipes.GetUserById(ctx, userId)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", &domain.ErrSecurity, err)
	}
	return context.WithValue(ctx, ctxKeyUser, &user), nil
}
