package util

import "net/http"

const (
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
		code:    http.StatusInternalServerError,
		message: InternalErrorMessage,
		err:     err,
	}
}
