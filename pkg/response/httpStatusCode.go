package response

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 // EmailInvalid
	ErrInvalidToken     = 30001 // Token is invalid

	// Register Code
	ErrCodeUserHasExists = 50001 // User has already registered
)

// Message
var msg = map[int]string{
	ErrCodeSuccess:       "success",
	ErrCodeParamInvalid:  "Email is invalid",
	ErrInvalidToken:      "Invalid token",
	ErrCodeUserHasExists: "User Has Exists",
}
