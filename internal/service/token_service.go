package service

import (
	"fmt"
	"time"

	"github.com/hifat/go-todo-hexagonal/helper/errs"
	"github.com/hifat/go-todo-hexagonal/helper/zlog"
	"github.com/hifat/go-todo-hexagonal/internal/repository"
	"github.com/hifat/go-todo-hexagonal/internal/token"
	"gorm.io/gorm"
)

type tokenService struct {
	db       repository.AuthRepository
	jwtMaker token.Maker
}

func NewTokenService(db repository.AuthRepository, jwtMaker token.Maker) TokenService {
	return &tokenService{db, jwtMaker}
}

func (t tokenService) RenewAccessToken(renewAccessTokenReq RenewAccessTokenRequest) (*RenewAccessTokenResponse, error) {
	refreshPayload, err := t.jwtMaker.VerifyToken(renewAccessTokenReq.RefreshToken)
	if err != nil {
		zlog.Error(err)
		return nil, errs.Unauthorizetion(errs.Unauthorized)
	}

	session, err := t.db.ShowSession(fmt.Sprintf("%v", refreshPayload.ID))
	if err != nil {
		zlog.Error(err)
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NotFound(err.Error())
		}

		return nil, errs.Unexpected()
	}

	if session.IsBlocked {
		zlog.Error(err)
		return nil, errs.Unauthorizetion("blocked session")
	}

	if session.Username != refreshPayload.Username {
		zlog.Error(err)
		return nil, errs.Unauthorizetion("incorrect session user")
	}

	if session.RefreshToken != renewAccessTokenReq.RefreshToken {
		zlog.Error(err)
		return nil, errs.Unauthorizetion("mismatched session token")
	}

	if time.Now().After(session.ExpiresAt) {
		zlog.Error(err)
		return nil, errs.Unauthorizetion("expired session")
	}

	userPayload := token.UserPayload{
		UserID:   refreshPayload.UserID,
		Username: refreshPayload.Username,
	}

	accessToken, accessPayload, err := t.jwtMaker.CreateToken(userPayload, 24*time.Hour)
	if err != nil {
		zlog.Error(err)
		return nil, errs.Unexpected()
	}

	rsp := RenewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}

	return &rsp, nil
}
