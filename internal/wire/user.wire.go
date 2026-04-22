//go:build wire

package wire

import (
	"go-ecomerce-backend-api/internal/controller"
	"go-ecomerce-backend-api/internal/repo"
	"go-ecomerce-backend-api/internal/service"

	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(controller.NewUserController, service.NewUserService, repo.NewUserRepository)
	return new(controller.UserController), nil
}
