package db

import (
	"context"
	"database/sql"
	"github.com/wolfsblu/go-chef/domain"
	_ "modernc.org/sqlite"
)

func Connect(path string) (*Queries, error) {
	con, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	return New(con), nil
}

type SqliteStore struct {
	DB *Queries
}

func (s *SqliteStore) CreateRecipe(ctx context.Context, r domain.RecipeDetails) (recipe domain.Recipe, _ error) {
	payload := CreateRecipeParams{
		Name:      r.Name,
		CreatedBy: r.CreatedBy.ID,
	}

	result, err := s.DB.CreateRecipe(ctx, payload)
	if err != nil {
		return recipe, err
	}

	return result.AsDomainModel(), nil
}

func (s *SqliteStore) CreateUser(ctx context.Context, credentials domain.Credentials) (user domain.User, _ error) {
	result, err := s.DB.CreateUser(ctx, CreateUserParams{
		Email:        credentials.Email,
		PasswordHash: credentials.PasswordHash,
	})
	if err != nil {
		return user, err
	}
	return result.AsDomainModel(), nil
}

func (s *SqliteStore) DeleteRecipe(ctx context.Context, id int64) error {
	return s.DB.DeleteRecipe(ctx, id)
}

func (s *SqliteStore) GetRecipeById(ctx context.Context, id int64) (recipe domain.Recipe, _ error) {
	result, err := s.DB.GetRecipe(ctx, id)
	if err != nil {
		return recipe, err
	}
	return result.AsDomainModel(), nil
}

func (s *SqliteStore) GetRecipesByUser(ctx context.Context, user *domain.User) (recipes []domain.Recipe, _ error) {
	result, err := s.DB.ListRecipes(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	for _, item := range result {
		recipes = append(recipes, item.AsDomainModel())
	}
	return recipes, nil
}

func (s *SqliteStore) GetUserByEmail(ctx context.Context, email string) (user domain.User, _ error) {
	result, err := s.DB.GetUserByEmail(ctx, email)
	if err != nil {
		return user, err
	}
	return result.AsDomainModel(), nil
}

func (s *SqliteStore) GetUserById(ctx context.Context, id int64) (user domain.User, _ error) {
	result, err := s.DB.GetUser(ctx, id)
	if err != nil {
		return user, err
	}
	return result.AsDomainModel(), nil
}
