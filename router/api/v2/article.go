package v2

import (
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/util"
	"blackhole-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ArticleFindById(c *gin.Context) {
	// bindings
	param := dto.IdParam{}
	util.BindUri(c, &param)

	article := service.Article.FindById(param.Id)
	c.JSON(http.StatusOK, util.RespOK(article))
}

func ArticleFindList(c *gin.Context) {
	// bindings
	page := dto.PageParam{}
	util.BindQuery(c, &page)

	// query
	articles := service.Article.FindList(page.GetPage(), page.GetSize())
	c.JSON(http.StatusOK, util.RespOK(articles))
}

func ArticleFindListByCategory(c *gin.Context) {
	// bindings
	category := dto.StringParam{}
	util.BindUri(c, &category)
	page := dto.PageParam{}
	util.BindQuery(c, &page)

	// query
	articles := service.Article.FindListByCategory(category.Name, page.GetPage(), page.GetSize())
	c.JSON(http.StatusOK, util.RespOK(articles))
}

func ArticleFindListByTag(c *gin.Context) {
	// bindings
	tag := dto.StringParam{}
	util.BindUri(c, &tag)
	page := dto.PageParam{}
	util.BindQuery(c, &page)

	// query
	articles := service.Article.FindListByTag(tag.Name, page.GetPage(), page.GetSize())
	c.JSON(http.StatusOK, util.RespOK(articles))
}
