package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
}

func (pr *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// public router
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/search", func(c *gin.Context) {
			c.JSON(http.StatusNotImplemented, gin.H{"message": "search endpoint is not implemented yet"})
		})
		productRouterPublic.GET("/detail/:id", func(c *gin.Context) {
			c.JSON(http.StatusNotImplemented, gin.H{"message": "detail endpoint is not implemented yet"})
		})
	}

	// private router
	productRouterPrivate := Router.Group("/product")
	{
		productRouterPrivate.GET("/search", func(c *gin.Context) {
			c.JSON(http.StatusNotImplemented, gin.H{"message": "private search endpoint is not implemented yet"})
		})
		productRouterPrivate.GET("/detail/:id", func(c *gin.Context) {
			c.JSON(http.StatusNotImplemented, gin.H{"message": "private detail endpoint is not implemented yet"})
		})
	}
}
