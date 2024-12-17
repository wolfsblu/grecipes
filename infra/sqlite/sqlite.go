package sqlite

import (
	"ariga.io/atlas-go-sdk/atlasexec"
	"context"
	"fmt"
	"github.com/wolfsblu/go-chef/domain"
	"github.com/wolfsblu/go-chef/domain/security"
	"io/fs"
	_ "modernc.org/sqlite"
)

type Store struct {
	db   *Queries
	path string
}

func (s *Store) Migrate() error {
	subFS, err := fs.Sub(migrationFS, "migrations")
	if err != nil {
		return err
	}

	workdir, err := atlasexec.NewWorkingDir(
		atlasexec.WithMigrations(subFS),
	)
	if err != nil {
		return err
	}

	defer func(workdir *atlasexec.WorkingDir) {
		err = workdir.Close()
	}(workdir)

	client, err := atlasexec.NewClient(workdir.Path(), "atlas")
	if err != nil {
		return err
	}

	_, err = client.MigrateApply(context.Background(), &atlasexec.MigrateApplyParams{
		URL: fmt.Sprintf("sqlite://%s", s.path),
	})
	return err
}

func (s *Store) CreateRecipe(ctx context.Context, r domain.RecipeDetails) (recipe domain.Recipe, _ error) {
	payload := CreateRecipeParams{
		Name:      r.Name,
		CreatedBy: r.CreatedBy.ID,
	}

	result, err := s.db.CreateRecipe(ctx, payload)
	if err != nil {
		return recipe, err
	}

	return result.AsDomainModel(), nil
}

func (s *Store) CreatePasswordResetToken(ctx context.Context, user domain.User) (token domain.PasswordResetToken, _ error) {
	result, err := s.db.CreatePasswordResetToken(ctx, CreatePasswordResetTokenParams{
		UserID: user.ID,
		Token:  security.GenerateToken(security.DefaultTokenLength),
	})
	if err != nil {
		return token, err
	}

	token = result.AsDomainModel()
	token.User = user
	return token, nil
}

func (s *Store) CreateUser(ctx context.Context, credentials domain.Credentials) (user domain.User, _ error) {
	result, err := s.db.CreateUser(ctx, CreateUserParams{
		Email:        credentials.Email,
		PasswordHash: credentials.PasswordHash,
	})
	if err != nil {
		return user, err
	}
	return result.AsDomainModel(), nil
}

func (s *Store) DeleteRecipe(ctx context.Context, id int64) error {
	return s.db.DeleteRecipe(ctx, id)
}

func (s *Store) GetRecipeById(ctx context.Context, id int64) (recipe domain.Recipe, _ error) {
	result, err := s.db.GetRecipe(ctx, id)
	if err != nil {
		return recipe, err
	}
	return result.AsDomainModel(), nil
}

func (s *Store) GetRecipesByUser(ctx context.Context, user *domain.User) (recipes []domain.Recipe, _ error) {
	result, err := s.db.ListRecipes(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	for _, item := range result {
		recipes = append(recipes, item.AsDomainModel())
	}
	return recipes, nil
}

func (s *Store) GetUserByEmail(ctx context.Context, email string) (user domain.User, _ error) {
	result, err := s.db.GetUserByEmail(ctx, email)
	if err != nil {
		return user, err
	}
	return result.AsDomainModel(), nil
}

func (s *Store) GetUserById(ctx context.Context, id int64) (user domain.User, _ error) {
	result, err := s.db.GetUser(ctx, id)
	if err != nil {
		return user, err
	}
	return result.AsDomainModel(), nil
}
