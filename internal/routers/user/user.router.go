package user

import (
	"go-ecomerce-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	// public router
	// this is non-dependency
	//ur := repo.NewUserRepository()
	//us := service.NewUserService(ur)
	//userHandlerNonDependency := controller.NewUserController(us)
	userController, _ := wire.InitUserRouterHandler()

	// Wire go
	// Dependency Injection (DI)

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")
	}

	// private router
	userRouterPrivate := Router.Group("/user")
	//userRouterPrivate.Use(middlewares.Limiter())
	//userRouterPrivate.Use(middlewares.JWTAuth()).Use(middlewares.CORS())
	//userRouterPrivate.Use(middlewares.PermissionCheck())
	{
		userRouterPrivate.GET("/get_info")
	}
}
