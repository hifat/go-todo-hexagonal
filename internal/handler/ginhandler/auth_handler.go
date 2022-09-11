package ginhandler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hifat/go-todo-hexagonal/internal/handler"
	"github.com/hifat/go-todo-hexagonal/internal/service"
)

type authHandler struct {
	authSrv service.AuthService
}

func NewAuthHandler(authSrv service.AuthService) authHandler {
	return authHandler{authSrv}
}

func (t authHandler) Register(ctx *gin.Context) {
	var reqAuth handler.Register
	err := ctx.ShouldBind(&reqAuth)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	registerServ := service.Register{
		Username: reqAuth.Username,
		Password: reqAuth.Password,
		Name:     reqAuth.Name,
	}

	register, err := t.authSrv.Register(registerServ)
	if err != nil {
		handlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": register,
	})
}

func (t authHandler) Login(ctx *gin.Context) {
	var reqAuth handler.Login
	err := ctx.ShouldBind(&reqAuth)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	loginServ := service.Login{
		Username: reqAuth.Username,
		Password: reqAuth.Password,
	}

	login, err := t.authSrv.Login(loginServ)
	if err != nil {
		handlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": login,
	})
}

func (t authHandler) Me(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	accessToken := strings.TrimPrefix(authHeader, "Bearer ")

	login, err := t.authSrv.Me(accessToken)
	if err != nil {
		handlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": login,
	})
}
