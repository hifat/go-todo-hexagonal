package ginhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/go-todo-hexagonal/internal/handler"
	"github.com/hifat/go-todo-hexagonal/internal/service"
	"github.com/hifat/go-todo-hexagonal/internal/token"
)

type userHandler struct {
	userSrv service.UserService
}

func NewUserHandler(userSrv service.UserService) userHandler {
	return userHandler{userSrv}
}

func (t userHandler) Update(ctx *gin.Context) {
	userAuth := ctx.MustGet("user").(*token.Payload)

	var userReq handler.EditUser
	err := ctx.ShouldBind(&userReq)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	editUser := service.EditUser{
		Username: userReq.Username,
		Name:     userReq.Name,
	}

	user, err := t.userSrv.Update(userAuth.UserID, editUser)
	if err != nil {
		handlerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
