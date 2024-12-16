package factories

import (
	"github.com/wolfsblu/go-chef/api"
	"github.com/wolfsblu/go-chef/infra/handlers"
)

func NewApiServer() (*api.Server, error) {
	recipeService, err := newRecipeService()
	if err != nil {
		return nil, err
	}

	rh := handlers.NewRecipeHandler(recipeService)
	sh := handlers.NewSecurityHandler(recipeService)
	return api.NewServer(rh, sh)
}
