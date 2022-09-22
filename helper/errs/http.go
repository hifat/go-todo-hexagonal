package errs

import (
	"net/http"

	"github.com/hifat/go-todo-hexagonal/helper/rules"
)

const (
	Unauthorized = "unauthorized"
)

type AppError struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Errors  []rules.Errors `json:"errors"`
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

func UnprocessableEntity(errors []rules.Errors) AppError {
	return AppError{
		Code:    http.StatusUnprocessableEntity,
		Message: "unprocessable entity",
		Errors:  errors,
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
		Message: "unexpected error",
	}
}
