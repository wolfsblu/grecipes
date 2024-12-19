package sqlite

import (
	"context"
	"fmt"
	"github.com/wolfsblu/go-chef/domain"
	"github.com/wolfsblu/go-chef/domain/security"
)

func (s *Store) CreatePasswordResetToken(ctx context.Context, user *domain.User) (token domain.PasswordResetToken, _ error) {
	result, err := s.query().CreatePasswordResetToken(ctx, CreatePasswordResetTokenParams{
		UserID: user.ID,
		Token:  security.GenerateToken(security.DefaultTokenLength),
	})
	if err != nil {
		return token, err
	}

	token = result.AsDomainModel()
	token.User = user
	return token, nil
}

func (s *Store) CreateUser(ctx context.Context, credentials domain.Credentials) (user domain.User, _ error) {
	result, err := s.query().CreateUser(ctx, CreateUserParams{
		Email:        credentials.Email,
		PasswordHash: credentials.PasswordHash,
	})
	if err != nil {
		return user, err
	}
	return result.AsDomainModel(), nil
}

func (s *Store) CreateUserRegistration(ctx context.Context, user *domain.User) (registration domain.UserRegistration, _ error) {
	result, err := s.query().CreateUserRegistration(ctx, CreateUserRegistrationParams{
		UserID: user.ID,
		Token:  security.GenerateToken(security.DefaultTokenLength),
	})
	if err != nil {
		return registration, err
	}

	registration = result.AsDomainModel()
	registration.User = user
	return registration, nil
}

func (s *Store) GetPasswordResetTokenByUser(ctx context.Context, user *domain.User) (token domain.PasswordResetToken, _ error) {
	result, err := s.query().GetPasswordResetTokenByUser(ctx, user.ID)
	if err != nil {
		return token, err
	}
	token = result.AsDomainModel()
	token.User = user
	return token, nil
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (user domain.User, _ error) {
	result, err := s.query().GetUserByEmail(ctx, email)
	if err != nil {
		return user, fmt.Errorf("%w: %w", &domain.ErrUserNotFound, err)
	}
	return result.AsDomainModel(), nil
}

func (s *Store) GetUserById(ctx context.Context, id int64) (user domain.User, _ error) {
	result, err := s.query().GetUser(ctx, id)
	if err != nil {
		return user, err
	}
	return result.AsDomainModel(), nil
}

func (s *Store) UpdatePasswordByToken(ctx context.Context, searchToken, hashedPassword string) error {
	err := s.Begin(ctx)
	if err != nil {
		return err
	}
	defer s.Rollback()

	token, err := s.query().GetPasswordResetToken(ctx, searchToken)
	if err != nil {
		return err
	}
	if err = s.query().UpdatePasswordByUserId(ctx, UpdatePasswordByUserIdParams{
		PasswordHash: hashedPassword,
		ID:           token.User.ID,
	}); err != nil {
		return err
	}
	if err = s.query().DeletePasswordResetTokenByUserId(ctx, token.User.ID); err != nil {
		return err
	}
	return s.Commit()
}
