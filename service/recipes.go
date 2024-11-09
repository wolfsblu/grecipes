package service

import (
	"context"
	"sync"

	"github.com/wolfsblu/grecipes/api"
)

type RecipesService struct {
	Recipes map[int64]api.Recipe
	id      int64
	mux     sync.Mutex
}

func (p *RecipesService) AddRecipe(ctx context.Context, req *api.Recipe) (*api.Recipe, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	p.Recipes[p.id] = *req
	p.id++
	return req, nil
}

func (p *RecipesService) DeleteRecipe(ctx context.Context, params api.DeleteRecipeParams) error {
	p.mux.Lock()
	defer p.mux.Unlock()

	delete(p.Recipes, params.RecipeId)
	return nil
}

func (p *RecipesService) GetRecipeById(ctx context.Context, params api.GetRecipeByIdParams) (api.GetRecipeByIdRes, error) {
	p.mux.Lock()
	defer p.mux.Unlock()

	Recipe, ok := p.Recipes[params.RecipeId]
	if !ok {
		// Return Not Found.
		return &api.GetRecipeByIdNotFound{}, nil
	}
	return &Recipe, nil
}

func (p *RecipesService) UpdateRecipe(ctx context.Context, params api.UpdateRecipeParams) error {
	p.mux.Lock()
	defer p.mux.Unlock()

	Recipe := p.Recipes[params.RecipeId]
	Recipe.Status = params.Status
	if val, ok := params.Name.Get(); ok {
		Recipe.Name = val
	}
	p.Recipes[params.RecipeId] = Recipe

	return nil
}
