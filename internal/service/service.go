package service

import (
	"github.com/hifat/go-todo-hexagonal/helper/errs"
	"github.com/hifat/go-todo-hexagonal/helper/rules"
	"github.com/hifat/go-todo-hexagonal/helper/zlog"
)

func validateForm(st interface{}) *errs.AppError {
	errFields, err := rules.Validate(st)
	if err != nil {
		zlog.Error(err)
		return errs.Unexpected()
	}

	if errFields != nil {
		return errs.UnprocessableEntity(errFields)
	}

	return nil
}
