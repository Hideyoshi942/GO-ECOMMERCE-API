package repo

import (
	"go-ecomerce-backend-api/global"
	"go-ecomerce-backend-api/internal/models"
)

//type UserRepo struct{}
//
//func NewUserRepo() *UserRepo {
//	return &UserRepo{}
//}
//
//func (ur *UserRepo) GetInfoUser() string {
//	return "hideyoshi"
//}
//
//func (p *UserRepo) FindByUsername(c *gin.Context) {
//
//}

// Interface_Version

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userRepository struct {
}

func (u userRepository) GetUserByEmail(email string) bool {
	// Select * from user where email = email
	row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&models.GoCrmUser{}).RowsAffected
	return row != NumberNull
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
