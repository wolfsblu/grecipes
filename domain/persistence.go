package domain

import "context"

type RecipeStore interface {
	CreateRecipe(ctx context.Context, recipe RecipeDetails) (Recipe, error)
	CreateUser(ctx context.Context, credentials Credentials) (User, error)
	DeleteRecipe(ctx context.Context, id int64) error
	GetRecipeById(ctx context.Context, id int64) (Recipe, error)
	GetRecipesByUser(ctx context.Context, user *User) ([]Recipe, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserById(ctx context.Context, id int64) (User, error)
}
