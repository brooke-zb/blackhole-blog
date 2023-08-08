package v2

import (
	"blackhole-blog/middleware/auth"
	"blackhole-blog/models"
	"blackhole-blog/models/dto"
	"blackhole-blog/pkg/util"
	"blackhole-blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommentFindListByArticleId(c *gin.Context) {
	// bindings
	param := models.IdParam{}
	util.BindUri(c, &param)
	page := models.PageParam{}
	util.BindQuery(c, &page)

	// query
	comments := service.Comment.FindListByArticleId(param.Id, page.Page(), page.Size())
	c.JSON(http.StatusOK, util.RespOK(comments))
}

func CommentAdd(c *gin.Context) {
	// bindings
	body := dto.CommentAddDto{}
	util.BindJSON(c, &body)
	body.Ip = c.ClientIP()
	user, exist := auth.ShouldGetUser(c)
	if exist {
		body.Uid = &user.Uid
	}

	// insert
	service.Comment.Insert(dto.ToCommentModel(body))
	c.JSON(http.StatusOK, util.RespMsg("评论成功"))
}
