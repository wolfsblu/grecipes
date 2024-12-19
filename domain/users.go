package domain

import (
	"context"
	"errors"
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
	User      *User
	Token     string
	CreatedAt time.Time
}

type UserRegistration struct {
	User      *User
	Token     string
	CreatedAt time.Time
}

func (s *RecipeService) UpdatePasswordByToken(ctx context.Context, searchToken, hashedPassword string) error {
	return s.store.UpdatePasswordByToken(ctx, searchToken, hashedPassword)
}

func (s *RecipeService) GetUserById(ctx context.Context, id int64) (User, error) {
	return s.store.GetUserById(ctx, id)
}

func (s *RecipeService) GetUserByEmail(ctx context.Context, email string) (User, error) {
	return s.store.GetUserByEmail(ctx, email)
}

func (s *RecipeService) RegisterUser(ctx context.Context, credentials Credentials) error {
	if err := s.store.Begin(ctx); err != nil {
		return err
	}
	defer s.store.Rollback()
	user, err := s.store.GetUserByEmail(ctx, credentials.Email)
	if err == nil {
		return nil
	} else if !errors.Is(err, &ErrUserNotFound) {
		return err
	}
	user, err = s.store.CreateUser(ctx, credentials)
	if err != nil {
		return err
	}
	registration, err := s.store.CreateUserRegistration(ctx, &user)
	if err != nil {
		return err
	}
	if err = s.store.Commit(); err != nil {
		return err
	}
	go func() {
		_ = s.sender.SendUserRegistration(registration)
	}()
	return nil
}

func (s *RecipeService) ResetPasswordByEmail(ctx context.Context, email string) error {
	user, err := s.store.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	token, err := s.store.GetPasswordResetTokenByUser(ctx, &user)
	if err == nil { // When there already is a reset token we want to do nothing and exit early
		return nil
	}
	token, err = s.store.CreatePasswordResetToken(ctx, &user)
	if err != nil {
		return err
	}
	go func() {
		_ = s.sender.SendPasswordReset(token)
	}()
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
