package user

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register") // register -> YES <-> NO
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
