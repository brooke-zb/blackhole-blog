package router

import (
	"blackhole-blog/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, util.RespOK("OK"))
}