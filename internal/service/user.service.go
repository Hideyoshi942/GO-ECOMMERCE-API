package service

import (
	"go-ecomerce-backend-api/internal/repo"
	"go-ecomerce-backend-api/pkg/response"
)

//type UserService struct {
//	userRepo *repo.UserRepo
//}
//
//func NewUserService() *UserService {
//	return &UserService{
//		userRepo: repo.NewUserRepo(),
//	}
//}
//
//func (us *UserService) GetInfoUserUserService() string {
//	return us.userRepo.GetInfoUser()
//}

// Interface_version

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) Register(email string, purpose string) int {
	// 1. check email exists
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	return response.ErrCodeSuccess
}
