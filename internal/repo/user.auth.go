package repo

import (
	"fmt"
	"go-ecomerce-backend-api/global"
	"time"
)

type IUserAuthRepository interface {
	AddOTP(email string, otp int, expiration int64) error
}

type userAuthRepository struct {
}

func (u *userAuthRepository) AddOTP(email string, otp int, expiration int64) error {
	key := fmt.Sprintf("user:%s:otp", email)
	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expiration)).Err()
}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}
