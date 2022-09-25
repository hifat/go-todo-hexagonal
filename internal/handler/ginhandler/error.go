package ginhandler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hifat/go-todo-hexagonal/helper/errs"
)

func handlerError(ctx *gin.Context, err error) {
	switch e := err.(type) {
	case *errs.AppError:
		fmt.Print("here 1")
		ctx.JSON(e.Code, e)
	case error:
		fmt.Print("here 2")
		ctx.JSON(http.StatusInternalServerError, e)
	}
}
