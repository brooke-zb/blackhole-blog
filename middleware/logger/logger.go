package logger

import (
	"blackhole-blog/middleware/recovery"
	"blackhole-blog/pkg/log"
	"github.com/gin-gonic/gin"
)

func RouterLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		msg := "ok"
		if c.GetBool(recovery.AbortKey) {
			msg = "abort"
		}
		log.Api.Infow(msg, "status", c.Writer.Status(), "method", c.Request.Method, "path", c.Request.URL.Path, "query", c.Request.URL.RawQuery, "ip", c.ClientIP())
	}
}
