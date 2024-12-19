package handlers

import (
	"context"
	"errors"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/wolfsblu/go-chef/api"
	"github.com/wolfsblu/go-chef/domain"
)

func (h *RecipeHandler) NewError(_ context.Context, err error) (r *api.ErrorStatusCode) {
	var serviceError *domain.RecipeServiceError
	var securityError *ogenerrors.SecurityError

	if errors.As(err, &securityError) {
		serviceError = domain.ErrSecurity
	} else if !errors.As(err, &serviceError) {
		serviceError = domain.ErrUnhandled
	}

	return &api.ErrorStatusCode{
		StatusCode: serviceError.HttpStatusCode,
		Response: api.Error{
			Code:    serviceError.Code,
			Message: err.Error(),
		},
	}
}
