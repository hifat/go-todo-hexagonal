package rules

import (
	"github.com/go-playground/validator"
	"github.com/hifat/go-todo-hexagonal/internal/resource"
)

type Errors map[string]string

func Validate(err validator.FieldError) map[string]string {
	rulesMSG := map[string]string{
		"required": "field " + err.Field() + " has required",
	}

	msg := map[string]string{
		resource.Fields[err.StructField()]: rulesMSG[err.Tag()],
	}

	return msg
}
