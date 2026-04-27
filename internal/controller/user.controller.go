package controller

import (
	"fmt"
	"go-ecomerce-backend-api/internal/service"
	"go-ecomerce-backend-api/internal/vo"
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
	var params = vo.UserRegistratorRequest{}
	if err := c.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(c, response.ErrCodeParamInvalid, err.Error())
		return
	}
	fmt.Printf("Email: %s, Purpose: %s\n", params.Email, params.Purpose)
	result := uc.userService.Register(params.Email, params.Purpose)
	response.SuccessResponse(c, result, "Success", nil)
}
