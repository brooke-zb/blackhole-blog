// Package service is a layer between router and dao, it is used to handle business logic.
// when error occurs, it will panic an Error, and should be handled in middleware.
package service

import (
	"blackhole-blog/pkg/util"
	"errors"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var (
	User     = userService{}
	Article  = articleService{}
	Comment  = commentService{}
	Category = categoryService{}
	Tag      = tagService{}
	Role     = roleService{}
	Friend   = friendService{}
	AiChat   = aiChatService{}
)

type errorEntry struct {
	Code        uint16
	Message     string
	MsgProducer func(msg string) string
}

func entryErr(errCode uint16, errMsg string) errorEntry {
	return errorEntry{Code: errCode, Message: errMsg}
}

func entryErrProducer(errCode uint16, msgProducer func(msg string) string) errorEntry {
	return errorEntry{Code: errCode, MsgProducer: msgProducer}
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
				if entry.MsgProducer != nil {
					panic(util.NewError(http.StatusBadRequest, entry.MsgProducer(mysqlErr.Message)))
				}
				panic(util.NewError(http.StatusBadRequest, entry.Message))
			}
		}
	}
	panic(util.NewInternalError(err))
}

// panicNotFoundErrIfNotNil panics a util.Error if err is not nil.
// use to handle gorm.ErrRecordNotFound and other dao err with custom error message.
func panicNotFoundErrIfNotNil(err error, notFoundMsg string, entries ...errorEntry) {
	if err == nil {
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(util.NewError(http.StatusNotFound, notFoundMsg))
	}
	panicErrIfNotNil(err, entries...)
}
