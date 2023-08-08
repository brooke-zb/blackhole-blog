package admin

import (
	"blackhole-blog/middleware/auth"
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
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
	// bindings
	article := dto.ArticleAddDto{}
	util.BindJSON(c, &article)
	user := auth.MustGetUser(c)
	article.Uid = user.Uid

	// insert article
	service.Article.Add(article)
	c.JSON(http.StatusOK, util.RespMsg("添加成功"))
}

func ArticleUpdate(c *gin.Context) {
	// bindings
	article := dto.ArticleUpdateDto{}
	util.BindJSON(c, &article)

	// update article
	service.Article.Update(article)
	c.JSON(http.StatusOK, util.RespMsg("更新成功"))
}

func ArticleDelete(c *gin.Context) {
	// bindings
	param := models.IdParam{}
	util.BindUri(c, &param)

	// delete article
	service.Article.Delete(param.Id)
	c.JSON(http.StatusOK, util.RespMsg("删除成功"))
}
