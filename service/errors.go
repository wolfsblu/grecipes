package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/wolfsblu/go-chef/api"
)

type RecipeServiceError struct {
	HttpStatusCode int
	ErrorCode      int
	ErrorMessage   string
}

func (e RecipeServiceError) Error() string {
	return fmt.Sprintf("status: %d, code %d: %+v", e.HttpStatusCode, e.ErrorCode, e.ErrorMessage)
}

var (
	ErrUnhandled      = RecipeServiceError{HttpStatusCode: 500, ErrorCode: 1, ErrorMessage: "internal server error"}
	ErrRecipeNotFound = RecipeServiceError{HttpStatusCode: 404, ErrorCode: 2, ErrorMessage: "recipe not found"}
	ErrSecurity       = RecipeServiceError{HttpStatusCode: 403, ErrorCode: 3, ErrorMessage: "authentication required"}
)

func (p *RecipesService) NewError(ctx context.Context, err error) (r *api.ErrorStatusCode) {
	var serviceError RecipeServiceError
	var securityError *ogenerrors.SecurityError

	switch {
	case errors.As(err, &securityError):
		serviceError = ErrSecurity
	default:
		serviceError = ErrUnhandled
	}

	return &api.ErrorStatusCode{
		StatusCode: serviceError.HttpStatusCode,
		Response: api.Error{
			Code:    serviceError.ErrorCode,
			Message: fmt.Sprintf("%s: %s", serviceError.ErrorMessage, err),
		},
	}
}
