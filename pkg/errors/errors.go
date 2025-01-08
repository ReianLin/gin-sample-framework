package errors

import "fmt"

type Error struct {
	code    int
	message string
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Code() int {
	return e.code
}

func New(code int, message string) *Error {
	return &Error{
		code:    code,
		message: message,
	}
}

type ErrorCode struct {
	Code    int
	Message string
}

func (e ErrorCode) Error() string {
	return e.Message
}

func (e ErrorCode) String() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

// 预定义错误码
var (
	Success             = ErrorCode{Code: 0, Message: "success"}
	BadRequest          = ErrorCode{Code: 400, Message: "bad request"}
	Unauthorized        = ErrorCode{Code: 401, Message: "unauthorized"}
	Forbidden           = ErrorCode{Code: 403, Message: "forbidden"}
	NotFound            = ErrorCode{Code: 404, Message: "not found"}
	InternalServerError = ErrorCode{Code: 500, Message: "internal server error"}
	NotPermission       = ErrorCode{Code: 4003, Message: "no permission"}
)

// 业务错误码 (4xxx)
var (
	UserNotFound      = ErrorCode{Code: 4001, Message: "user not found"}
	RoleNotFound      = ErrorCode{Code: 4002, Message: "role not found"}
	InvalidParameter  = ErrorCode{Code: 4004, Message: "invalid parameter"}
	DuplicateUsername = ErrorCode{Code: 4005, Message: "duplicate username"}
	DuplicateEmail    = ErrorCode{Code: 4006, Message: "duplicate email"}
)
