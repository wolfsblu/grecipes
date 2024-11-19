package service

import (
	"context"
	"github.com/go-faster/errors"
	"github.com/wolfsblu/go-chef/api"
	"github.com/wolfsblu/go-chef/db"
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

func (p *RecipesService) Register(ctx context.Context, c *api.Credentials) (*api.ReadUser, error) {
	hash, err := security.CreateHash(c.Password, security.DefaultHashParams)
	if err != nil {
		return nil, errors.Wrap(err, ErrSecurity.Error())
	}
	creds := db.CreateUserParams{
		Email:        c.Email,
		PasswordHash: hash,
	}
	user, err := p.Db.CreateUser(ctx, creds)
	if err != nil {
		return nil, errors.Wrap(err, ErrSecurity.Error())
	}
	return &api.ReadUser{
		ID: user.ID,
	}, nil
}
