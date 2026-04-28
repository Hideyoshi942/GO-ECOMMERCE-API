package service

import (
	"fmt"
	"go-ecomerce-backend-api/internal/repo"
	"go-ecomerce-backend-api/internal/utils/crypto"
	"go-ecomerce-backend-api/internal/utils/random"
	"go-ecomerce-backend-api/internal/utils/sendto"
	"go-ecomerce-backend-api/pkg/response"
	"strconv"
	"time"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
}

func NewUserService(userRepo repo.IUserRepository, userAuthRepo repo.IUserAuthRepository) IUserService {
	return &userService{
		userRepo:     userRepo,
		userAuthRepo: userAuthRepo,
	}
}

func (us *userService) Register(email string, purpose string) int {
	// 0. Hash email
	hashEmail := crypto.GetHash(email)
	fmt.Printf("Hash email: %s", hashEmail)

	// 5. check OTP is available

	// 6. user spam ...

	// 1. check email exists in db
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}

	// 2. new OTP
	otp := random.GenerateSixDigitOtp()
	if purpose == "TEST_USER" {
		otp = 123456
	}

	fmt.Printf("OTP for email %s: %d\n", email, otp)

	// 3. save OTP in redis with expiration time
	err := us.userAuthRepo.AddOTP(hashEmail, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrInvalidOTP
	}

	// 4. send email OTP
	//err = sendto.SendTemplateEmailOtp([]string{email}, "kietchu972@gmail.com", "otp-auth.html",
	//	map[string]interface{}{
	//		"otp": strconv.Itoa(otp),
	//	},
	//)
	//if err != nil {
	//	return response.ErrSendEmailOTP
	//}

	// send email OTP by JAVA
	err = sendto.SendEmailToJavaByAPI(strconv.Itoa(otp), email, "otp-auth.html")
	//fmt.Printf("err sendto : JAVA :: %v\n", err)
	if err != nil {
		return response.ErrSendEmailOTP
	}

	return response.ErrCodeSuccess
}
