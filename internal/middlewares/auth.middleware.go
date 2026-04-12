package middlewares

import (
	"go-ecomerce-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(context, response.ErrInvalidToken, "")
			context.Abort()
			return
		}

		context.Next()
	}
}
