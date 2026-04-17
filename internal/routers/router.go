package routers

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(mode string) *gin.Engine {
	if mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		return gin.Default()
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	return r
}
