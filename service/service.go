// Package service is a layer between router and dao, it is used to handle business logic.
// when error occurs, it will panic an Error, and should be handled in middleware.
package service

import (
	"blackhole-blog/pkg/util"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

var (
	User    = userService{}
	Article = articleService{}
)

func panicErrIfNotNil(err error) {
	if err != nil {
		panic(util.NewInternalError(err))
	}
}

func panicNotFoundErrIfNotNil(err error, notFoundMsg string) {
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			panic(util.NewError(http.StatusNotFound, notFoundMsg))
		} else {
			panic(util.NewInternalError(err))
		}
	}
}
