package errors

import "fmt"

type APIError struct {
	StatusCode int    `json:"-"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Details    any    `json:"details,omitempty"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func NewBadRequestError(code, message string, details ...any) *APIError {
	var detailsData any
	if len(details) > 0 {
		detailsData = details[0]
	}
	return &APIError{
		StatusCode: 400,
		Code:       code,
		Message:    message,
		Details:    detailsData,
	}
}

func NewNotFoundError(code, message string) *APIError {
	return &APIError{
		StatusCode: 404,
		Code:       code,
		Message:    message,
	}
}

func NewInternalServerError(message string) *APIError {
	return &APIError{
		StatusCode: 500,
		Code:       "internal_server_error",
		Message:    message,
	}
}
