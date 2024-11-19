package service

import (
	"context"
	ht "github.com/ogen-go/ogen/http"
	"github.com/wolfsblu/go-chef/api"
	"github.com/wolfsblu/go-chef/security"
)

func (p *RecipesService) Login(ctx context.Context, req *api.Credentials) (r *api.AuthenticatedUserHeaders, _ error) {
	var userId int64 = 10
	cookie, err := security.NewSessionCookie(userId)
	if err != nil {
		return nil, ErrSecurity
	}
	return &api.AuthenticatedUserHeaders{
		SetCookie: api.OptString{
			Set:   true,
			Value: cookie,
		},
		Response: api.ReadUser{
			ID: userId,
		},
	}, nil
}

func (p *RecipesService) Register(ctx context.Context) (r *api.ReadUser, _ error) {
	return r, ht.ErrNotImplemented
}
