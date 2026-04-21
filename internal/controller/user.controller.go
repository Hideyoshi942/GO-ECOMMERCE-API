package controller

import (
	"go-ecomerce-backend-api/internal/service"
	"go-ecomerce-backend-api/pkg/response"

	"github.com/gin-gonic/gin"
)

//type UserController struct {
//	userService *service.UserService
//}
//
//func NewUserController() *UserController {
//	return &UserController{
//		userService: service.NewUserService(),
//	}
//}
//
//func (uc *UserController) GetUserById(c *gin.Context) {
//
//	response.SuccessResponse(c, 20001, "User retrieved successfully", []string{"User retrieved successfully"})
//}

// Interface_Version
type UserController struct {
	userService service.IUserService
}

func NewUserController(
	userService service.IUserService,
) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(c, result, "Success", nil)
}
