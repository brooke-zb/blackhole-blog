package util

import (
	"blackhole-blog/pkg/setting"
	"net/http"
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
		message: setting.InternalErrorMessage,
		err:     err,
	}
}
