package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {
	fmt.Println("--->> My handler")
	name := c.DefaultQuery("name", "hideyoshi")
	//name := c.Param("name")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong" + name,
		"uid":     uid,
		"users":   []string{"cr7", "m10", "hideyoshi"},
	})
}
