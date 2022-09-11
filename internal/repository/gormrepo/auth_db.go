package gormrepo

import (
	"github.com/hifat/go-todo-hexagonal/internal/repository"
	"gorm.io/gorm"
)

type authRepositoryDB struct {
	db *gorm.DB
}

func NewAuthGorm(db *gorm.DB) repository.AuthRepository {
	return authRepositoryDB{db}
}

type user struct {
	Username string
	Password string
	Name     string
}

func (r authRepositoryDB) Register(register repository.Register) (*repository.Auth, error) {
	newUser := user{
		Username: register.Username,
		Password: register.Password,
		Name:     register.Name,
	}

	var userRepo repository.User

	tx := r.db.Create(&newUser).Find(&userRepo)
	if tx.Error != nil {
		return nil, tx.Error
	}

	auth := repository.Auth{
		User:        userRepo,
		AccessToken: "",
	}

	return &auth, nil
}

func (r authRepositoryDB) Login(login repository.Login) (*repository.Auth, error) {
	credentials := repository.User{
		Username: login.Username,
	}

	tx := r.db.Where("username = ?", credentials.Username).Find(&credentials)
	if tx.Error != nil {
		return nil, tx.Error
	}

	user := repository.User{
		ID:        credentials.ID,
		Username:  credentials.Username,
		Password:  credentials.Password,
		Name:      credentials.Name,
		CreatedAt: credentials.CreatedAt,
		UpdatedAt: credentials.UpdatedAt,
	}

	auth := repository.Auth{
		User:        user,
		AccessToken: "",
	}

	return &auth, nil
}

func (r authRepositoryDB) Me(token string) (*repository.Auth, error) {
	return nil, nil
}
