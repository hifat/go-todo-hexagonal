package rules

import (
	"github.com/go-playground/validator"
	"github.com/hifat/go-todo-hexagonal/helper/zlog"
	"github.com/hifat/go-todo-hexagonal/internal/resource"
)

type FieldError struct {
	Message string              `json:"message"`
	Errors  map[string][]string `json:"error"`
}

var validate *validator.Validate

func handleMessage(err validator.FieldError) (field string, msg string) {
	field = resource.Fields[err.Field()]

	rulesMSG := map[string]string{
		"required": "field " + field + " has required",
	}

	return field, rulesMSG[err.Tag()]
}

func Validate(st interface{}) (map[string][]string, error) {
	validate = validator.New()
	err := validate.Struct(st)
	errFields := make(map[string][]string)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			zlog.Error(err)
			return nil, err
		}

		for _, err := range err.(validator.ValidationErrors) {
			field, errMSG := handleMessage(err)
			errFields[field] = append(errFields[field], errMSG)
		}

		return errFields, nil
	}

	return nil, nil
}
