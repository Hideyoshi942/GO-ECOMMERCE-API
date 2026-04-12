package routers

import (
	"fmt"
	c "go-ecomerce-backend-api/internal/controller"
	"go-ecomerce-backend-api/internal/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("Before --> AA")
		context.Next()
		fmt.Println("Alter --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("Before --> BB")
		context.Next()
		fmt.Println("Alter --> BB")
	}
}

func CC(context *gin.Context) {
	fmt.Println("Before --> CC")
	context.Next()
	fmt.Println("Alter --> CC")
}

func NewRouter() *gin.Engine {
	r := gin.Default()

	// use the middleware
	r.Use(middlewares.AuthenMiddleware(), BB(), CC)

	v1 := r.Group("/v1/2026")
	{
		v1.GET("/pingpong", c.NewPongController().Pong)
		v1.GET("/pinguser", c.NewUserController().GetUserById)
		v1.POST("/ping", Pong)
		v1.PATCH("/ping", Pong)
		v1.DELETE("/ping", Pong)
		v1.HEAD("/ping", Pong)
		v1.OPTIONS("/ping", Pong)

	}

	v2 := r.Group("/v2/2026")
	{
		v2.GET("/ping", Pong)
		v2.POST("/ping", Pong)
		v2.PATCH("/ping", Pong)
		v2.DELETE("/ping", Pong)
		v2.HEAD("/ping", Pong)
		v2.OPTIONS("/ping", Pong)

	}

	return r
}

func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "hideyoshi")
	//name := c.Param("name")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"uid":     uid,
		"users":   []string{"cr7", "m10", "hideyoshi"},
	})

}
