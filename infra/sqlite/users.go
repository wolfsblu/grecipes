package sqlite

import (
	"context"
	"database/sql"
	"github.com/wolfsblu/go-chef/domain"
	"github.com/wolfsblu/go-chef/domain/security"
)

func (s *Store) CreatePasswordResetToken(ctx context.Context, user *domain.User) (token domain.PasswordResetToken, _ error) {
	result, err := s.db.CreatePasswordResetToken(ctx, CreatePasswordResetTokenParams{
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
	result, err := s.db.CreateUser(ctx, CreateUserParams{
		Email:        credentials.Email,
		PasswordHash: credentials.PasswordHash,
	})
	if err != nil {
		return user, err
	}
	return result.AsDomainModel(), nil
}

func (s *Store) DeletePasswordResetTokenByUser(ctx context.Context, user *domain.User) error {
	return s.db.DeletePasswordResetTokenByUserId(ctx, user.ID)
}

func (s *Store) GetPasswordResetToken(ctx context.Context, searchToken string) (token domain.PasswordResetToken, _ error) {
	result, err := s.db.GetPasswordResetToken(ctx, searchToken)
	if err != nil {
		return token, err
	}
	token = result.PasswordReset.AsDomainModel()
	user := result.User.AsDomainModel()
	token.User = &user
	return token, nil
}

func (s *Store) GetPasswordResetTokenByUser(ctx context.Context, user *domain.User) (token domain.PasswordResetToken, _ error) {
	result, err := s.db.GetPasswordResetTokenByUser(ctx, user.ID)
	if err != nil {
		return token, err
	}
	token = result.AsDomainModel()
	token.User = user
	return token, nil
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (user domain.User, _ error) {
	result, err := s.db.GetUserByEmail(ctx, email)
	if err != nil {
		return user, err
	}
	return result.AsDomainModel(), nil
}

func (s *Store) GetUserById(ctx context.Context, id int64) (user domain.User, _ error) {
	result, err := s.db.GetUser(ctx, id)
	if err != nil {
		return user, err
	}
	return result.AsDomainModel(), nil
}

func (s *Store) UpdatePasswordByUser(ctx context.Context, user *domain.User, hashedPassword string) error {
	return s.db.UpdatePasswordByUserId(ctx, UpdatePasswordByUserIdParams{
		PasswordHash: hashedPassword,
		ID:           user.ID,
	})
}

func (s *Store) UpdatePasswordByToken(ctx context.Context, searchToken, hashedPassword string) error {
	tx, err := s.con.Begin()
	if err != nil {
		return err
	}

	defer func(tx *sql.Tx) {
		_ = tx.Rollback()
	}(tx)

	qtx := s.db.WithTx(tx)
	token, err := qtx.GetPasswordResetToken(ctx, searchToken)
	if err != nil {
		return err
	}
	if err = qtx.UpdatePasswordByUserId(ctx, UpdatePasswordByUserIdParams{
		PasswordHash: hashedPassword,
		ID:           token.User.ID,
	}); err != nil {
		return err
	}
	if err = qtx.DeletePasswordResetTokenByUserId(ctx, token.User.ID); err != nil {
		return err
	}
	return tx.Commit()
}
