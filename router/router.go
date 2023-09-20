package router

import (
	"blackhole-blog/middleware/auth"
	"blackhole-blog/middleware/csrf"
	"blackhole-blog/middleware/logger"
	"blackhole-blog/middleware/no_route"
	"blackhole-blog/middleware/recovery"
	"blackhole-blog/middleware/security"
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/setting"
	"blackhole-blog/router/api/v2"
	"blackhole-blog/router/api/v2/admin"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// set production mode
	if setting.Config.Production {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// proxy settings
	err := r.SetTrustedProxies(setting.Config.Server.Proxies)
	if err != nil {
		log.Default.Errorf("set trusted proxies fail with reason: %s", err.Error())
	}
	r.RemoteIPHeaders = setting.Config.Server.ProxyHeaders

	// set global middlewares
	r.NoRoute(no_route.NoRoute())
	r.Use(logger.RouterLog())
	r.Use(recovery.Recovery())
	r.Use(csrf.Filter())
	r.Use(auth.Authorization())

	// routes
	account := r.Group("/accounts")
	{
		account.POST("/tokens", v2.AccountLogin)             // 登录
		account.DELETE("/tokens", v2.AccountLogout)          // 登出
		account.GET("", v2.AccountInfo)                      // 获取账号信息
		account.PUT("", v2.AccountUpdateInfo)                // 更新账号信息
		account.PATCH("/password", v2.AccountUpdatePassword) // 更改账号密码
	}

	article := r.Group("/articles")
	{
		article.GET("/:id", v2.ArticleFindById)                     // 获取文章详情
		article.GET("", v2.ArticleFindList)                         // 获取文章列表
		article.GET("/:id/comments", v2.CommentFindListByArticleId) // 获取文章评论列表
	}

	comment := r.Group("/comments")
	{
		comment.POST("", v2.CommentAdd) // 添加评论
	}

	category := r.Group("/categories")
	{
		category.GET("/:name/articles", v2.ArticleFindListByCategory) // 获取分类文章列表
		category.GET("", v2.CategoryFindList)                         // 获取分类列表
	}

	tag := r.Group("/tags")
	{
		tag.GET("/:name/articles", v2.ArticleFindListByTag) // 获取标签文章列表
		tag.GET("", v2.TagFindList)                         // 获取标签列表
	}

	friend := r.Group("/friends")
	{
		friend.GET("", v2.FriendFindList) // 获取友链列表
	}

	adminGroup := r.Group("/admin")
	{
		user := adminGroup.Group("/users", security.RequirePerm("USER:FULLACCESS"))
		{
			user.GET("/:id", admin.UserFindById)  // 获取用户详情
			user.GET("", admin.UserFindList)      // 获取用户列表
			user.POST("", admin.UserAdd)          // 添加用户
			user.PUT("", admin.UserUpdate)        // 修改用户
			user.DELETE("/:id", admin.UserDelete) // 删除用户
		}

		role := adminGroup.Group("/roles", security.RequirePerm("ROLE:FULLACCESS"))
		{
			role.GET("/:id", admin.RoleFindById)  // 获取角色详情
			role.GET("", admin.RoleFindList)      // 获取角色列表
			role.POST("", admin.RoleAdd)          // 添加角色
			role.PUT("", admin.RoleUpdate)        // 修改角色
			role.DELETE("/:id", admin.RoleDelete) // 删除角色
		}

		category := adminGroup.Group("/categories", security.RequirePerm("CATEGORY:FULLACCESS"))
		{
			category.GET("/:id", admin.CategoryFindById)  // 获取分类详情
			category.GET("", admin.CategoryFindList)      // 获取分类列表
			category.POST("", admin.CategoryAdd)          // 添加分类
			category.PUT("", admin.CategoryUpdate)        // 修改分类
			category.DELETE("/:id", admin.CategoryDelete) // 删除分类
		}

		tag := adminGroup.Group("/tags", security.RequirePerm("TAG:FULLACCESS"))
		{
			tag.GET("/:id", admin.TagFindById)        // 获取标签详情
			tag.GET("", admin.TagFindList)            // 获取标签列表
			tag.POST("", admin.TagAdd)                // 添加标签
			tag.PUT("", admin.TagUpdate)              // 修改标签
			tag.DELETE("/*ids", admin.TagDeleteBatch) // 批量删除标签
		}

		article := adminGroup.Group("/articles", security.RequirePerm("ARTICLE:FULLACCESS"))
		{
			article.GET("/:id", admin.ArticleFindById)                  // 获取文章详情
			article.GET("", admin.ArticleFindList)                      // 获取文章列表
			article.POST("", admin.ArticleAdd)                          // 添加文章
			article.POST("/attachments", admin.ArticleUploadAttachment) // 上传附件
			article.PUT("", admin.ArticleUpdate)                        // 修改文章
			article.DELETE("/:id", admin.ArticleDelete)                 // 删除文章
		}

		comment := adminGroup.Group("/comments", security.RequirePerm("COMMENT:FULLACCESS"))
		{
			comment.GET("/:id", admin.CommentFindById)        // 获取评论详情
			comment.GET("", admin.CommentFindList)            // 获取评论列表
			comment.PUT("", admin.CommentUpdate)              // 修改评论
			comment.DELETE("/*ids", admin.CommentDeleteBatch) // 删除评论
		}

		friend := adminGroup.Group("/friends", security.RequirePerm("FRIEND:FULLACCESS"))
		{
			friend.GET("/:id", admin.FriendFindById)  // 获取友链详情
			friend.GET("", admin.FriendFindList)      // 获取友链列表
			friend.POST("", admin.FriendAdd)          // 添加友链
			friend.PUT("", admin.FriendUpdate)        // 修改友链
			friend.DELETE("/:id", admin.FriendDelete) // 删除友链
		}
	}

	return r
}
