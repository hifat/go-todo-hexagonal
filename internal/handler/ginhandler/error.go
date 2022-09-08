package ginhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/go-todo-hexagonal/helper/errs"
)

func handlerError(ctx *gin.Context, err error) {
	switch e := err.(type) {
	case errs.AppError:
		ctx.JSON(e.Code, gin.H{
			"message": e.Message,
		})
	case error:
		ctx.JSON(http.StatusInternalServerError, e)
	}
}
