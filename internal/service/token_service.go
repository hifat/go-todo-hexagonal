package service

import (
	"fmt"
	"time"

	"github.com/hifat/go-todo-hexagonal/helper/errs"
	"github.com/hifat/go-todo-hexagonal/helper/zlog"
	"github.com/hifat/go-todo-hexagonal/internal/repository"
	"github.com/hifat/go-todo-hexagonal/internal/token"
)

type tokenService struct {
	authRepo repository.AuthRepository
	jwtMaker token.Maker
}

func NewTokenService(authRepo repository.AuthRepository, jwtMaker token.Maker) TokenService {
	return &tokenService{authRepo, jwtMaker}
}

func (t tokenService) RenewAccessToken(renewAccessTokenReq RenewAccessTokenRequest, userDevice UserDevice) (*RenewAccessTokenResponse, error) {
	refreshPayload, err := t.jwtMaker.VerifyToken(renewAccessTokenReq.RefreshToken)
	if err != nil {
		zlog.Error(err)
		return nil, errs.Unauthorizetion(errs.Unauthorized)
	}

	session, err := t.authRepo.ShowSession(fmt.Sprintf("%v", refreshPayload.ID))
	if err != nil {
		zlog.Error(err)
		return nil, errs.HttpError(err)
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

	accessToken, accessPayload, err := t.jwtMaker.CreateToken(userPayload, 5*time.Minute)
	if err != nil {
		zlog.Error(err)
		return nil, errs.Unexpected()
	}

	newRefreshToken, newRefreshPayload, err := t.jwtMaker.CreateToken(userPayload, 24*time.Hour)
	if err != nil {
		zlog.Error(err)
		return nil, errs.Unexpected()
	}

	err = t.authRepo.DeleteSession(refreshPayload.ID.String())
	if err != nil {
		zlog.Error(err)
		return nil, errs.HttpError(err)
	}

	newSessionRepo := repository.NewSession{
		ID:           newRefreshPayload.ID.String(),
		Username:     newRefreshPayload.Username,
		RefreshToken: newRefreshToken,
		UserAgent:    userDevice.UserAgent,
		ClientIP:     userDevice.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    newRefreshPayload.ExpiredAt,
	}

	_, err = t.authRepo.CreateSession(newSessionRepo)
	if err != nil {
		return nil, errs.HttpError(err)
	}

	rsp := RenewAccessTokenResponse{
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          newRefreshToken,
		RefreshTokenExpiresAt: newRefreshPayload.ExpiredAt,
	}

	return &rsp, nil
}
