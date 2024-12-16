package factories

import (
	"fmt"
	"github.com/wolfsblu/go-chef/domain"
	"github.com/wolfsblu/go-chef/infra/env"
	"github.com/wolfsblu/go-chef/infra/smtp"
	"github.com/wolfsblu/go-chef/infra/sqlite"
)

func newRecipeService() (*domain.RecipeService, error) {
	dbPath := env.MustGet("DB_PATH")
	query, err := sqlite.Connect(dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	err = sqlite.Migrate(fmt.Sprintf("sqlite://%s", dbPath))
	if err != nil {
		return nil, fmt.Errorf("failed to apply database migrations: %w", err)
	}

	return &domain.RecipeService{
		Store:  &sqlite.Store{DB: query},
		Sender: &smtp.Mailer{},
	}, nil
}
