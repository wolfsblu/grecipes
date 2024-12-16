package handlers

import (
	"context"
	"github.com/wolfsblu/go-chef/api"
	"github.com/wolfsblu/go-chef/domain"
)

type RecipeHandler struct {
	Recipes *domain.RecipeService
}

func NewRecipeHandler(service *domain.RecipeService) *RecipeHandler {
	return &RecipeHandler{
		Recipes: service,
	}
}

func (h *RecipeHandler) AddRecipe(ctx context.Context, req *api.WriteRecipe) (*api.ReadRecipe, error) {
	user := ctx.Value(ctxKeyUser).(*domain.User)
	recipe, err := h.Recipes.Add(ctx, domain.RecipeDetails{
		Name:      req.Name,
		CreatedBy: user,
	})
	if err != nil {
		return nil, err
	}
	return &api.ReadRecipe{
		ID:   recipe.ID,
		Name: recipe.Name,
	}, nil
}

func (h *RecipeHandler) DeleteRecipe(ctx context.Context, params api.DeleteRecipeParams) error {
	err := h.Recipes.Delete(ctx, params.RecipeId)
	if err != nil {
		return err
	}
	return nil
}

func (h *RecipeHandler) GetRecipes(ctx context.Context) ([]api.ReadRecipe, error) {
	user := ctx.Value(ctxKeyUser).(*domain.User)
	recipes, err := h.Recipes.GetByUser(ctx, user)
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

func (h *RecipeHandler) GetRecipeById(ctx context.Context, params api.GetRecipeByIdParams) (*api.ReadRecipe, error) {
	recipe, err := h.Recipes.GetById(ctx, params.RecipeId)
	if err != nil {
		return nil, &domain.ErrRecipeNotFound
	}
	return &api.ReadRecipe{
		ID:   recipe.ID,
		Name: recipe.Name,
	}, nil
}

func (h *RecipeHandler) UpdateRecipe(_ context.Context, _ *api.WriteRecipe, _ api.UpdateRecipeParams) (*api.ReadRecipe, error) {
	// TODO: Implement
	return &api.ReadRecipe{}, nil
}
