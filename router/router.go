package router

import (
	"blackhole-blog/middleware/auth"
	"blackhole-blog/middleware/logger"
	"blackhole-blog/middleware/no_route"
	"blackhole-blog/middleware/recovery"
	"blackhole-blog/router/api/v2"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.NoRoute(no_route.NoRoute())
	r.Use(logger.RouterLog())
	r.Use(recovery.Recovery())
	r.Use(auth.Authorization())

	account := r.Group("/account")
	{
		account.POST("/token", v2.AccountLogin)
		account.DELETE("/token", v2.AccountLogout)
		account.GET("", v2.AccountInfo)
		account.PUT("", v2.AccountUpdateInfo)
		account.PATCH("/password", v2.AccountUpdatePassword)
	}

	article := r.Group("/article")
	{
		article.GET("/:id", v2.ArticleFindById)
		article.GET("", v2.ArticleFindList)
		article.GET("/category/:name", v2.ArticleFindListByCategory)
		article.GET("/tag/:name", v2.ArticleFindListByTag)
	}

	return r
}
