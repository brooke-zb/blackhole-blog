package no_route

import (
	"blackhole-blog/pkg/log"
	"blackhole-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, util.RespFail("404 Not Found"))
		log.Err.Errorf("%s %s : 404 Not Found", c.Request.Method, c.Request.URL.Path)
	}
}
