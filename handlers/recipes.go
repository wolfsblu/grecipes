package handlers

import (
	"context"
	"github.com/wolfsblu/go-chef/api"
	"github.com/wolfsblu/go-chef/db"
)

type RecipeHandler struct {
	DB *db.Queries
}

func NewRecipeHandler(query *db.Queries) *RecipeHandler {
	return &RecipeHandler{
		DB: query,
	}
}

func (h *RecipeHandler) AddRecipe(ctx context.Context, req *api.WriteRecipe) (*api.ReadRecipe, error) {
	user := ctx.Value(CtxKeyUser).(*db.User)
	payload := db.CreateRecipeParams{
		Name:      req.Name,
		CreatedBy: user.ID,
	}
	recipe, err := h.DB.CreateRecipe(ctx, payload)
	if err != nil {
		return nil, err
	}
	return &api.ReadRecipe{
		ID:   recipe.ID,
		Name: recipe.Name,
	}, nil
}

func (h *RecipeHandler) DeleteRecipe(ctx context.Context, params api.DeleteRecipeParams) error {
	err := h.DB.DeleteRecipe(ctx, params.RecipeId)
	if err != nil {
		return err
	}
	return nil
}

func (h *RecipeHandler) GetRecipes(ctx context.Context) ([]api.ReadRecipe, error) {
	user := ctx.Value(CtxKeyUser).(*db.User)
	recipes, err := h.DB.ListRecipes(ctx, user.ID)
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
	recipe, err := h.DB.GetRecipe(ctx, params.RecipeId)
	if err != nil {
		return nil, &ErrRecipeNotFound
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
