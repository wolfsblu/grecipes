package handlers

import (
	"context"
	"fmt"
	"github.com/wolfsblu/go-chef/api"
	"github.com/wolfsblu/go-chef/domain"
	"github.com/wolfsblu/go-chef/domain/security"
)

func (h *RecipeHandler) GetUserProfile(ctx context.Context) (*api.ReadUser, error) {
	user := ctx.Value(ctxKeyUser).(*domain.User)
	return &api.ReadUser{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}

func (h *RecipeHandler) Login(ctx context.Context, req *api.Credentials) (r *api.AuthenticatedUserHeaders, _ error) {
	user, err := h.Recipes.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	err = h.Recipes.VerifyPassword(user, req.Password)
	if err != nil {
		return nil, err
	}
	cookie, err := createSessionCookie(user.ID)
	if err != nil {
		return nil, &domain.ErrSecurity
	}

	return &api.AuthenticatedUserHeaders{
		SetCookie: api.OptString{
			Set:   true,
			Value: cookie,
		},
		Response: api.ReadUser{
			ID:    user.ID,
			Email: user.Email,
		},
	}, nil
}

func (h *RecipeHandler) Logout(_ context.Context) (*api.LogoutOK, error) {
	cookie := expireSessionCookie()
	return &api.LogoutOK{
		SetCookie: api.OptString{
			Set:   true,
			Value: cookie,
		},
	}, nil
}

func (h *RecipeHandler) Register(ctx context.Context, c *api.Credentials) error {
	hash, err := security.CreateHash(c.Password, security.DefaultHashParams)
	if err != nil {
		return fmt.Errorf("%w: %w", &domain.ErrSecurity, err)
	}
	err = h.Recipes.RegisterUser(ctx, domain.Credentials{
		Email:        c.Email,
		PasswordHash: hash,
	})
	if err != nil {
		return fmt.Errorf("%w: %w", &domain.ErrRegistration, err)
	}
	return nil
}

func (h *RecipeHandler) ResetPassword(ctx context.Context, req *api.ResetPasswordReq) error {
	return h.Recipes.ResetPasswordByEmail(ctx, req.Email)
}

func (h *RecipeHandler) UpdatePassword(ctx context.Context, req *api.UpdatePassword) error {
	hashedPassword, err := security.CreateHash(req.Password, security.DefaultHashParams)
	if err != nil {
		return err
	}
	return h.Recipes.UpdatePasswordByToken(ctx, req.Token, hashedPassword)
}
