package logger

import (
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/setting"
	"github.com/gin-gonic/gin"
)

func RouterLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		msg := "ok"
		if c.GetBool(setting.RecoveryAbortKey) {
			msg = "abort"
		}
		log.Api.Infow(msg, "status", c.Writer.Status(), "method", c.Request.Method, "path", c.Request.URL.Path, "query", c.Request.URL.RawQuery, "ip", c.ClientIP())
	}
}
