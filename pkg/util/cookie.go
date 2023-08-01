package util

import (
	"blackhole-blog/pkg/setting"
	"github.com/gin-gonic/gin"
)

func SetCookie(c *gin.Context, name string, value string, maxAge int, httpOnly bool) {
	c.SetCookie(name, value, maxAge, setting.Config.Server.Cookie.Path, setting.Config.Server.Cookie.Domain, setting.Config.Server.Cookie.Secure, httpOnly)
}
