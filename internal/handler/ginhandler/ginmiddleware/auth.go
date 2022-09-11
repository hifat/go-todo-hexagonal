package ginmiddleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hifat/go-todo-hexagonal/helper/zlog"
	"github.com/hifat/go-todo-hexagonal/internal/token"
)

func Auth(tokenMaker token.Maker) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		accessToken := strings.TrimPrefix(authHeader, "Bearer ")
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			zlog.Error(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
		}

		ctx.Set("Authorization", payload)
		ctx.Next()
	}
}
