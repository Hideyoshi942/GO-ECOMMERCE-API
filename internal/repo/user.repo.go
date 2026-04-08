package repo

import "github.com/gin-gonic/gin"

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ur *UserRepo) GetInfoUser() string {
	return "hideyoshi"
}

func (p *UserRepo) FindByUsername(c *gin.Context) {

}
