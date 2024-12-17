package sqlite

import (
	"context"
	"database/sql"
	"github.com/wolfsblu/go-chef/domain"
	_ "modernc.org/sqlite"
)

type Store struct {
	con  *sql.DB
	db   *Queries
	path string
}

func (s *Store) CreateRecipe(ctx context.Context, r domain.RecipeDetails) (recipe domain.Recipe, _ error) {
	payload := CreateRecipeParams{
		Name:      r.Name,
		CreatedBy: r.CreatedBy.ID,
	}

	result, err := s.db.CreateRecipe(ctx, payload)
	if err != nil {
		return recipe, err
	}

	return result.AsDomainModel(), nil
}

func (s *Store) DeleteRecipe(ctx context.Context, id int64) error {
	return s.db.DeleteRecipe(ctx, id)
}

func (s *Store) GetRecipeById(ctx context.Context, id int64) (recipe domain.Recipe, _ error) {
	result, err := s.db.GetRecipe(ctx, id)
	if err != nil {
		return recipe, err
	}
	return result.AsDomainModel(), nil
}

func (s *Store) GetRecipesByUser(ctx context.Context, user *domain.User) (recipes []domain.Recipe, _ error) {
	result, err := s.db.ListRecipes(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	for _, item := range result {
		recipes = append(recipes, item.AsDomainModel())
	}
	return recipes, nil
}
