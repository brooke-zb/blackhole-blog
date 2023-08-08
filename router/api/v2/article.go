package v2

import (
	"blackhole-blog/models"
	"blackhole-blog/pkg/setting"
	"blackhole-blog/pkg/util"
	"blackhole-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArticleFindById(c *gin.Context) {
	// bindings
	param := models.IdParam{}
	util.BindUri(c, &param)

	article := service.Article.FindById(param.Id)
	if article.Status != setting.StatusArticlePublished {
		panic(util.NewError(http.StatusNotFound, "未找到该文章"))
	}
	increment := service.Article.IncrAndGetReadCount(param.Id, c.ClientIP())
	article.ReadCount += increment
	c.JSON(http.StatusOK, util.RespOK(article))
}

func ArticleFindList(c *gin.Context) {
	// bindings
	page := models.PageParam{}
	util.BindQuery(c, &page)

	// query
	articles := service.Article.FindList(models.ArticleClause{
		PageParam: page,
		Status:    &setting.StatusArticlePublished,
	})
	c.JSON(http.StatusOK, util.RespOK(articles))
}

func ArticleFindListByCategory(c *gin.Context) {
	// bindings
	category := models.StringParam{}
	util.BindUri(c, &category)
	page := models.PageParam{}
	util.BindQuery(c, &page)

	// query
	articles := service.Article.FindList(models.ArticleClause{
		Category:  &category.Name,
		PageParam: page,
		Status:    &setting.StatusArticlePublished,
	})
	c.JSON(http.StatusOK, util.RespOK(articles))
}

func ArticleFindListByTag(c *gin.Context) {
	// bindings
	tag := models.StringParam{}
	util.BindUri(c, &tag)
	page := models.PageParam{}
	util.BindQuery(c, &page)

	// query
	articles := service.Article.FindList(models.ArticleClause{
		Tag:       &tag.Name,
		PageParam: page,
		Status:    &setting.StatusArticlePublished,
	})
	c.JSON(http.StatusOK, util.RespOK(articles))
}
