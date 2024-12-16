package domain

import (
	"context"
	"fmt"
	"github.com/wolfsblu/go-chef/security"
)

type Credentials struct {
	Email        string
	PasswordHash string
}

type User struct {
	ID int64
	Credentials
}

func (s *RecipeService) GenerateSessionCookie(user User) (string, error) {
	cookie, err := security.NewSessionCookie(user.ID)
	if err != nil {
		return "", &ErrSecurity
	}
	return cookie, nil
}

func (s *RecipeService) GetUserById(ctx context.Context, id int64) (User, error) {
	return s.Store.GetUserById(ctx, id)
}

func (s *RecipeService) GetUserByEmail(ctx context.Context, email string) (User, error) {
	return s.Store.GetUserByEmail(ctx, email)
}

func (s *RecipeService) RegisterUser(ctx context.Context, credentials Credentials) (User, error) {
	return s.Store.CreateUser(ctx, credentials)
}

func (s *RecipeService) ResetPasswordByEmail(ctx context.Context, email string) error {
	// TODO: Implement
	_, err := s.Store.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	return nil
}

func (s *RecipeService) VerifyPassword(user User, password string) error {
	ok, err := security.ComparePasswordAndHash(password, user.PasswordHash)
	if err != nil {
		return fmt.Errorf("%w: %w", &ErrSecurity, err)
	} else if !ok {
		return fmt.Errorf("hallo welt: %w", &ErrInvalidCredentials)
	}
	return nil
}
