package service

import (
	"context"
	"github.com/wolfsblu/grecipes/api"
	"github.com/wolfsblu/grecipes/db"
)

type RecipesService struct{}

func (p *RecipesService) AddRecipe(ctx context.Context, req *api.WriteRecipe) (*api.ReadRecipe, error) {
	recipe, err := db.Query.CreateRecipe(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &api.ReadRecipe{
		ID:   recipe.ID,
		Name: recipe.Name,
	}, nil
}

func (p *RecipesService) DeleteRecipe(ctx context.Context, params api.DeleteRecipeParams) error {
	err := db.Query.DeleteRecipe(ctx, params.RecipeId)
	if err != nil {
		return err
	}
	return nil
}

func (p *RecipesService) GetRecipes(ctx context.Context) ([]api.ReadRecipe, error) {
	recipes, err := db.Query.ListRecipes(ctx)
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
	recipe, err := db.Query.GetRecipe(ctx, params.RecipeId)
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
