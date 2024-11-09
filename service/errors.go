package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/wolfsblu/grecipes/api"
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
)

func (p *RecipesService) NewError(ctx context.Context, err error) (r *api.ErrorStatusCode) {
	var serr RecipeServiceError
	ok := errors.As(err, &serr)
	if !ok {
		serr = ErrUnhandled
	}

	return &api.ErrorStatusCode{
		StatusCode: serr.HttpStatusCode,
		Response: api.Error{
			Code:    serr.ErrorCode,
			Message: serr.ErrorMessage,
		},
	}
}
