package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 // EmailInvalid
	ErrInvalidToken     = 30001 // Token is invalid
	ErrInvalidOTP       = 30002
	ErrSendEmailOTP     = 30003
	// Register Code
	ErrCodeUserHasExists = 50001 // User has already registered
)

// Message
var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "Invalid token",
	ErrInvalidOTP:        "Invalid OTP",
	ErrSendEmailOTP:      "Send email OTP failed",
	ErrCodeUserHasExists: "User Has Exists",
}
