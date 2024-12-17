package domain

import (
	"context"
	"fmt"
	"github.com/wolfsblu/go-chef/domain/security"
	"time"
)

type Credentials struct {
	Email        string
	PasswordHash string
}

type User struct {
	ID int64
	Credentials
}

type PasswordResetToken struct {
	User      User
	Token     string
	CreatedAt time.Time
}

func (s *RecipeService) GetUserById(ctx context.Context, id int64) (User, error) {
	return s.store.GetUserById(ctx, id)
}

func (s *RecipeService) GetUserByEmail(ctx context.Context, email string) (User, error) {
	return s.store.GetUserByEmail(ctx, email)
}

func (s *RecipeService) RegisterUser(ctx context.Context, credentials Credentials) (User, error) {
	return s.store.CreateUser(ctx, credentials)
}

func (s *RecipeService) ResetPasswordByEmail(ctx context.Context, email string) error {
	user, err := s.store.GetUserByEmail(ctx, email)
	// TODO: Extend existing token before creating a new one (there's a unique key on user id)
	if err != nil {
		return err
	}
	token, err := s.store.CreatePasswordResetToken(ctx, user)
	if err != nil {
		return err
	}
	err = s.sender.SendPasswordReset(token)
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
