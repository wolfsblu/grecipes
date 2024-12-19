package domain

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
	ErrUserNotFound       = RecipeServiceError{HttpStatusCode: 404, Code: 3, Message: "user not found"}
	ErrSecurity           = RecipeServiceError{HttpStatusCode: 403, Code: 4, Message: "authentication required"}
	ErrInvalidCredentials = RecipeServiceError{HttpStatusCode: 403, Code: 5, Message: "invalid credentials"}
	ErrRegistration       = RecipeServiceError{HttpStatusCode: 422, Code: 6, Message: "failed to register user"}
	ErrPersistence        = RecipeServiceError{HttpStatusCode: 500, Code: 7, Message: "failed to persist data"}
	ErrRetrieval          = RecipeServiceError{HttpStatusCode: 500, Code: 8, Message: "failed to retrieve data"}
)
