package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	ErrParam            = &Errno{Code: 10003, Message: "get request param fail"}

	ErrValidation    = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase      = &Errno{Code: 20002, Message: "Database error."}
	ErrToken         = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	ErrAuthority     = &Errno{Code: 20004, Message: "no authority to call func"}
	ErrFileInit      = &Errno{Code: 20005, Message: "read or store file fail or no file upload"}
	ErrJsonMarshal   = &Errno{Code: 20006, Message: "Json Marshal fail"}
	ErrJsonUnmarshal = &Errno{Code: 20006, Message: "Json Unmarshal fail"}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid or expired."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
	ErrClassFound        = &Errno{Code: 20105, Message: "no class for this teacher"}
)
