package admin

import (
	"blackhole-blog/models"
	"blackhole-blog/pkg/util"
	"blackhole-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArticleFindById(c *gin.Context) {
	// bindings
	param := models.IdParam{}
	util.BindUri(c, &param)

	// query
	c.JSON(http.StatusOK, util.RespOK(service.Article.FindById(param.Id)))
}

func ArticleFindList(c *gin.Context) {
	// bindings
	clause := models.ArticleClause{}
	util.BindQuery(c, &clause)

	// query
	c.JSON(http.StatusOK, util.RespOK(service.Article.FindList(clause)))
}

func ArticleAdd(c *gin.Context) {
	// TODO
}

func ArticleUpdate(c *gin.Context) {
	// TODO
}

func ArticleDelete(c *gin.Context) {
	// TODO
}
