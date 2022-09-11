package ginhandler

import (
	"net/http"

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
