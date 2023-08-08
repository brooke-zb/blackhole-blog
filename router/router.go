package router

import (
	"blackhole-blog/middleware/auth"
	"blackhole-blog/middleware/csrf"
	"blackhole-blog/middleware/logger"
	"blackhole-blog/middleware/no_route"
	"blackhole-blog/middleware/recovery"
	"blackhole-blog/middleware/security"
	"blackhole-blog/router/api/v2"
	"blackhole-blog/router/api/v2/admin"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.NoRoute(no_route.NoRoute())
	r.Use(logger.RouterLog())
	r.Use(recovery.Recovery())
	r.Use(csrf.Filter())
	r.Use(auth.Authorization())

	account := r.Group("/account")
	{
		account.POST("/token", v2.AccountLogin)              // 登录
		account.DELETE("/token", v2.AccountLogout)           // 登出
		account.GET("", v2.AccountInfo)                      // 获取账号信息
		account.PUT("", v2.AccountUpdateInfo)                // 更新账号信息
		account.PATCH("/password", v2.AccountUpdatePassword) // 更改账号密码
	}

	article := r.Group("/article")
	{
		article.GET("/:id", v2.ArticleFindById)                    // 获取文章详情
		article.GET("", v2.ArticleFindList)                        // 获取文章列表
		article.GET("/:id/comment", v2.CommentFindListByArticleId) // 获取文章评论列表
	}

	comment := r.Group("/comment")
	{
		comment.POST("", v2.CommentAdd) // 添加评论
	}

	category := r.Group("/category")
	{
		category.GET("/:name/article", v2.ArticleFindListByCategory) // 获取分类文章列表
		category.GET("", v2.CategoryFindList)                        // 获取分类列表
	}

	tag := r.Group("/tag")
	{
		tag.GET("/:name/article", v2.ArticleFindListByTag) // 获取标签文章列表
		tag.GET("", v2.TagFindList)                        // 获取标签列表
	}

	adminGroup := r.Group("/admin")
	{
		user := adminGroup.Group("/user", security.RequirePermission("USER:FULLACCESS"))
		{
			user.GET("/:id", admin.UserFindById)  // 获取用户详情
			user.GET("", admin.UserFindList)      // 获取用户列表
			user.POST("", admin.UserAdd)          // 添加用户
			user.PUT("/:id", admin.UserUpdate)    // 修改用户
			user.DELETE("/:id", admin.UserDelete) // 删除用户
		}

		role := adminGroup.Group("/role", security.RequirePermission("ROLE:FULLACCESS"))
		{
			role.GET("/:id", admin.RoleFindById)  // 获取角色详情
			role.GET("", admin.RoleFindList)      // 获取角色列表
			role.POST("", admin.RoleAdd)          // 添加角色
			role.PUT("", admin.RoleUpdate)        // 修改角色
			role.DELETE("/:id", admin.RoleDelete) // 删除角色
		}
	}

	return r
}
