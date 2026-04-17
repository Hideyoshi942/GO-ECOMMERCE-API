package manager

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	//// public router
	//userRouterPublic := Router.Group("/admin/user")
	//{
	//	userRouterPublic.GET("/search")
	//	userRouterPublic.GET("/detail/:id")
	//}

	// private router
	userRouterPrivate := Router.Group("/admin/user")
	//userRouterPrivate.Use(middlewares.Limiter())
	//userRouterPrivate.Use(middlewares.JWTAuth()).Use(middlewares.CORS())
	//userRouterPrivate.Use(middlewares.PermissionCheck())
	{
		userRouterPrivate.POST("/active_user", func(c *gin.Context) {
			c.JSON(http.StatusNotImplemented, gin.H{"message": "manager active_user endpoint is not implemented yet"})
		})
	}
}
