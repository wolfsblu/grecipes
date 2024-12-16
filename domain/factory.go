package domain

func NewRecipeService(notifier NotificationSender, store RecipeStore) *RecipeService {
	return &RecipeService{
		store:  store,
		sender: notifier,
	}
}
