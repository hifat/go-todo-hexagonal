package gormrepo

import (
	"github.com/hifat/go-todo-hexagonal/internal/repository"
	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.AuthRepository {
	return userRepositoryDB{db}
}

func (r userRepositoryDB) Register(register repository.Register) (*repository.Auth, error) {
	newUser := repository.User{
		Username: register.Username,
		Password: register.Password,
		Name:     register.Name,
	}

	tx := r.db.Create(&newUser)
	if tx.Error != nil {
		return nil, tx.Error
	}

	user := repository.User{
		ID:        newUser.ID,
		Username:  newUser.Username,
		Password:  newUser.Password,
		Name:      newUser.Name,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}

	auth := repository.Auth{
		User:  user,
		Token: "",
	}

	return &auth, nil
}

func (r userRepositoryDB) Login(login repository.Login) (*repository.Auth, error) {
	credentials := repository.User{
		Username: login.Username,
	}

	tx := r.db.Where("username = ?", credentials.Username).Find(&credentials)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// compare password

	return nil, nil
}

func (r userRepositoryDB) Me(token string) (*repository.Auth, error) {
	return nil, nil
}
