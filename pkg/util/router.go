package util

import (
	"blackhole-blog/service"
	"github.com/gin-gonic/gin"
)

// BindJSON is a wrapper of gin.Context.ShouldBindJSON.
// It will panic if any error occurs
func BindJSON(c *gin.Context, obj any) {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		panic(service.NewError(422, err.Error()))
	}
}

// BindQuery is a wrapper of gin.Context.ShouldBindQuery.
// It will panic if any error occurs
func BindQuery(c *gin.Context, obj any) {
	err := c.ShouldBindQuery(obj)
	if err != nil {
		panic(service.NewError(422, err.Error()))
	}
}

// BindUri is a wrapper of gin.Context.ShouldBindUri.
// It will panic if any error occurs
func BindUri(c *gin.Context, obj any) {
	err := c.ShouldBindUri(obj)
	if err != nil {
		panic(service.NewError(422, err.Error()))
	}
}

// BindHeader is a wrapper of gin.Context.ShouldBindHeader.
// It will panic if any error occurs
func BindHeader(c *gin.Context, obj any) {
	err := c.ShouldBindHeader(obj)
	if err != nil {
		panic(service.NewError(422, err.Error()))
	}
}
