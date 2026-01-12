package models

type ApiError struct {
	Error string
}

func NewApiError(error string) *ApiError {
	return &ApiError{Error: error}
}
