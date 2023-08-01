package router

import (
	"blackhole-blog/middleware"
	"blackhole-blog/router/api/v2"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(middleware.Authorization())

	account := r.Group("/account")
	{
		account.POST("/token", v2.AccountLogin)
		account.DELETE("/token", v2.AccountLogout)
	}

	return r
}
