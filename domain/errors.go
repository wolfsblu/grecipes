package domain

import (
	"fmt"
)

type RecipeServiceError struct {
	HttpStatusCode int
	Code           int
	Message        string
	Children       []error
}

var (
	ErrUnhandled          = &RecipeServiceError{HttpStatusCode: 500, Code: 1, Message: "internal server error"}
	ErrNotFound           = &RecipeServiceError{HttpStatusCode: 404, Code: 2, Message: "the requested resource was not found"}
	ErrSecurity           = &RecipeServiceError{HttpStatusCode: 403, Code: 3, Message: "authentication required"}
	ErrInvalidCredentials = &RecipeServiceError{HttpStatusCode: 403, Code: 4, Message: "invalid credentials"}
	ErrRegistration       = &RecipeServiceError{HttpStatusCode: 422, Code: 5, Message: "failed to register user"}
	ErrPersistence        = &RecipeServiceError{HttpStatusCode: 500, Code: 6, Message: "failed to persist data"}
	ErrRetrieval          = &RecipeServiceError{HttpStatusCode: 500, Code: 7, Message: "failed to retrieve data"}
)

func (e *RecipeServiceError) Error() string {
	return e.Message
}

func WrapError(err error, message string, root *RecipeServiceError) error {
	return fmt.Errorf("%w: %s: %w", root, message, err)
}
