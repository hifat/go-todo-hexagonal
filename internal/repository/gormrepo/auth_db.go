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
	credentials := user{
		Username: login.Username,
	}

	tx := r.db.Where("username = ?", credentials.Username).First(&credentials)
	if tx.Error != nil {
		return nil, tx.Error
	}

	user := repository.User{
		ID:        fmt.Sprintf("%v", credentials.ID),
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

func (r authRepositoryDB) Me(username string) (*repository.Auth, error) {
	credentials := user{
		Username: username,
	}

	tx := r.db.Where("username = ?", credentials.Username).First(&credentials)
	if tx.Error != nil {
		return nil, tx.Error
	}

	user := repository.User{
		ID:        fmt.Sprintf("%v", credentials.ID),
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

func (s authRepositoryDB) ShowSession(id string) (*repository.Session, error) {
	var session repository.Session

	tx := s.db.Where("id = ?", id).First(&session)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &session, nil
}

type Session struct {
	ID           string
	Username     string
	RefreshToken string
	UserAgent    string
	ClientIP     string
	IsBlocked    bool
	ExpiresAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (s authRepositoryDB) CreateSession(newSession repository.NewSession) (*repository.Session, error) {
	createSession := Session{
		ID:           newSession.ID,
		Username:     newSession.Username,
		RefreshToken: newSession.RefreshToken,
		UserAgent:    newSession.UserAgent,
		ClientIP:     newSession.ClientIP,
		IsBlocked:    newSession.IsBlocked,
		ExpiresAt:    newSession.ExpiresAt,
	}

	tx := s.db.Create(&createSession)
	if tx.Error != nil {
		return nil, tx.Error
	}

	session := repository.Session{
		ID:           createSession.ID,
		Username:     createSession.Username,
		RefreshToken: createSession.RefreshToken,
		UserAgent:    createSession.UserAgent,
		ClientIP:     createSession.ClientIP,
		IsBlocked:    createSession.IsBlocked,
		ExpiresAt:    createSession.ExpiresAt,
		CreatedAt:    createSession.CreatedAt,
		UpdatedAt:    createSession.UpdatedAt,
	}

	return &session, nil
}

func (s authRepositoryDB) DeleteSession(id string) error {
	var session Session
	tx := s.db.Where("id = ?", id).Delete(&session)
	return tx.Error
}
