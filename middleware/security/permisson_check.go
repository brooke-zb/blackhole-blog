package security

import (
	"blackhole-blog/middleware/auth"
	"blackhole-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RequirePerm 检查用户是否有指定权限，包含其中一个权限即可通过
func RequirePerm(perm ...string) gin.HandlerFunc {
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

// RequireAllPerms 检查用户是否有指定权限，必须包含所有权限才能通过
func RequireAllPerms(perm ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := auth.MustGetUser(c)
		for _, p := range perm {
			hasPerm := false
			for _, rp := range user.Role.Permissions {
				if p == rp.Name {
					hasPerm = true
					break
				}
			}
			if !hasPerm {
				panic(util.NewError(http.StatusForbidden, "您没有权限执行该操作"))
			}
		}
	}
}
