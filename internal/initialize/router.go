package initialize

import (
	"go-ecomerce-backend-api/global"
	"go-ecomerce-backend-api/internal/controller"
	"go-ecomerce-backend-api/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := routers.NewRouter(global.Config.Server.Mode)
	pongController := controller.NewPongController()

	// middlewares
	//r.Use() // logging
	//r.Use() // cross
	//r.Use() // limiter global
	manageRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/api/v1/2026")
	{
		MainGroup.GET("/checkStatus", pongController.Pong) // tracking monitor
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
		manageRouter.InitAdminUserRouter(MainGroup)
	}
	return r
}
