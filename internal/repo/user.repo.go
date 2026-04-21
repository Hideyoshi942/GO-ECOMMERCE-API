package repo

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
	NewUserRepository
}

type userRepository struct {
}

func (u userRepository) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
