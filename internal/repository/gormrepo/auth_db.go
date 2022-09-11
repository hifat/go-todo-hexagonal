package gormrepo

import (
	"fmt"
	"time"

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
	ID        uint
	Username  string
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r authRepositoryDB) Register(register repository.Register) (*repository.Auth, error) {
	newUser := user{
		Username: register.Username,
		Password: register.Password,
		Name:     register.Name,
	}

	tx := r.db.Create(&newUser)
	if tx.Error != nil {
		return nil, tx.Error
	}

	auth := repository.Auth{
		User: repository.User{
			ID:        fmt.Sprintf("%d", newUser.ID),
			Username:  newUser.Username,
			Password:  newUser.Password,
			Name:      newUser.Name,
			CreatedAt: newUser.CreatedAt,
			UpdatedAt: newUser.UpdatedAt,
		},
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
		User: user,
	}

	return &auth, nil
}

func (r authRepositoryDB) Me(token string) (*repository.Auth, error) {
	return nil, nil
}
