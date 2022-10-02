package gormrepo

import (
	"strconv"

	"github.com/hifat/go-todo-hexagonal/helper/errs"
	"github.com/hifat/go-todo-hexagonal/internal/repository"
	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserGorm(db *gorm.DB) repository.UserRepository {
	return &userRepositoryDB{db}
}

func (r userRepositoryDB) Update(id string, editUser repository.EditUser) (*repository.User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errs.NaN("UserID")
	}

	var user repository.User
	tx := r.db.First(&user, userID)
	if tx.Error != nil {
		return nil, errHandler(tx.Error)
	}

	tx = r.db.Model(&user).Updates(map[string]interface{}{
		"username": editUser.Username,
		"name":     editUser.Name,
	})
	if tx.Error != nil {
		return nil, errHandler(tx.Error)
	}

	return &user, nil
}
