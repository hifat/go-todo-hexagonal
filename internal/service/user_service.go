package service

import (
	"github.com/hifat/go-todo-hexagonal/helper/errs"
	"github.com/hifat/go-todo-hexagonal/helper/zlog"
	"github.com/hifat/go-todo-hexagonal/internal/repository"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (u userService) Update(id string, user EditUser) (*User, error) {
	errValidate := validateForm(user)
	if errValidate != nil {
		return nil, errValidate
	}

	editUser := repository.EditUser{
		Username: user.Username,
		Name:     user.Name,
	}

	updateUser, err := u.userRepo.Update(id, editUser)

	if err != nil {
		zlog.Error(err)
		return nil, errs.HttpError(err)
	}

	userResponse := User{
		ID:        updateUser.ID,
		Username:  updateUser.Username,
		Name:      updateUser.Name,
		CreatedAt: updateUser.CreatedAt,
		UpdatedAt: updateUser.UpdatedAt,
	}

	return &userResponse, nil
}
