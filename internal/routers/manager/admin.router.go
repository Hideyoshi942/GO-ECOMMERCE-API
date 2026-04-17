package manager

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminRouter struct {
}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusNotImplemented, gin.H{"message": "admin login endpoint is not implemented yet"})
		})
	}

	// private router
	adminRouterPrivate := Router.Group("/admin/user")
	//adminRouterPrivate.Use(middlewares.Limiter())
	//adminRouterPrivate.Use(middlewares.JWTAuth()).Use(middlewares.CORS())
	//adminRouterPrivate.Use(middlewares.PermissionCheck())
	{
		adminRouterPrivate.POST("/active_user", func(c *gin.Context) {
			c.JSON(http.StatusNotImplemented, gin.H{"message": "admin active_user endpoint is not implemented yet"})
		})
	}
}
