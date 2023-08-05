package recovery

import (
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/setting"
	"blackhole-blog/pkg/util"
	"errors"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Set(setting.RecoveryAbortKey, true)
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					var se *os.SyscallError
					if errors.As(ne, &se) {
						seStr := strings.ToLower(se.Error())
						if strings.Contains(seStr, "broken pipe") ||
							strings.Contains(seStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				if brokenPipe {
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					log.Err.Errorf("%s %s : %s\n%s", c.Request.Method, c.Request.URL.Path, c.Errors.String(), string(debug.Stack()))
				} else if serviceErr, ok := err.(util.Error); ok {
					// service error
					c.AbortWithStatusJSON(serviceErr.Code(), util.RespFail(serviceErr.Error()))
					if serviceErr.E() != nil {
						log.Err.Errorf("%s %s : %s %s", c.Request.Method, c.Request.URL.Path, serviceErr.Error(), serviceErr.E())
					} else {
						log.Err.Errorf("%s %s : %s", c.Request.Method, c.Request.URL.Path, serviceErr.Error())
					}
				} else {
					// unknown error
					c.AbortWithStatusJSON(http.StatusInternalServerError, util.RespFail("发生了未知错误"))
					log.Err.Errorf("%s %s : 发生了未知错误: %v\n%s", c.Request.Method, c.Request.URL.Path, err, string(debug.Stack()))
				}
			}
		}()
		c.Next()
	}
}
