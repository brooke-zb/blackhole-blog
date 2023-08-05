// Package service is a layer between router and dao, it is used to handle business logic.
// when error occurs, it will panic an Error, and should be handled in middleware.
package service

import (
	"blackhole-blog/pkg/util"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var (
	User    = userService{}
	Article = articleService{}
)

type errorEntry struct {
	Code    uint16
	Message string
}

func entryErr(errCode uint16, errMsg string) errorEntry {
	return errorEntry{Code: errCode, Message: errMsg}
}

// panicErrIfNotNil panics a util.Error if err is not nil.
// use to handle dao error with custom error message.
func panicErrIfNotNil(err error, entries ...errorEntry) {
	if err == nil {
		return
	}
	mysqlErr, ok := err.(*mysql.MySQLError)
	if ok {
		for _, entry := range entries {
			if mysqlErr.Number == entry.Code {
				panic(util.NewError(http.StatusBadRequest, entry.Message))
			}
		}
	}
	panic(util.NewInternalError(err))
}

// panicSelectErrIfNotNil panics a util.Error if err is not nil.
// use to handle gorm.ErrRecordNotFound with custom error message.
func panicSelectErrIfNotNil(err error, msg string) {
	if err == nil {
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(util.NewError(http.StatusNotFound, msg))
	}
	panic(util.NewInternalError(err))
}
