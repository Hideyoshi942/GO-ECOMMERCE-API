package controller

import (
	"go-ecomerce-backend-api/internal/service"
	"go-ecomerce-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

func (uc *UserController) GetUserById(c *gin.Context) {

	response.SuccessResponse(c, 20001, "User retrieved successfully", []string{"User retrieved successfully"})
}
