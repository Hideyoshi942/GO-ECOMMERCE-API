package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	// private router
	adminRouterPrivate := Router.Group("/admin/user")
	//adminRouterPrivate.Use(middlewares.Limiter())
	//adminRouterPrivate.Use(middlewares.JWTAuth()).Use(middlewares.CORS())
	//adminRouterPrivate.Use(middlewares.PermissionCheck())
	{
		adminRouterPrivate.POST("/active_user")
	}
}
