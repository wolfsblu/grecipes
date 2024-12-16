package domain

import (
	"context"
)

type RecipeService struct {
	sender NotificationSender
	store  RecipeStore
}

type RecipeDetails struct {
	Name      string
	CreatedBy *User
}

type Recipe struct {
	ID int64
	RecipeDetails
}

func (s *RecipeService) Add(ctx context.Context, r RecipeDetails) (Recipe, error) {
	return s.store.CreateRecipe(ctx, r)
}

func (s *RecipeService) Delete(ctx context.Context, id int64) error {
	return s.store.DeleteRecipe(ctx, id)
}

func (s *RecipeService) GetByUser(ctx context.Context, user *User) ([]Recipe, error) {
	return s.store.GetRecipesByUser(ctx, user)
}

func (s *RecipeService) GetById(ctx context.Context, id int64) (Recipe, error) {
	return s.store.GetRecipeById(ctx, id)
}
