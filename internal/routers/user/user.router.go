package user

import (
	"go-ecomerce-backend-api/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController := controller.NewUserController()

	// public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", func(c *gin.Context) {
			c.JSON(http.StatusNotImplemented, gin.H{"message": "register endpoint is not implemented yet"})
		})
		userRouterPublic.POST("/otp", func(c *gin.Context) {
			c.JSON(http.StatusNotImplemented, gin.H{"message": "otp endpoint is not implemented yet"})
		})
	}

	// private router
	userRouterPrivate := Router.Group("/user")
	//userRouterPrivate.Use(middlewares.Limiter())
	//userRouterPrivate.Use(middlewares.JWTAuth()).Use(middlewares.CORS())
	//userRouterPrivate.Use(middlewares.PermissionCheck())
	{
		userRouterPrivate.GET("/get_info", userController.GetUserById)
	}
}
