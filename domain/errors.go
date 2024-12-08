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
	ErrSecurity           = RecipeServiceError{HttpStatusCode: 403, Code: 3, Message: "authentication required"}
	ErrInvalidCredentials = RecipeServiceError{HttpStatusCode: 403, Code: 4, Message: "invalid credentials"}
	ErrRegistration       = RecipeServiceError{HttpStatusCode: 422, Code: 5, Message: "failed to register user"}
)
