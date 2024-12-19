package sqlite

import (
	"context"
	"database/sql"
	"github.com/go-faster/errors"
	"github.com/wolfsblu/go-chef/domain"
	"github.com/wolfsblu/go-chef/domain/security"
	"log"
)

func (s *Store) CreatePasswordResetToken(ctx context.Context, user *domain.User) (token domain.PasswordResetToken, _ error) {
	result, err := s.query().CreatePasswordResetToken(ctx, CreatePasswordResetTokenParams{
		UserID: user.ID,
		Token:  security.GenerateToken(security.DefaultTokenLength),
	})
	if err != nil {
		log.Println("failed to create password reset token:", err)
		return token, &domain.ErrPersistence
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
		log.Println("failed to create user:", err)
		return user, &domain.ErrPersistence
	}
	return result.AsDomainModel(), nil
}

func (s *Store) CreateUserRegistration(ctx context.Context, user *domain.User) (registration domain.UserRegistration, _ error) {
	result, err := s.query().CreateUserRegistration(ctx, CreateUserRegistrationParams{
		UserID: user.ID,
		Token:  security.GenerateToken(security.DefaultTokenLength),
	})
	if err != nil {
		log.Println("failed to create user registration:", err)
		return registration, &domain.ErrPersistence
	}

	registration = result.AsDomainModel()
	registration.User = user
	return registration, nil
}

func (s *Store) GetPasswordResetTokenByUser(ctx context.Context, user *domain.User) (token domain.PasswordResetToken, _ error) {
	result, err := s.query().GetPasswordResetTokenByUser(ctx, user.ID)
	if err != nil {
		log.Println("failed to retrieve password reset token:", err)
		return token, &domain.ErrRetrieval
	}
	token = result.AsDomainModel()
	token.User = user
	return token, nil
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (user domain.User, _ error) {
	result, err := s.query().GetUserByEmail(ctx, email)
	if errors.Is(err, sql.ErrNoRows) {
		return user, &domain.ErrUserNotFound
	} else if err != nil {
		log.Println("failed to retrieve user by email:", err)
		return user, &domain.ErrRetrieval
	}
	return result.AsDomainModel(), nil
}

func (s *Store) GetUserById(ctx context.Context, id int64) (user domain.User, _ error) {
	result, err := s.query().GetUser(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return user, &domain.ErrUserNotFound
	} else if err != nil {
		log.Println("failed to retrieve user by id:", err)
		return user, &domain.ErrRetrieval
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
		log.Println("failed to retrieve password reset token:", err)
		return &domain.ErrRetrieval
	}
	if err = s.query().UpdatePasswordByUserId(ctx, UpdatePasswordByUserIdParams{
		PasswordHash: hashedPassword,
		ID:           token.User.ID,
	}); err != nil {
		log.Println("failed to update password:", err)
		return &domain.ErrPersistence
	}
	if err = s.query().DeletePasswordResetTokenByUserId(ctx, token.User.ID); err != nil {
		log.Println("failed to delete password reset token:", err)
		return &domain.ErrPersistence
	}
	return s.Commit()
}
