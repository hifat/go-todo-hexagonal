package errs

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func NotFound(message string) AppError {
	return AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func BadRequest(message string) AppError {
	return AppError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func Unauthorizetion(message string) AppError {
	return AppError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func Unexpected() AppError {
	return AppError{
		Code:    http.StatusInternalServerError,
		Message: "Unexpected error",
	}
}
