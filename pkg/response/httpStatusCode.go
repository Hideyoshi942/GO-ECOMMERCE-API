package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 // EmailInvalid
	ErrInvalidToken     = 30001
)

// Message
var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email is invalid",
	ErrInvalidToken:     "Invalid token",
}
