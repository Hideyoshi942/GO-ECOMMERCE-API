package routers

import (
	c "go-ecomerce-backend-api/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

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
