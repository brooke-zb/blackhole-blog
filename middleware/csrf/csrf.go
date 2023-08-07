package csrf

import (
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/setting"
	"blackhole-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

const (
	headerName = "X-CSRF-TOKEN"
)

// Filter CSRF过滤器，专事专干，只针对CSRF攻击原理进行防御
func Filter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 不拦截GET, OPTIONS及HEAD请求
		switch c.Request.Method {
		case "GET", "OPTIONS", "HEAD":
			return
		}

		// 检查路径是否在排除列表中
		url := c.Request.URL.Path
		for _, pattern := range setting.Config.Server.Csrf.ExcludePatterns {
			match, err := path.Match(pattern, url)
			if err != nil {
				log.Default.Error("CSRF Filter: path.Match fail with reason: " + err.Error())
				continue
			}
			if match {
				return
			}
		}

		// 检查是否存在token
		if token := c.Request.Header.Get(headerName); token == "" {
			panic(util.NewError(http.StatusForbidden, "缺少CSRF令牌"))
		}
	}
}
