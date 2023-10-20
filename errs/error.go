package errs

import "net/http"

type ApiError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
}

func (e ApiError) AsMessage() *ApiError {
	return &ApiError{
		Message: e.Message,
	}
}

func NewNotFoundError(message string) *ApiError {
	return &ApiError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewUnexpectedError(message string) *ApiError {
	return &ApiError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}

func NewValidationError(message string) *ApiError {
	return &ApiError{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
	}
}
