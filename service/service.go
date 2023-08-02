// Package service is a layer between router and dao, it is used to handle business logic.
// when error occurs, it will panic an Error, and should be handled in middleware.
package service

var User = userService{}

const (
	Unauthorized  = 401
	NotFound      = 404
	InternalError = 500

	UnauthorizedMessage  = "请登录后再进行操作"
	InternalErrorMessage = "内部错误，请联系管理员"
)

// Error is an interface for error with error code
type Error interface {
	error
	Code() int
	E() error
}

type errorImpl struct {
	code    int
	message string
	err     error
}

func (e *errorImpl) Error() string {
	return e.message
}

func (e *errorImpl) Code() int {
	return e.code
}

func (e *errorImpl) E() error {
	return e.err
}

func NewError(code int, message string) Error {
	return &errorImpl{code: code, message: message}
}

func NewInternalError(err error) Error {
	return &errorImpl{
		code:    InternalError,
		message: InternalErrorMessage,
		err:     err,
	}
}
