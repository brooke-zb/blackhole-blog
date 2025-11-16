package v2

import (
	"blackhole-blog/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, util.RespOK("OK"))
}
