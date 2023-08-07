package security

import (
	"blackhole-blog/middleware/auth"
	"blackhole-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RequireLogin 检查用户是否登录
func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth.MustGetUser(c)
	}
}

// RequirePermission 检查用户是否有指定权限，包含其中一个权限即可通过
func RequirePermission(perm ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		for _, p := range perm {
			for _, rp := range user.Role.Permissions {
				if p == rp.Name {
					return
				}
			}
		}
		panic(util.NewError(http.StatusForbidden, "您没有权限执行该操作"))
	}
}
