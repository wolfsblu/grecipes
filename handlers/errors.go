package handlers

import (
	"context"
	"errors"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/wolfsblu/go-chef/api"
)

type RecipeServiceError struct {
	HttpStatusCode int
	Code           int
	Message        string
}

func (e *RecipeServiceError) Error() string {
	return e.Message
}

var (
	ErrUnhandled          = RecipeServiceError{HttpStatusCode: 500, Code: 1, Message: "internal server error"}
	ErrRecipeNotFound     = RecipeServiceError{HttpStatusCode: 404, Code: 2, Message: "recipe not found"}
	ErrSecurity           = RecipeServiceError{HttpStatusCode: 403, Code: 3, Message: "authentication required"}
	ErrInvalidCredentials = RecipeServiceError{HttpStatusCode: 403, Code: 4, Message: "invalid credentials"}
	ErrRegistration       = RecipeServiceError{HttpStatusCode: 422, Code: 5, Message: "failed to register user"}
)

func (h *RecipeHandler) NewError(_ context.Context, err error) (r *api.ErrorStatusCode) {
	var serviceError *RecipeServiceError
	var securityError *ogenerrors.SecurityError

	if errors.As(err, &securityError) {
		serviceError = &ErrSecurity
	} else if !errors.As(err, &serviceError) {
		serviceError = &ErrUnhandled
	}

	return &api.ErrorStatusCode{
		StatusCode: serviceError.HttpStatusCode,
		Response: api.Error{
			Code:    serviceError.Code,
			Message: err.Error(),
		},
	}
}
