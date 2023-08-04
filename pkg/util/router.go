package util

import (
	"blackhole-blog/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

// BindJSON is a wrapper of gin.Context.ShouldBindJSON.
// It will panic if any error occurs
func BindJSON(c *gin.Context, obj any) {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		panicReadable(err)
	}
}

// BindQuery is a wrapper of gin.Context.ShouldBindQuery.
// It will panic if any error occurs
func BindQuery(c *gin.Context, obj any) {
	err := c.ShouldBindQuery(obj)
	if err != nil {
		panicReadable(err)
	}
}

// BindUri is a wrapper of gin.Context.ShouldBindUri.
// It will panic if any error occurs
func BindUri(c *gin.Context, obj any) {
	err := c.ShouldBindUri(obj)
	if err != nil {
		panicReadable(err)
	}
}

// BindHeader is a wrapper of gin.Context.ShouldBindHeader.
// It will panic if any error occurs
func BindHeader(c *gin.Context, obj any) {
	err := c.ShouldBindHeader(obj)
	if err != nil {
		panicReadable(err)
	}
}

func panicReadable(err error) {
	if vErr, ok := err.(validator.ValidationErrors); ok {
		errMap := vErr.Translate(translator)
		errorMsg := "请求参数错误:"
		for _, val := range errMap {
			errorMsg += "\n" + val
		}
		panic(service.NewError(http.StatusUnprocessableEntity, errorMsg))
	}
	if _, ok := err.(*strconv.NumError); ok {
		panic(service.NewError(http.StatusUnprocessableEntity, "参数转换失败，类型错误"))
	}
	panic(err)
}
