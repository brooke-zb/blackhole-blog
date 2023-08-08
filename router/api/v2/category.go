package v2

import (
	"blackhole-blog/pkg/util"
	"blackhole-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CategoryFindList(c *gin.Context) {
	c.JSON(http.StatusOK, util.RespOK(service.Category.FindListWithArticleCount()))
}
