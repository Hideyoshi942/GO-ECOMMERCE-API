package initialize

import (
	"go-ecomerce-backend-api/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	return routers.NewRouter()
}
