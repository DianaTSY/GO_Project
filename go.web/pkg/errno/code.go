package errno

var(
	//Common Errors
	OK = &Errno{Code: 0, Message: "OK"}
	IntegernalServerError = &Errno{Code: 10001, Message: "Internal server errno"}
	ErrBind = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct "}

	ErrValidation = &Errno{Code: 20001, Message: "Validation Failed"}
	ErrDatabase = &Errno{Code: 20002, Message: "Database Error"}

	//User Errors
	ErrUserNotFound = &Errno{Code: 20101, Message:"The User was not found"}
	ErrPasswordIncorrect = &Errno{Code: 20102, Message: "The password is incorrect"}
)
