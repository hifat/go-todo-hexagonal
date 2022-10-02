package service

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/hifat/go-todo-hexagonal/helper/errs"
	"github.com/hifat/go-todo-hexagonal/helper/zlog"
	"github.com/hifat/go-todo-hexagonal/internal/repository"
	"github.com/hifat/go-todo-hexagonal/internal/token"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var jwtMaker token.Maker

type authService struct {
	db repository.AuthRepository
}

func init() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		zlog.Error(err)
		return
	}

	jwtMaker, err = token.NewJWTMaker(os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		zlog.Error(err)
		return
	}
}

func NewAuthService(db repository.AuthRepository) AuthService {
	return authService{db}
}

// Function for create session
func createSession(r authService, user User, login Login) (*Auth, error) {
	userPayload := token.UserPayload{
		UserID:   user.ID,
		Username: user.Username,
	}

	accessToken, accessPayload, err := jwtMaker.CreateToken(userPayload, 15*time.Minute)
	if err != nil {
		zlog.Error(err)
		return nil, errs.Unexpected()
	}

	refreshToken, refreshPayload, err := jwtMaker.CreateToken(userPayload, 24*time.Hour)
	if err != nil {
		zlog.Error(err)
		return nil, errs.Unexpected()
	}

	newSession := repository.NewSession{
		ID:           fmt.Sprintf("%v", refreshPayload.ID),
		Username:     userPayload.Username,
		RefreshToken: refreshToken,
		UserAgent:    login.UserAgent,
		ClientIP:     login.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	}

	session, err := r.db.CreateSession(newSession)
	if err != nil {
		zlog.Error(err)
		return nil, errs.Unexpected()
	}

	auth := Auth{
		User:                  user,
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
	}

	return &auth, nil
}

func (r authService) Register(register Register) (*Auth, error) {
	errValidate := validateForm(register)
	if errValidate != nil {
		return nil, errValidate
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		zlog.Error(err)
		return nil, errs.Unexpected()
	}

	registerRepo := repository.Register{
		Username: register.Username,
		Password: string(hashPassword),
		Name:     register.Name,
	}

	registerDB, err := r.db.Register(registerRepo)
	if err != nil {
		zlog.Error(err)
		return nil, errs.Unexpected()
	}

	user := User{
		ID:        registerDB.ID,
		Username:  registerDB.Username,
		Name:      registerDB.Name,
		CreatedAt: registerDB.CreatedAt,
		UpdatedAt: registerDB.UpdatedAt,
	}

	login := Login{
		Username:  register.Username,
		Password:  register.Password,
		UserAgent: register.UserAgent,
		ClientIP:  register.ClientIP,
	}

	auth, err := createSession(r, user, login)
	if err != nil {
		return nil, errs.Unexpected()
	}

	return auth, nil
}

func (r authService) Login(login Login) (*Auth, error) {
	errValidate := validateForm(login)
	if errValidate != nil {
		return nil, errValidate
	}

	loginRepo := repository.Login{
		Username: login.Username,
		Password: login.Password,
	}

	loginDB, err := r.db.Login(loginRepo)
	if err != nil {
		zlog.Error(err)

		if errors.Is(err, errs.ErrRecordNotFound) {
			return nil, errs.Unauthorizetion("Invalid credentials")
		}

		return nil, errs.HttpError(err)
	}

	user := User{
		ID:        loginDB.ID,
		Username:  loginDB.Username,
		Name:      loginDB.Name,
		CreatedAt: loginDB.CreatedAt,
		UpdatedAt: loginDB.UpdatedAt,
	}

	auth, err := createSession(r, user, login)
	if err != nil {
		return nil, errs.Unexpected()
	}

	return auth, nil
}

func (r authService) Me(accessToken string) (*Auth, error) {
	payload, err := jwtMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, errs.Unexpected()
	}

	meDB, err := r.db.Me(payload.Username)
	if err != nil {
		return nil, errs.Unexpected()
	}

	user := Auth{
		User: User{
			ID:        meDB.ID,
			Username:  meDB.Username,
			Name:      meDB.Name,
			CreatedAt: meDB.CreatedAt,
			UpdatedAt: meDB.UpdatedAt,
		},
		AccessToken: accessToken,
	}

	return &user, nil
}
