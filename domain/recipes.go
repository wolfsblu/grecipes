package domain

import (
	"context"
)

type RecipeService struct {
	Sender NotificationSender
	Store  RecipeStore
}

type RecipeDetails struct {
	Name      string
	CreatedBy *User
}

type Recipe struct {
	ID int64
	RecipeDetails
}

func NewRecipeService(store RecipeStore, sender NotificationSender) *RecipeService {
	return &RecipeService{Store: store, Sender: sender}
}

func (s *RecipeService) Add(ctx context.Context, r RecipeDetails) (Recipe, error) {
	return s.Store.CreateRecipe(ctx, r)
}

func (s *RecipeService) Delete(ctx context.Context, id int64) error {
	return s.Store.DeleteRecipe(ctx, id)
}

func (s *RecipeService) GetByUser(ctx context.Context, user *User) ([]Recipe, error) {
	return s.Store.GetRecipesByUser(ctx, user)
}

func (s *RecipeService) GetById(ctx context.Context, id int64) (Recipe, error) {
	return s.Store.GetRecipeById(ctx, id)
}
