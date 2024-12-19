package domain

import "context"

type RecipeStore interface {
	Begin(ctx context.Context) error
	Commit() error
	CreateRecipe(ctx context.Context, recipe RecipeDetails) (Recipe, error)
	CreatePasswordResetToken(ctx context.Context, user *User) (PasswordResetToken, error)
	CreateUser(ctx context.Context, credentials Credentials) (User, error)
	CreateUserRegistration(ctx context.Context, user *User) (UserRegistration, error)
	DeleteRecipe(ctx context.Context, id int64) error
	GetPasswordResetTokenByUser(ctx context.Context, user *User) (PasswordResetToken, error)
	GetRecipeById(ctx context.Context, id int64) (Recipe, error)
	GetRecipesByUser(ctx context.Context, user *User) ([]Recipe, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserById(ctx context.Context, id int64) (User, error)
	Rollback()
	UpdatePasswordByToken(ctx context.Context, token, hashedPassword string) error
}
