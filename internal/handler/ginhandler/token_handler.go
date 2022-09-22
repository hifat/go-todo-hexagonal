package ginhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/go-todo-hexagonal/internal/handler"
	"github.com/hifat/go-todo-hexagonal/internal/service"
)

type tokenHandler struct {
	tokenSrv service.TokenService
}

func NewTokenHandler(tokenSrv service.TokenService) tokenHandler {
	return tokenHandler{tokenSrv}
}

func (t tokenHandler) RenewAccessToken(ctx *gin.Context) {
	var req handler.RenewAccessTokenRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	renewAccessToken := service.RenewAccessTokenRequest{
		RefreshToken: req.RefreshToken,
	}

	userDevice := service.UserDevice{
		UserAgent: ctx.Request.UserAgent(),
		ClientIP:  ctx.ClientIP(),
	}

	accessToken, err := t.tokenSrv.RenewAccessToken(renewAccessToken, userDevice)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": accessToken,
	})
}
