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
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *ApiError {
	return &ApiError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
