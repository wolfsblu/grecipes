package service

import (
	"context"
	"github.com/wolfsblu/go-chef/api"
	"github.com/wolfsblu/go-chef/db"
)

type RecipesService struct {
	Db *db.Queries
}

func New(query *db.Queries) *RecipesService {
	return &RecipesService{
		Db: query,
	}
}

func (p *RecipesService) AddRecipe(ctx context.Context, req *api.WriteRecipe) (*api.ReadRecipe, error) {
	recipe, err := p.Db.CreateRecipe(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &api.ReadRecipe{
		ID:   recipe.ID,
		Name: recipe.Name,
	}, nil
}

func (p *RecipesService) DeleteRecipe(ctx context.Context, params api.DeleteRecipeParams) error {
	err := p.Db.DeleteRecipe(ctx, params.RecipeId)
	if err != nil {
		return err
	}
	return nil
}

func (p *RecipesService) GetRecipes(ctx context.Context) ([]api.ReadRecipe, error) {
	recipes, err := p.Db.ListRecipes(ctx)
	if err != nil {
		return nil, err
	}
	var response []api.ReadRecipe
	for _, recipe := range recipes {
		response = append(response, api.ReadRecipe{
			ID:   recipe.ID,
			Name: recipe.Name,
		})
	}
	return response, nil
}

func (p *RecipesService) GetRecipeById(ctx context.Context, params api.GetRecipeByIdParams) (*api.ReadRecipe, error) {
	recipe, err := p.Db.GetRecipe(ctx, params.RecipeId)
	if err != nil {
		return nil, ErrRecipeNotFound
	}
	return &api.ReadRecipe{
		ID:   recipe.ID,
		Name: recipe.Name,
	}, nil
}

func (p *RecipesService) UpdateRecipe(ctx context.Context, req *api.WriteRecipe, params api.UpdateRecipeParams) (*api.ReadRecipe, error) {
	// TODO: Implement
	return &api.ReadRecipe{}, nil
}
