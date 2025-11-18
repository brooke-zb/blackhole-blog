package v2

import (
	"blackhole-blog/pkg/util"
	"blackhole-blog/pkg/version"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, util.RespOK(version.Version))
}
