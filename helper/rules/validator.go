package rules

import (
	"github.com/go-playground/validator"
	"github.com/hifat/go-todo-hexagonal/internal/resource"
)

type Errors map[string]string

func Validate(err validator.FieldError) map[string]string {
	field := resource.Fields[err.Field()]

	rulesMSG := map[string]string{
		"required": "field " + field + " has required",
	}

	msg := map[string]string{
		field: rulesMSG[err.Tag()],
	}

	return msg
}
