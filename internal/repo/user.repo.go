package repo

import (
	"go-ecomerce-backend-api/global"
	"go-ecomerce-backend-api/internal/database"
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
	sqlc *database.Queries
}

func (up *userRepository) GetUserByEmail(email string) bool {
	// Select * from user where email = '?' order by email
	//row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&models.GoCrmUser{}).RowsAffected
	user, err := up.sqlc.GetUserByEmailSQLC(ctx, email)
	if err != nil {
		return false
	}
	return user.UsrID.Int32 != NumberNull
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}
