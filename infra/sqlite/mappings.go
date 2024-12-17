package sqlite

import (
	"github.com/wolfsblu/go-chef/domain"
)

func (t *PasswordReset) AsDomainModel() domain.PasswordResetToken {
	return domain.PasswordResetToken{
		Token:     t.Token,
		CreatedAt: t.CreatedAt,
	}
}
func (r *Recipe) AsDomainModel() domain.Recipe {
	return domain.Recipe{
		ID: r.ID,
		RecipeDetails: domain.RecipeDetails{
			Name: r.Name,
			CreatedBy: &domain.User{
				ID: r.CreatedBy,
			},
		},
	}
}

func (r *User) AsDomainModel() domain.User {
	return domain.User{
		ID: r.ID,
		Credentials: domain.Credentials{
			Email:        r.Email,
			PasswordHash: r.PasswordHash,
		},
	}
}
